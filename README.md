# Simple TCP Server in Go

This repository contains a basic TCP server implementation written in Go that demonstrates the fundamental concepts of network programming using the Go standard library.

## Overview

This TCP server:
- Listens for incoming TCP connections on port 1729
- Accepts a client connection
- Reads client requests using a buffer
- Processes the request (simulated with a 1-second delay)
- Responds with an HTTP 200 OK message
- Closes the connection after responding

The implementation serves as a foundational example of socket programming in Go, illustrating core networking concepts.

## Concepts Covered

- **TCP (Transmission Control Protocol)**: A reliable protocol for machine-to-machine communication over networks
- **Sockets**: OS-provided abstractions that enable two-way communication
- **TCP Server**: A process that listens on a specific port and handles TCP connections
- **Listening & Accepting**: Fundamental operations in server-side socket programming
- **Buffers**: Temporary memory storage for data transfer between processes or devices
- **Blocking I/O Operations**: System calls that pause execution until data is available or operation completes
- **HTTP Response**: Simple implementation of HTTP protocol response formatting

## How It Works

1. The server creates a TCP listener on port 1729
2. It waits for an incoming client connection (blocking on the `Accept()` call)
3. When a connection is established, it calls the `ops` function to handle the connection
4. Inside `ops`:
   - A 1KB buffer is created to store incoming data
   - The server waits for client data (blocking on the `Read()` call)
   - After receiving data, the server simulates processing with a 1-second delay
   - The server sends back an HTTP 200 OK response with "Hello World!"
   - The connection is closed
5. The server prints a confirmation message with connection details

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

You can test the server using various TCP clients. The server will respond with an HTTP 200 OK message:

### Using curl (Recommended)
```
curl http://localhost:1729
```
You should see: `Hello World!`

### Using Telnet
```
telnet localhost 1729
```
After connecting, type any message and press Enter. You'll receive the HTTP response.

### Using Netcat
```
nc localhost 1729
```
After connecting, type any message and press Enter. You'll receive the HTTP response.

### Using a Web Browser
Open your browser and navigate to:
```
http://localhost:1729
```

## Key Components Explained

### Buffers
The server uses a buffer (`buf := make([]byte, 1024)`) to temporarily store incoming data from clients:
- **Purpose**: Provides a memory area to hold client requests for processing
- **Size**: 1KB (1024 bytes) is allocated to store incoming data
- **Operation**: The `conn.Read(buf)` call fills this buffer with data from the client

### Blocking I/O Operations
The server uses several blocking operations:
- **`listener.Accept()`**: Blocks until a client connects
- **`conn.Read(buf)`**: Blocks until data is received from the client
- These operations pause execution of the program until the event they're waiting for occurs

### Process Simulation
The server includes a simulated processing delay:
- **`time.Sleep(1*time.Second)`**: Simulates server-side processing time
- In a real application, this would be replaced with actual request handling logic