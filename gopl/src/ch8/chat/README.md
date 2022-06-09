# 聊天服务器

聊天服务器可用于在几个用户之间相互广播消息

共4个goroutine
- 主goroutine
- 广播goroutine
- 每一个连接里面有一个连接处理goroutine和客户写入goroutine
