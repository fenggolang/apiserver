实战：启动一个最简单的RESTful API服务器

项目使用govendor做包管理工具
```bash
# 1. 下载
go get -u -v github.com/kardianos/govendor

# 2. 初始化项目为vendor管理的项目
cd 项目根目录
govendor init

# 3. 去GOPATH中去找本项目依赖的包到vendor目录中
govendor add +external
```