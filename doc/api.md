# Millionare 项目后台 API 部分
### 创建项目
- 在`$PATH`/src mkdir -p lightsaid.com/millionare
- 初始化 go mod init lightsaid.com/millionare


### docker 

### mongodb 驱动降级
- 先是安装了 1.9.x版本, 出行依赖错误
- 解决：
    1. go get go.mongodb.org/mongo-driver/mongo@v1.8.2
    1. go mod edit -require=go.mongodb.org/mongo-driver/mongo@v1.8.2
    1. go mod tidy
