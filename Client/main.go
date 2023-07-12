package main

import (
	"net"
	"os"
)

func main(){
	client()
}

func client(){
	client, err := net.Dial("tcp", ":4444")
	if err != nil{
		panic(err)
	}

	msg := os.Args[1]

	client.Write([]byte(msg))
}