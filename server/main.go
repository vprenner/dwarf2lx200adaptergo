package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"vprenner.com/comm"
)

const LISTEN_PORT = "9999"

// Server ...
type Server struct {
	host string
	port string
}

// Client ...
type Client struct {
	conn net.Conn
}

// Run ...
func (server *Server) Run() {
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", server.host, server.port))
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatal(err)
		}

		client := &Client{
			conn: conn,
		}
		go client.handleRequest()
	}
}

func (client *Client) handleRequest() {
	reader := bufio.NewReader(client.conn)
	for {
		message, err := reader.ReadString('#')
		if err != nil {
			client.conn.Close()
			return
		}
		msg := strings.Trim(strings.TrimSpace(message), "#")
		if msg != ":GR" && msg != ":GD" {
			fmt.Printf("Message incoming: '%s'\n", msg)
		}

		// do stuff
		// send result
		result := comm.Send(msg)

		client.conn.Write([]byte(result))
	}
}

func main() {
	server := Server{host: "0.0.0.0", port: LISTEN_PORT}
	server.Run()
	fmt.Println("ALLDONEFOLKS")
}
