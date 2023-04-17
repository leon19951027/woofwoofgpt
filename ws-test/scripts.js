// 替换为你自己的WebSocket服务器地址
const wsUrl = "ws://127.0.0.1:8080/api/v1/stream-chat";

const messages = document.getElementById("messages");
const messageForm = document.getElementById("message-form");
const messageInput = document.getElementById("message-input");

const ws = new WebSocket(wsUrl);

ws.onopen = () => {
    console.log("WebSocket连接已建立");
};

ws.onmessage = (event) => {
    const message = document.createElement("div");
    message.textContent = event.data;
    messages.appendChild(message);
    messages.scrollTop = messages.scrollHeight;
};

ws.onclose = () => {
    console.log("WebSocket连接已关闭");
};

ws.onerror = (error) => {
    console.error("WebSocket发生错误", error);
};

messageForm.addEventListener("submit", (event) => {
    event.preventDefault();
    const message = messageInput.value.trim();
    if (message) {
        ws.send(message);
        messageInput.value = "";
    }
});
