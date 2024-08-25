package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
	"time"
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
				function_call_res := handler()
				response = http.Response{
					StatusCode: 200,
					ProtoMajor: 1,
					ProtoMinor: 0,
					Body: io.NopCloser(strings.NewReader(function_call_res)),
				}
			} else {
				// Respond with 404 Not Found for other paths
				response = http.Response{
					StatusCode: 404,
					ProtoMajor: 1,
					ProtoMinor: 0,
					Body: io.NopCloser(strings.NewReader("404 Not Found\n")),
				}
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


func check_goroutine_mechanism () string{
	events := make(chan int)

	// goを外すとただの即時実行関数を同期的に実行しているだけになって停止する
	go func(){
		for i := 0;i<=5; i ++{
			// goroutine内でないと、受信者がまだ作られていないためデットロックを引き起こす
			events <- i
			time.Sleep(1 * time.Second)
		}
			close(events)
	}()

	for event := range events {
		fmt.Printf("Recieved event %d\n", event)
	}

	return "All events processed"
}