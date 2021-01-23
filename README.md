# gRPC demo with GoLang

Simple client-server communication. Just a simple message being sent from the client, printed by the server, and the response being print by the client too.

Command lines required to build this demo:

```
brew install protobuf
go get google.golang.org/protobuf/cmd/protoc-gen-go \
google.golang.org/grpc/cmd/protoc-gen-go-grpc
```

### How to run it?

Run the server

`go run ./server.go`

Run the client 

`go run ./client.go`