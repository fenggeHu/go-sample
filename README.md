# go-sample
打包命令, 编译websocket目录下的所有文件，生产可执行文件ws
```shell
go build -o quote-ws websocket/main.go  # 本地编译
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o quote-ws websocket/main.go  # kafka依赖不支持交叉编译
```

# samples
kafka
gin
websocket <- kafka
