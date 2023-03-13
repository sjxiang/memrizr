

# 单词本，百词斩



=== Model / Domain

User
Token
Interfaces
Errors


从下到上


Handler 
__________

解析 / 验证 传入请求、调用服务



Service / Usecase
______________________

User
Token


Repository / Data
____________________

User
Image
Token
Events


Data Sources
______________

Redis
Postgres
Cloud Storge
PubSub