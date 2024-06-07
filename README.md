# highwire
A high wire is a tight rope. It supports and provides firm footing. This project is a reference implementation of connectrpc which sends reliable data across the wire in a tight, reliable, wire-format. 

> See https://connectrpc.com/docs/go/getting-started for more information
~~~
go install github.com/bufbuild/buf/cmd/buf@latest
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install connectrpc.com/connect/cmd/protoc-gen-connect-go@latest
~~~

> Initializing the remote module
~~~
buf config init buf.build/sideshow/highwire
~~~

> Updating the dependencies. This allows us to import the remote protocol buffer definitions
~~~
buf dep update
~~~

> Fetching local version of connect
~~~
go get golang.org/x/net/http2
go get connectrpc.com/connect@v1.16.2
go get buf.build/gen/go/sideshow/coffeeco/protocolbuffers/go@latest
~~~

> Update Go Modules
~~~
go mod tidy
~~~

## Building the ConnectRPC Client

~~~
go get github.com/google/uuid
~~~

### Fun with Buf

Initially the thought was we'd need to do an export from BSR to utilize the local protobufs.
~~~
sync:
	@buf export buf.build/sideshow/coffeeco -o proto/
~~~

But that is not actually the case. Instead, this just gives us a nice way 
to exchange proto definitions and their dependencies.
