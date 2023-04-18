# 使用
## 启动
保证可执行文件同目录下有config.yaml,之后直接启动即可
## 接口
两个接口，data raw一样的
```json
{
    "messages": [
        {
            "message": "你是一个sql命令行翻译程序，你可以将人类自然语言描述的指令翻译成对应的命令行语句。",
            "role":"system"          
        },
        {
            "message":"创建一张表并授权给user@'%'所有权限"
        }
    ]
}
```
/api/v1/stream-chat是ws传输
/api/v1/chunk-chat是http chunk传输
