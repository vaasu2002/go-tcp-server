package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// TCP    - most reliable way for machines to communicate over a network
// Socket - abstraction(logical entities which wraps entier communication specifics) provided by OS to enable communication(end points for two way communication)
// TCP Server - a process that runs in a machine that listens to a port that understands TCP
// any machine that needs to talk with a server needs to connect to the port and establish connection

// Step1 - listening on a port
// Step2 - wait for client to connection(we will use accept system call which is a blocking system call so server could not proced until some client connects)
// Step3 - Read the request and send the response (read and write system call are blocking)
// Step4 - Continuouly waiting for client to connect with the tcp server(will use infinite loop here)
// We can currently process only one request at a time
// Step5 - Handling multiple request concurrently 
// Running in parallel
// Once a client connects, fork a thread to process and respend, main thread should came back to accept as soon as possible

// A client connects → the accept() system call returns a new socket for communication with that client.
// Immediately after that we will creare a new Thread to (read clients req, process it, send response and close the connection(socket))
// Meanwhile, the main thread (the one that called accept()) should not handle the client directly. Instead, it goes right back to calling accept() again to wait for the next client.


func ops(conn net.Conn){
	// when a client sends data to your server, the operating system places this data in a system buffer. 
	// then we copy data system buffer to application buffer
	// once we have data in out(applciation) buffer, we can examine it, modify it etc.
	buf := make([]byte,1024) // creating byte buffer. 1KB is size of buffer
	fmt.Println("Reading request");
	_,err := conn.Read(buf) // waiting until client sends us the request 
	if(err!=nil){
		log.Fatalln(err)
	}
	fmt.Println("Processing!!");
	// doing some processs
	time.Sleep(4*time.Second)
	fmt.Println("Writing response");
	conn.Write([]byte("HTTP/1.1 200 OK\r\n\r\nHello World!\r\n"));
	fmt.Println("Closing current connection");
	conn.Close()
}
func main() {
	// listening for tcp protocol on port :1729
	// this tcp server(which is a process) is now reserving a port 1729
	listner, err := net.Listen("tcp",":1729") 
	if(err!=nil){
		log.Fatalln(err)
	}
	for{
		fmt.Println("Waiting for connection to be established")
		conn,err := listner.Accept(); //  waiting for some client to connect
		if(err!=nil){
			log.Fatalln(err)
		}
		fmt.Println("Connection established");
		go ops(conn)
	}
}
 