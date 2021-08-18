proto:
	find . -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && protoc -I=. --go_out=paths=source_relative:. --go-grpc_out=paths=source_relative:. *.proto'


clear:
	find . -type d -depth 1 -print | xargs -L 1 bash -c 'cd "$$0" && rm *.pb.go'
