#! /bin/bash

# 编译google.api
protoc -I . --go_out=plugins=grpc,Mgoogle/protobuf/descriptor.proto=github.com/golang/protobuf/protoc-gen-go/descriptor:. google/api/*.proto

#进入proto目录， 编译hello.proto
protoc -I . --go_out=plugins=grpc,Mgoogle/api/annotations.proto=grpc-hello-world/proto/google/api:. hello/hello.proto

#hello.proto为hello.pb.gw.proto
protoc --grpc-gateway_out=logtostderr=true:. hello/hello.proto


#生成datafile.go文件
go-bindata --nocompress -pkg swagger -o pkg/ui/data/swagger/datafile.go third_party/swagger-ui/...

#生成hello.swagger.json
protoc -I/usr/local/include -I. -I$GOPATH/src/grpc-hello-world/proto/google/api --swagger_out=logtostderr=true:. hello/hello.proto
