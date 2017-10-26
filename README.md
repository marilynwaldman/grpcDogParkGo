
This is an example grpc recipe modeled from the following tutorial:

https://medium.com/@shijuvar/building-high-performance-apis-in-go-using-grpc-and-protocol-buffers-2eda5b80771b

To compile with protoc, run the following command from root directory (grpc directory):

protoc -I dogpark/ dogpark/dogpark.proto --go_out=plugins=grpc:dogpark
