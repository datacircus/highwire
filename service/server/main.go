package main

import (
	"connectrpc.com/connect"
	"context"
	"fmt"
	coffeeservicev1 "github.com/datacircus/highwire/gen/coffeeservice/v1"
	"github.com/datacircus/highwire/gen/coffeeservice/v1/coffeeservicev1connect"
	"golang.org/x/net/http2"
	"golang.org/x/net/http2/h2c"
	"log"
	"net/http"
)

type CoffeeserviceServer struct{}

func (s *CoffeeserviceServer) CoffeeOrder(ctx context.Context,
	req *connect.Request[coffeeservicev1.CoffeeOrderRequest]) (*connect.Response[coffeeservicev1.CoffeeOrderResponse], error) {
	log.Println("Request Headers: ", req.Header())
	var order = req.Msg.Order
	res := connect.NewResponse(&coffeeservicev1.CoffeeOrderResponse{
		Response: fmt.Sprintf("Thanks for the Order, %s\n", order.Customer.Name),
	})
	res.Header().Set("CoffeeService-Version", "v1")
	return res, nil
}

func main() {
	coffeeservice := &CoffeeserviceServer{}
	mux := http.NewServeMux()
	path, handler := coffeeservicev1connect.NewCoffeeServiceHandler(coffeeservice)
	mux.Handle(path, handler)
	log.Println(fmt.Sprintf("CoffeeService:OpenForBusiness (port=8080)"))
	http.ListenAndServe("localhost:8080", h2c.NewHandler(mux, &http2.Server{}))
}