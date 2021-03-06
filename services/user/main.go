package main

import (
	"fmt"

	hello "github.com/notsu/gomono/services/post/hello"
	helloV4 "github.com/notsu/gomono/services/post/v4/hello"
	helloV5 "github.com/notsu/gomono/services/post/v5/hello"
	helloV6 "github.com/notsu/gomono/services/post/v6/hello"
)

func main() {
	fmt.Println("Hello from USER")
	fmt.Println(hello.Hello())
	fmt.Println(helloV4.Hello())
	fmt.Println(helloV5.Hello())
	fmt.Println(helloV6.Hello())
}
