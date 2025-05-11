package main

import (
	"fmt"
	"log"
	"net"
)
// TCP    - most reliable way for machines to communicate over a network
// Socket - abstraction(logical entities which wraps entier communication specifics) provided by OS to enable communication(end points for two way communication)
// TCP Server - a process that runs in a machine that listens to a port that understands TCP
// any machine that needs to talk with a server needs to connect to the port and establish connection

// Step1 - listening on a port
// Step2 - wait for client to connection(we will use accept system call which is a blocking system call so server could not proced until some client connects)

func main() {
	// listening for tcp protocol on port :1729
	// this tcp server(which is a process) is now reserving a port 1729
	listner, err := net.Listen("tcp",":1729") 
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println("Waiting for connection to be established")
	conn,err := listner.Accept(); //  waiting for some client to connect
	if(err!=nil){
		log.Fatal(err)
	}
	fmt.Println("hello world ",conn)
}
