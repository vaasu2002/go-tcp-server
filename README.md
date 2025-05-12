# Concurrent TCP Server in Go

This repository contains a concurrent TCP server implementation written in Go that demonstrates fundamental concepts of network programming and concurrency using the Go standard library.

## Overview

This TCP server:
- Listens for incoming TCP connections on port 1729
- Accepts multiple client connections concurrently
- Handles each connection in a separate goroutine
- Reads client requests using a buffer
- Processes each request (simulated with a 4-second delay)
- Responds with an HTTP 200 OK message
- Closes the connection after responding
- Continues accepting new connections without waiting for previous ones to complete

The implementation serves as a practical example of socket programming and concurrency in Go, illustrating core networking concepts.

## Concepts Covered

- **TCP (Transmission Control Protocol)**: A reliable protocol for machine-to-machine communication over networks
- **Sockets**: OS-provided abstractions that enable two-way communication
- **TCP Server**: A process that listens on a specific port and handles TCP connections
- **Listening & Accepting**: Fundamental operations in server-side socket programming
- **Buffers**: Temporary memory storage for data transfer between processes or devices
- **Blocking I/O Operations**: System calls that pause execution until data is available or operation completes
- **HTTP Response**: Simple implementation of HTTP protocol response formatting
- **Concurrency**: Handling multiple client connections simultaneously
- **Goroutines**: Go's lightweight threads for concurrent execution
- **Infinite Loop Server**: Continuous listening for and accepting new connections

## How It Works

1. The server creates a TCP listener on port 1729
2. It enters an infinite loop to continuously accept connections
3. For each incoming connection:
   - The server accepts the connection (blocking on the `Accept()` call)
   - Once established, it immediately spawns a new goroutine with `go ops(conn)`
   - The main thread returns to accepting new connections while the goroutine handles the client
4. Inside each goroutine (`ops` function):
   - A 1KB buffer is created to store incoming data
   - The server waits for client data (blocking on the `Read()` call)
   - After receiving data, the server simulates processing with a 4-second delay
   - The server sends back an HTTP 200 OK response with "Hello World!"
   - The connection is closed
5. Multiple clients can be served simultaneously without waiting for previous requests to complete

## Prerequisites

- Go 1.15 or higher

## Running the Server

1. Clone this repository
   ```
   git clone https://github.com/yourusername/go-tcp-server.git
   cd go-tcp-server
   ```

2. Run the server
   ```
   go run main.go
   ```

3. The server will display: "Waiting for connection to be established"

## Testing the Connection

You can test the server using various TCP clients. The server will respond with an HTTP 200 OK message after a 4-second simulated processing time:

### Testing Concurrency
To test the concurrency feature, open multiple terminal windows and run the client commands simultaneously. You'll see that the server handles all connections concurrently.

### Using curl (Recommended)
```
curl http://localhost:1729
```
You should see: `Hello World!` after a 4-second delay.

### Using Netcat
```
nc localhost 1729
```
After connecting, type any message and press Enter. You'll receive the HTTP response after 4 seconds.

### Using Load Testing Tools
You can also use tools like Apache Bench or wrk to test multiple concurrent connections:
```
ab -n 100 -c 10 http://localhost:1729/
```
This will send 100 requests with a concurrency level of 10.