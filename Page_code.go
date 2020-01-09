package main

import (
    "fmt"
    "net"
	"strings"
)

// Split port and Host
func main() {
   // host, port, err := net.SplitHostPort("127.0.0.1:5432")
    //fmt.Println(host, port, err)
	username := strings.Split("vkalipat@gmail.com", "@")
 	fmt.Printf("%q\n", username[0])
 }