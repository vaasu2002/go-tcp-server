# Simple TCP Server in Go

This repository contains a basic TCP server implementation written in Go that demonstrates the fundamental concepts of network programming using the Go standard library.

## Overview

This TCP server:
- Listens for incoming TCP connections on port 1729
- Accepts a single client connection
- Displays a confirmation message when a connection is established

The implementation serves as a foundational example of socket programming in Go, illustrating core networking concepts.

## Concepts Covered

- **TCP (Transmission Control Protocol)**: A reliable protocol for machine-to-machine communication over networks
- **Sockets**: OS-provided abstractions that enable two-way communication
- **TCP Server**: A process that listens on a specific port and handles TCP connections
- **Listening & Accepting**: Fundamental operations in server-side socket programming

## How It Works

1. The server creates a TCP listener on port 1729
2. It waits for an incoming client connection (blocking on the `Accept()` call)
3. Once a connection is established, it prints a confirmation message with connection details

## Prerequisites

- Go 1.15 or higher

## Running the Server

1. Clone this repository
   ```
   git clone https://github.com/vaasu2002/go-tcp-server.git
   cd go-tcp-server
   ```

2. Run the server
   ```
   go run main.go
   ```

3. The server will display: "Waiting for connection to be established"

## Testing the Connection

You can test the server using various TCP clients:

### Using Telnet
```
telnet localhost 1729
```

### Using Netcat
```
nc localhost 1729
```

### Using curl
```
curl telnet://localhost:1729
```
