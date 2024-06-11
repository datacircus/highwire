package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	protovalidate "github.com/bufbuild/protovalidate-go"
	"github.com/confluentinc/confluent-kafka-go/kafka"
	coffeeservicev1 "github.com/datacircus/highwire/gen/coffeeservice/v1"
	"github.com/datacircus/highwire/gen/coffeeservice/v1/coffeeservicev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"google.golang.org/protobuf/proto"
	"log"
	"net/http"
)

type CoffeeserviceServer struct {
	Kafka     *kafka.Producer
	TopicName string
	Validator *protovalidate.Validator
}

func (s *CoffeeserviceServer) CoffeeOrder(ctx context.Context,
	req *connect.Request[coffeeservicev1.CoffeeOrderRequest]) (*connect.Response[coffeeservicev1.CoffeeOrderResponse], error) {
	log.Println("Request Headers: ", req.Header())
	var order = req.Msg.Order
	
	response := ""

	if err := s.Validator.Validate(req.Msg); err != nil {
		log.Println("validation failed:", err)
		response = err.Error()
	} else {
		data, err := proto.Marshal(order)
		if err != nil {
			log.Println(err)
		}
		err = s.Kafka.Produce(&kafka.Message{
			TopicPartition: kafka.TopicPartition{
				Topic:     &s.TopicName,
				Partition: kafka.PartitionAny,
			},
			Key:   []byte(order.Customer.Name),
			Value: data,
		}, nil)
		if err != nil {
			log.Println("Failed to Send Order : kafka-rp:29092/coffeeco.v1.orders")
			response = fmt.Sprintf("We Failed to Send your Order, %s\n", order.Customer.Name)
		} else {
			log.Println("Order Published Successfully : kafka-rp:29092/coffeeco.v1.orders")
			response = fmt.Sprintf("Thanks for the Order, %s\n", order.Customer.Name)
		}
	}
	log.Println("New Order: ", order)
	res := connect.NewResponse(&coffeeservicev1.CoffeeOrderResponse{
		Response: response,
	})
	res.Header().Set("CoffeeService-Version", "v1")
	return res, nil
}

func genKafkaProducer() *kafka.Producer {
	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:29092",
		"acks":              "all",
	})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Kafka Producer Created Successfully : localhost:29092")
	return p
}

func main() {
	/* add the protovalidate instance */
	v, err := protovalidate.New()
	if err != nil {
		fmt.Println(err)
	}
	/* to the coffeeservice */
	coffeeservice := &CoffeeserviceServer{
		Kafka:     genKafkaProducer(),
		TopicName: "coffeeco.v1.orders",
		Validator: v,
	}
	mux := http.NewServeMux()
	path, handler := coffeeservicev1connect.NewCoffeeServiceHandler(coffeeservice)
	mux.Handle(path, handler)
	log.Println(fmt.Sprintf("CoffeeService:OpenForBusiness (port=8080)"))
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}
