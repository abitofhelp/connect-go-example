package main

import (
	"fmt"
	greetv1 "github.com/abitofhelp/connect-go-example/bazel-bin/proto/greet/v1/greetv1_go_proto_/connect-go-example/gen/greet/v1"
)

func main() {

	r := greetv1.GreetRequest{
		Name: "Mike",
	}
	fmt.Printf("Request: %+v", r.String())
}
