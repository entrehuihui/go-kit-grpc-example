.PHONY: build generate proto clean

build:
	@echo "todo build shell"

generate: clean
	@echo "generate postgres migrate"w
	go generate ./orderpostgres/migrate.go

proto:
	# protoc --go_out=plugins=grpc:. pb/*.proto 
	protoc --proto_path=../  -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./proto/*.proto
	protoc --proto_path=../  -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
	protoc --proto_path=../  -I. -I$GOPATH/src -I$GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. ./proto/*.proto
clean:
	@rm -f ./orderpostgres/bindata.go