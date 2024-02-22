// Copyright 2015 The Gorilla WebSocket Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package main

import (
	"flag"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
	"google.golang.org/protobuf/proto"
	"vprenner.com/protos"
)

var addr = flag.String("addr", "localhost:9900", "http service address")

var upgrader = websocket.Upgrader{} // use default options

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			log.Println("read:", err)
			break
		}

		recvM := &protos.ReqGotoDSO{}
		proto.Unmarshal(message, recvM)
		log.Printf("recv: %s", recvM.String())
		response := protos.ComResponse{Code: 1}
		byteMsg, _ := proto.Marshal(&response)
		err = c.WriteMessage(mt, byteMsg)
		if err != nil {
			log.Println("write:", err)
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/", echo)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
