# Proto Hub

## 安装插件

```shell

go get -u github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/grpc/cmd/protoc-gen-go-grpc \
google.golang.org/protobuf/cmd/protoc-gen-go \
github.com/envoyproxy/protoc-gen-validate

go install \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc \
github.com/envoyproxy/protoc-gen-validate

```