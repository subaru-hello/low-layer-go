package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}

	fmt.Println("Server is running at localhost:8888")

	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}

		go func() {
			defer conn.Close()
			fmt.Printf("Accept %v\n", conn.RemoteAddr())

			// Read the request
			request, err := http.ReadRequest(bufio.NewReader(conn))
			if err != nil {
				fmt.Println("Error reading request:", err)
				return
			}

			// Dump the request for debugging purposes
			dump, err := httputil.DumpRequest(request, true)
			if err != nil {
				fmt.Println("Error dumping request:", err)
				return
			}
			fmt.Println(string(dump))

			// Check the request path
			var response http.Response
            if handler,ok := routes[request.URL.Path]; ok{
				
			}

			// Write the response
			err = response.Write(conn)
			if err != nil {
				fmt.Println("Error writing response:", err)
				return
			}
		}()
	}
}
