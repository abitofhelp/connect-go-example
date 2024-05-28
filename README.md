# connect-go-example

This repository contains a client/server solution implementing the classic  Greet service using the following technologies:
* Bazel
* Buf
* ConnectRPC
* Go
* Protocol Buffers

## File Generation
There are three ways to generate the Protocol Buffer and ConnectRPC files:

1) protoc -I proto --go_out='gen/' --go_opt=paths=source_relative --connect-go_out='gen/' --connect-go_opt=paths=source_relative $(find proto -name '*.proto')

2) buf generate proto

3) bazel build
    * bazel build //proto/greet/v1:greetv1_connect_go_proto \
      This command will generate the Protocol Buffers & Connect RPC files.
    * bazel build //proto/greet/v1:greetv1_connect_go_proto \
      and \
      bazel build //proto/greet/v1:greetv1_go_proto \
      The first command generates the Protocol Buffer file. \
      The second command generates the ConnectRPC file. \
      It doesn't generate the Protocol Buffer file unless it has changed.

## Build
There are two alternatives to build the solution depending on how you \
generated the Protocol Buffer and ConnectRPC files:
1) Go Build From protoc or buf \
   go build ./...

2) Bazel Build  
   bazel build //...

### IMPORTANT: If you want to build interchangeably using 'go build' or 'bazel build', the import paths in the generated files must be consistent.  Essentially, each path in the plugin's 'out' field must be the root of the importpath & overrideimportpath in the go_proto_library blocks in proto/greet/v1/BUILD.bazel. Here is an example from this repository's code showing proper alignment between the configuration files using a path starting at 'gen':

```
# buf.get.yaml (redacted to relevant fields) 
plugins:
- remote: buf.build/protocolbuffers/go 
  out: gen/
- remote: buf.build/connectrpc/go 
  out: gen/
```

```
# proto/greet/v1/BUILD.bazel (redacted to relevant fields)
go_proto_library(
name = "greetv1_connect_go_proto",
importpath = "github.com/abitofhelp/connect-go-example/gen/greet/v1",
overrideimportpath = "github.com/abitofhelp/connect-go-example/gen/greet/v1/greetv1connect",
)

go_proto_library(
name = "greetv1_go_proto",
importpath = "github.com/abitofhelp/connect-go-example/__gen__/greet/v1",
)
```

## GreetService Queries
* grpcurl -plaintext -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
* grpcurl -plaintext -proto proto/greet/v1/greet.proto -d '{"name": "Jane"}' localhost:8080 greet.v1.GreetService/Greet
* curl -X POST -H "Content-Type: application/json" -d '{"name": "Jane"}' http://127.0.0.1:8080/greet.v1.GreetService/Greet
* buf curl --http2-prior-knowledge --schema=proto -d '{"name": "Jane"}' http://127.0.0.1:8080/greet.v1.GreetService/Greet

## GreetService Health Queries
* grpcurl -plaintext -d '{"service":"greet.v1.GreetService"}' localhost:8080 grpc.health.v1.Health/Check
* curl -X POST -H "Content-Type: application/json" -d '{"service":"greet.v1.GreetService"}' http://127.0.0.1:8080/grpc.health.v1.Health/Check
* buf curl --http2-prior-knowledge -d '{"service":"greet.v1.GreetService"}' http://127.0.0.1:8080/grpc.health.v1.Health/Check


## Resources
* https://connectrpc.com
* https://github.com/connectrpc/connect-go
* https://github.com/connectrpc/examples-go
* https://github.com/connectrpc/grpchealth-go
* https://github.com/connectrpc/grpcreflect-go