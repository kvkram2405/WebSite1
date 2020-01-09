package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	// Split port and Host

	host, port, err := net.SplitHostPort("127.0.0.1:5432")
	fmt.Println(host, port, err)

	// Split username and domain

	username := strings.Split("vkalipat@gmail.com", "@")
	fmt.Printf("%q\n", username[0])
}
