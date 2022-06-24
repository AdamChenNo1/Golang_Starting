# server

1. 要求创建一个 Web 服务器，用户可以在其中跟踪玩家赢了多少场游戏。
   - GET /players/{name} 应该返回一个表示获胜总数的数字
   - POST /players/{name} 应该为玩家赢得游戏记录一次得分，并随着每次 POST 递增
2. /league 返回一个玩家清单，返回JSON格式的数据。