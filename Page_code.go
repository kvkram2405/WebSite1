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

var i int
 	fmt.Print("Enter an integer to generate the multiplication table : ")
 	fmt.Scanln(&i)

 	for n := 1; n <= 10; n++ {
 		fmt.Println(i, " X ", n, " = ", i*n)
 	}

}

