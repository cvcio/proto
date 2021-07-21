REGISTRY=reg.cvcio.org

tools: gen-paths
	go get \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
		github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
		google.golang.org/protobuf/cmd/protoc-gen-go \
		google.golang.org/grpc/cmd/protoc-gen-go-grpc \
		github.com/golangci/golangci-lint \
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/annotations.proto > google/api/annotations.proto
	curl https://raw.githubusercontent.com/googleapis/googleapis/master/google/api/http.proto > google/api/http.proto

gen-proto: 
	protoc \
		--proto_path proto \
		--proto_path . \
		proto/cvcio/**/*.proto \
			--go_out=Mtransformer/annotations.proto=github.com/bold-commerce/protoc-gen-struct-transformer/options:. --go_opt paths=source_relative \
			--struct-transformer_out=package=transformer,goimports=false:. \
			--go-grpc_out=:. --go-grpc_opt paths=source_relative \
			--grpc-gateway_out=:. --grpc-gateway_opt paths=source_relative \
			--openapiv2_out=:swagger

gen-paths:
	rm -rf internal/cvcio
	mkdir -p internal/cvcio swagger proto/google/api

protobuf: gen-paths gen-proto