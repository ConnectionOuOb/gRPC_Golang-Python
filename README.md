# gRPC

## Server : Python

### Install
```pip install grpcio grpcio-tools```


### Build
```python -m grpc_tools.protoc --proto_path=. --python_out=. --grpc_python_out=. proto/calc.proto```



## Client : Golang

### Install
```go install google.golang.org/protobuf/cmd/protoc-gen-go@latest```


### Build
```protoc --proto_path=. --go_out=. --go-grpc_out=. proto/calc.proto```


