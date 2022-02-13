 protoc -I=. -I=$GOPATH/pkg/mod/github.com/gogo/protobuf@v1.3.2/ -I=$GOPATH/pkg/mod --gogo_out=. proto/user/users.proto

 protoc --go_out=. --go-grpc_out=. proto/user/users.proto