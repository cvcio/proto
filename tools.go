// +build tools linux darwin windows

package tools

import (
	_ "github.com/amsokol/protoc-gen-gotag"
	_ "github.com/amsokol/protoc-gen-gotag/tagger"
	_ "github.com/bold-commerce/protoc-gen-struct-transformer"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway"
	_ "github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2"
	_ "google.golang.org/grpc/cmd/protoc-gen-go-grpc"
	_ "google.golang.org/protobuf/cmd/protoc-gen-go"
)
