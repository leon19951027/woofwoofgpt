package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/leon19951027/woofwoofgpt/config"
	"github.com/leon19951027/woofwoofgpt/web/controller"
)

type Web struct {
	Host       string
	Port       string
	JwtSecrect string
}

func (w *Web) ApplyCfg(cfg *config.Cfg) {
	w.Host = cfg.Host
	w.Port = cfg.Port
	w.JwtSecrect = cfg.JwtSecrect
	controller.OpenaiApiToken = cfg.ApiToken
	controller.OpenaiApiUrlPrefix = cfg.UrlPrefix

}

func (w *Web) Run() {
	h := gin.Default()
	//h.POST("/login")
	v1 := h.Group("/api/v1", controller.Login)
	//v1.Use(w.jwtMiddleware())
	v1.GET("/stream-chat", controller.Stream_Chat)
	v1.POST("/chunk-chat", controller.Chunk_Chat)
	h.Run(w.Host + ":" + w.Port)
}

func (w *Web) generateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"expire":   time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString(w.JwtSecrect)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (w *Web) jwtMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.Request.Header.Get("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing"})
			c.Abort()
			return
		}

		token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
			}
			return w.JwtSecrect, nil
		})

		fmt.Println(token, tokenString)

		// if err != nil || !token.Valid {
		// 	fmt.Println("====")
		// 	fmt.Println(err)
		// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		// 	c.Abort()
		// 	return
		// }

		c.Next()
	}
}
