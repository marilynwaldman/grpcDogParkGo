
This is an example grpc recipe modeled from the following tutorial:

https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b

First go get the dependencies from the cmd line:

1. go get -u google.golang.org/grpc
2. go get -u github.com/golang/protobuf/protoc-gen-go

To compile with protoc, run from the cmd line the following command from root

protoc -I dogpark/ dogpark/dogpark.proto --go_out=plugins=grpc:dogpark

Run the server:
   go run server/main.go
   
Run the client:
   go run client/main.go  
