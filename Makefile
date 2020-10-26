.PHONY: kong build build_image vendor gen

kong:
	@docker cp proto hello-grpc_kong_1:/usr/local
	@deck sync -s kong/kong.yaml

build_image: build
	@docker build . --tag hello

build:
	@env GOOS=linux GOARCH=amd64 \
	go build \
	-o build/hello \
	main.go

vendor:
	@go mod tidy
	@go mod vendor

gen:
	@mkdir -p gen/go
	@protoc -Iproto \
    --go_out ./gen/go/ \
    --go_opt paths=source_relative \
    --go-grpc_out ./gen/go/ \
    --go-grpc_opt paths=source_relative \
    --grpc-gateway_out ./gen/go \
    --grpc-gateway_opt logtostderr=true \
    --grpc-gateway_opt paths=source_relative \
    proto/ivanrylach/hello_grpc/v1/hello.proto