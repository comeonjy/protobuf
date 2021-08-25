proto:
	find . -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && protoc -I=. -I=$$GOPATH/bin --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. --grpc-gateway_out=paths=source_relative,logtostderr=true,generate_unbound_methods=true:. *.proto'

clear:
	find . -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && rm *.pb.go'


