# Github tools
## search
功能：
- 查询github repo的issues
  - 用法  
    `github search issues -q=xxx`  
  - 例：搜索golang/go项目的issue跟踪接口，查找关于JSON编码的Open状态的bug列表  
    `github search issues -q=repo:golang/go is:open json decoder`
- 按照创建时间排列结果
  - 用法  
    `github search issues -q=xxx -s=xxx`
    - 选项： 
      - 按数量
        - comments
        - reactions
        - reactions-+1
        - reactions--1
        - reactions-smile
        - reactions-thinking_face
        - reactions-heart
        - reactions-tada
        - interactions
      - 按时间
        - created
        - updated
  - 例：以创建时间降序排列  
      `github search issues -q=repo:golang/go is:open json decoder -s=created`

## issues
功能：
通过命令行创建、读取、更新或者关闭Github的issues，当需要额外输入时可调用文本编辑器 
  - 创建
  - 读取
  - 编辑
  - 更新
  - 关闭
# Todo
- 调整搜索结果时间字段右对齐
- 增加按时间段搜索功能