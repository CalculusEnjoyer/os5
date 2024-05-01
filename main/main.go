package main

import (
	"fmt"
	"net"
	"strconv"
)

func main() {
	conn, err := net.Dial("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn.Close()

	x := 11

	_, err = conn.Write([]byte(fmt.Sprintf("f %d\n", x)))
	if err != nil {
		fmt.Println("Error sending data for calculating f(x):", err)
		return
	}

	result := make([]byte, 1024)
	n, err := conn.Read(result)
	if err != nil {
		fmt.Println("Error receiving result of calculating f(x):", err)
		return
	}
	fx := string(result[:n])

	conn1, err := net.Dial("tcp", "localhost:8081")
	if err != nil {
		fmt.Println("Error connecting:", err)
		return
	}
	defer conn1.Close()
	_, err = conn1.Write([]byte(fmt.Sprintf("g %d\n", x)))
	if err != nil {
		fmt.Println("Error sending data for calculating g(x):", err)
		return
	}

	n, err = conn1.Read(result)
	if err != nil {
		fmt.Println("Error receiving result of calculating g(x):", err)
		return
	}
	gx := string(result[:n])

	result = []byte(fmt.Sprintf("%s && %s", fx, gx))

	fint, _ := strconv.Atoi(fx)
	gint, _ := strconv.Atoi(gx)
	fmt.Println(fint)
	fmt.Println(gint)
	if fint != 0 && gint != 0 {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
