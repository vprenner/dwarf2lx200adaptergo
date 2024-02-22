package comm

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"net/url"
	"os"
	"os/signal"
	"time"
)

const DWARF_ADDRESS = "localhost:9900"

type DwarfConn struct {
	sendToDwarf      chan []byte
	receiveFromDwarf chan []byte
	connection       *websocket.Conn
}

func (conn *DwarfConn) connectToDwarf() <-chan bool {
	death := make(chan bool)
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	u := url.URL{Scheme: "ws", Host: DWARF_ADDRESS, Path: "/"}

	log.Printf("connecting to %s", u.String())
	nc, _, err := websocket.DefaultDialer.Dial(u.String(), nil)

	if err != nil {
		log.Fatal("dial:", err)
	}
	conn.connection = nc

	done := make(chan struct{})

	conn.sendToDwarf = make(chan []byte)
	conn.receiveFromDwarf = make(chan []byte)

	go func() {
		defer close(done)
		defer close(conn.receiveFromDwarf)
		conn.connection.SetReadLimit(1024)
		conn.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
		conn.connection.SetPongHandler(func(string) error {
			//fmt.Println("pong")
			conn.connection.SetReadDeadline(time.Now().Add(60 * time.Second))
			return nil
		})
		for {
			_, message, err := conn.connection.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				death <- true
				return
			}
			conn.receiveFromDwarf <- message
		}
	}()

	go func() {
		defer conn.connection.Close()
		defer close(conn.sendToDwarf)

		pingTicker := time.NewTicker(15 * time.Second)
		defer pingTicker.Stop()
		for {
			select {
			case <-pingTicker.C:
				conn.connection.SetWriteDeadline(time.Now().Add(10 * time.Second))
				//fmt.Println("ping")
				if err := conn.connection.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					death <- true
					return
				}
			case <-done:
				fmt.Println("done...")
				return
			case msgForDwarf := <-conn.sendToDwarf:
				err := conn.connection.WriteMessage(websocket.BinaryMessage, msgForDwarf)
				if err != nil {
					log.Println("write:", err)
					death <- true
					return
				}
			case <-interrupt:
				log.Println("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := conn.connection.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					log.Println("write close:", err)
					death <- false
					return
				}
				death <- false
				select {
				case <-done:
				}
				return
			}
		}
	}()

	return death

}
