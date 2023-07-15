# GRPC

## Prerequisite
1. Create google folders by using `mkdir ./pkg/google` and `mkdir ./pkg/google/api`
2. Download `annotations.proto` and `http.proto` files from `googleapis github repo`: https://github.com/googleapis/googleapis/blob/master/google/api
3. Add those files to `./pkg/google/api`

## Installation
1. Install Protocol Buffer Compiler: https://grpc.io/docs/protoc-installation/
2. Install Go plugins for protocol compiler:
   - Run `go install google.golang.org/protobuf/cmd/protoc-gen-go@latest`
   - Run `go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest`
   - Run `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway@latest`
   - Run `go install github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2@latest`
3. Export protoc: `export PATH="$PATH:$(go env GOPATH)/bin"`

## How to run gRPC
1. Generate go proto files by using `protoc  --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative .\proto\*.proto`
2. Move generated files to suitable folder, `e.g common.proto, then the generated files will be move to common folder`
3. Generate Transcode HTTP over gRPC `protoc -I . --grpc-gateway_out=logtostderr=true:. .\proto\*.proto`
4. Fix any errors if having