package main

import (
	"fmt"
	"log"

	"github.com/maooz4426/hello"
)

func main() {
	mes, err := hello.Hello("maoz")
	log.Println(err)
	fmt.Println(mes)
}
