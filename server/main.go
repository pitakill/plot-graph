package main

import (
	"encoding/json"
	"flag"
	"log"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

var addr = flag.String("addr", "localhost:8080", "http service address")

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		host := strings.Split(r.Host, ":")[0]

		if r.Host == "localhost" || host == "localhost" {
			return true
		}

		return false
	},
}

func connections(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()

	// Parse json to struct
	conns := new(Connections)

	if err := json.Unmarshal(jsonConnections, &conns); err != nil {
		log.Println(err)
	}

	var ws sync.WaitGroup
	var mutex sync.Mutex

	ws.Add(2)
	go func(ws *sync.WaitGroup) {
		i := 1
		for id, node := range conns.Nodes {
			info, err := json.Marshal(map[string]interface{}{"id": id, "details": node, "type": "node"})
			if err != nil {
				log.Println("marshaling: ", err)
			}
			time.Sleep(time.Duration(i*200) * time.Millisecond)
			mutex.Lock()
			if err := c.WriteMessage(websocket.TextMessage, info); err != nil {
				log.Println("write: ", err)
			}
			mutex.Unlock()
			i++
		}
		ws.Done()
	}(&ws)

	go func(ws *sync.WaitGroup) {
		for i, edge := range conns.Edges {
			info, err := json.Marshal(map[string]interface{}{"source": edge.Source, "destination": edge.Destination, "type": "edge"})
			if err != nil {
				log.Println("marshaling: ", err)
			}
			time.Sleep(time.Duration(i*200) * time.Millisecond)
			mutex.Lock()
			if err := c.WriteMessage(websocket.TextMessage, info); err != nil {
				log.Println("write: ", err)
			}
			mutex.Unlock()
		}
		ws.Done()
	}(&ws)
	ws.Wait()
}

func resources(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade: ", err)
		return
	}
	defer c.Close()
	for i, message := range jsonResourcesSlice {
		time.Sleep(time.Duration(i+1*200) * time.Millisecond)
		if err := c.WriteMessage(websocket.TextMessage, message); err != nil {
			log.Println("write: ", err)
		}
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/resources", resources)
	http.HandleFunc("/connections", connections)
	log.Fatal(http.ListenAndServe(*addr, nil))
}
