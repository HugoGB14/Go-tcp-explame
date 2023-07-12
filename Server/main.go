package main

import (
	"fmt"
	"net"
)

func main(){
	serverListen()
}

func serverListen(){
	server, err := net.Listen("tcp", ":4444")
	if err != nil{
		panic(err)
	}
	for {
		client ,err := server.Accept()
		if err != nil{
			fmt.Println(err)
			continue
		}
		go handleClient(client)
	}
}

func handleClient(client net.Conn){
	b := make([]byte, 2800)
	reqbs , err := client.Read(b)
	if err != nil{
		fmt.Println(err)
		return
	}
	
	fmt.Println("Msg: ", string(b[:reqbs]))


}