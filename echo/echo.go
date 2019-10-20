package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

// Compile binary version:
//  % go build
// This will create binary file `echo` in current folder.
//
// Run the server in one terminal window:
//  % ./echo -server
//
// Run the client in another terminal window.
//  % ./echo -client
//
// Run one more client in a third terminal window.
//  % ./echo -client
//
// Homework: Make it work so that both clients can connect at the same time.

func main() {
	var (
		server   = flag.Bool("server", false, "Start echo server if true")
		endpoint = flag.String("endpoint", "localhost:12100", "Endpoint on which server runs or to which the client connects")
	)
	flag.Parse()
	if *server {
		serverLoop(*endpoint)
	} else {
		clientLoop(*endpoint)
	}
}

func serverLoop(endpoint string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to resolve endpoint addr %s", endpoint)
	}
	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatalf("failed to create listener for addr %v", tcpAddr)
	}
	fmt.Println("Waiting for clients...")
	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		handleClient(conn)
	}
}

func handleClient(conn net.Conn) {
	defer conn.Close()
	var buf [512]byte
	for {
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "Connection closed by %s\n", conn.RemoteAddr())
			} else {
				fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
			}
			return
		}
		_, err = conn.Write(buf[0:n])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
			return
		}
		fmt.Println(string(buf[0:n]))
	}
}

func clientLoop(endpoint string) {
	tcpAddr, err := net.ResolveTCPAddr("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to resolve endpoint addr %s", endpoint)
	}
	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		log.Fatalf("failed to dial TCP addr %v", tcpAddr)
	}
	defer conn.Close()

	var buf [512]byte
	fmt.Print("Enter: ")
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		s := scanner.Bytes()
		if string(s) == "exit" {
			break
		}
		_, err := conn.Write(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Write error: %v\n", err)
			return
		}
		n, err := conn.Read(buf[0:])
		if err != nil {
			if err == io.EOF {
				fmt.Fprintf(os.Stderr, "Connection closed by %s\n", conn.RemoteAddr())
			} else {
				fmt.Fprintf(os.Stderr, "Read error: %v\n", err)
			}
			return
		}
		fmt.Printf("Server response: %s\n", string(buf[0:n]))
		fmt.Print("Enter: ")
	}
}
