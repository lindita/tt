# tt
go framework, go自用api框架,集成gin,gorm,go-redis,wire等

## 结构
server->controller->service->dao->[db,redis..]
只能从上往下调用

## 启动server
go run main.go server


## 执行脚本
go run main.go tt