package main

import (
	"fmt"

	greetv1 "github.com/abitofhelp/connect-go-example/gen/greet/v1"
	greetv1connect "github.com/abitofhelp/connect-go-example/gen/greet/v1/greetv1connect"
)

func main() {
	r := greetv1.GreetRequest{
		Name: "Mike",
	}
	fmt.Printf("Request: %+v, %s", r.String(), greetv1connect.GreetServiceName)
}
