package v1

//go:generate sh -c "protoc -I=. -I=$GOPATH/bin --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative,logtostderr=true,generate_unbound_methods=true:. *.proto"
