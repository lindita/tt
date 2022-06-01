# tt
go api框架
http:gin
mysql:gorm
redis:go-redis
依赖注入:wire

## 结构
server->controller->service->dao->[mysql,redis..]
task->service->dao->[mysql,redis..]
只能从上往下调用

## 启动server
go run main.go server


## 执行脚本
go run main.go tt