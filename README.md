# gRPC Demo

This is a gRPC demo.

## Quick start

### Install package

1. Install the protocol compiler plugins for Go using the following commands:

    ``` shell
    go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.26
    go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.1
    ```

    fixed protoc error:

    ``` shell
    go get -u github.com/golang/protobuf/protoc-gen-go@v1.3.2
    ```

2. Update your PATH so that the protoc compiler can find the plugins:

    ``` shell
    export PATH="$PATH:$(go env GOPATH)/bin"
    ```

### Regenerate gRPC code

``` shell
protoc --go_out=plugins=grpc:. product/ProductInfo.proto
```

## procto

- AddProduct
- GetProduct
