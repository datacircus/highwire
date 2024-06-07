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

---

# Running the End-to-End with Kafka

> Need to have `docker` installed
> Need to create the `dais` docker network

**Create the DAIS docker network**
~~~
docker network create dais
~~~

**Spin up the Kafka Service**
~~~
docker-compose \
  -f docker-compose-service_server.yaml up \
  --remove-orphans
~~~

## Bootstrap the Kafka Topic

1. Pop on over to the Kafka Container
~~~
docker exec -it kafka-rp bash
~~~

2. Create the `coffeeco.v1.orders` Topic
> Note: The protobuf package for the protobuf is `coffeeco.v1` so for the topic we are reusing the same package
> naming conventions and tracking the `type` of events we expect to receive (in this case `coffeeco.v1.Order`) as
> the plural `*.orders`.

> Run: within the `docker container` (Step 1)
~~~
rpk topic create coffeeco.v1.orders --brokers=localhost:9092
~~~

3. Ensure the Topic exists
> From within the `docker container` (Step 1) run the following.
~~~
rpk topic list
~~~

4. Ensure you can hit the Kafka endpoint from your local machine
Exit the docker bash session, or open a new terminal window, and test the connection.
~~~
nc -z 127.0.0.1 29092
~~~
You should see `Connection to 127.0.0.1 port 29092 [tcp/*] succeeded!`. If you don't, then you just need to make sure
`kafka` is running. You can modify the command in step 1 to use the `-d` flag to detach so the process doesn't close 
when you close the terminal window.

## Running the Full End to End
> You will need 3 terminal windows. 
> 1. The first will be to run `docker compose -f docker-compose-service_server.yaml`
> 2. The second will be to run the connect service `go run ./service/server/main.go`
> 3. The third will be to generate and send `CoffeeOrder`s. `go run ./service/server/main.go`

At this point, you can run `go run ./service/server/main.go` as many times as you need to create more coffee orders. 
You can also modify the `service/client/main.go` to add new generators.

To inspect the data in kafka. 

~~~
docker exec -it kafka-rp rpk topic consume coffeeco.v1.orders --brokers=localhost:9092
~~~

> You can also just monitor using `rpk topic consume coffeeco.v1.orders --brokers=localhost:9092 --offset end`

### Debugging
> Can't resolve `kafka-rp` from outside of the Docker Container
~~~
%3|1717796251.070|FAIL|rdkafka#producer-1| [thrd:kafka-rp:29092/0]: kafka-rp:29092/0: Failed to resolve 'kafka-rp:29092': nodename nor servname provided, or not known (after 41ms in state CONNECT)
~~~

Solution:
> Add `kafka-rp localhost` to your `etc/host

```bash
127.0.0.1 kafka-rp
```

