# 环境安装 
### https://go-zero.dev/cn/docs/prepare/protoc-install/
> brew install protobuf protoc-gen-go protoc-gen-go-grpc
#### 环境查看命令
> protoc --version && which protoc-gen-go

# 依赖包
> go get -u -v google.golang.org/grpc

# 生成go代码
#### ①两种方法任选一种
> protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative  ./proto/both_stream.proto

> protoc --go_out=. ./proto/both_stream.proto

# demo借鉴
#### https://bingjian-zhu.github.io/2020/04/14/Go-gRPC%E6%95%99%E7%A8%8B-%E5%8F%8C%E5%90%91%E6%B5%81%E5%BC%8FRPC%EF%BC%88%E4%BA%94%EF%BC%89/

# 启动服务
> go run server/server.go
> go run client/client.go
> 

# 官网
#### https://grpc.io/docs/languages/go/quickstart/#regenerate-grpc-code