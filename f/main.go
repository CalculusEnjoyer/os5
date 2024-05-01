package main

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

func main() {
	ln, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer ln.Close()
	fmt.Println("Server for f(x) is listening on port 8080")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buf := make([]byte, 1024)
	n, err := conn.Read(buf)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	data := string(buf[:n])
	parts := strings.Fields(data)
	if len(parts) != 2 {
		fmt.Println("Invalid request:", data)
		return
	}

	x, err := strconv.Atoi(parts[1])
	if err != nil {
		fmt.Println("Invalid value for x:", parts[1])
		return
	}
	fx := calculateFx(x)

	_, err = conn.Write([]byte(fmt.Sprintf("%d", fx)))
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}

func calculateFx(x int) int {
	return x % 2
}
