# Kong + gRPC-gateway
A prototype of serving both HTTP and gRPC traffic from a single spec.

HTTP traffic is proxied to the gRPC server by Kong with `grpc-gateway` [plugin](https://docs.konghq.com/hub/kong-inc/grpc-gateway/).

## Spec
`proto` [files](./proto)

## Run Server
```shell
# 1. build the gRPC server and its docker image
make build_image

# 2. start kong and gRPC server
docker-compose up

# 3. create Kong resources
make kong
```

## Client

Requires:
* `grpcurl` - https://github.com/fullstorydev/grpcurl

### Read entity
```shell
# HTTP
curl localhost:8000/v1/messages/qwe

# gRPC
grpcurl -plaintext -d '{"name":"qwe"}' localhost:9000 ivanrylach.hello_grpc.v1.HelloService/SayHello
```

### Create entity
```shell
# HTTP
curl -XPOST localhost:8000/v1/messages -d '{"name":"world"}'

# gRPC
grpcurl -plaintext -d '{"name":"world"}' localhost:9000 ivanrylach.hello_grpc.v1.HelloService/CreateHello
```

## Caveats

`proto` files for gRPC services with HTTP annotations require following import:
```protobuf
import "google/api/annotations.proto";
```

`grpc-gateway` fails to import these definitions, when the parent file is copied inside of Kong docker image.
It is commented out for now. 