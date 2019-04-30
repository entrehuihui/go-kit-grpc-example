echo windowbuild

echo "generate postgres migrate"
go generate ./orderpostgres/migrate.go

echo "grpc protoc"
protoc --proto_path=../  -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --go_out=plugins=grpc:. ./proto/*.proto
protoc --proto_path=../  -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --grpc-gateway_out=logtostderr=true:. ./proto/*.proto
protoc --proto_path=../  -I. -I%GOPATH%/src -I%GOPATH%/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis --swagger_out=logtostderr=true:. ./proto/*.proto

echo "build swagger UI JSON"
go get github.com/entrehuihui/aias
aias -d ./proto