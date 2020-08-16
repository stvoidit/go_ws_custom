package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/websocket"
)

var (
	connections = make(map[net.Addr]*websocket.Conn)
	data        = make([]map[string]string, 0)
	upgrade     = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }, ReadBufferSize: 1024, WriteBufferSize: 1024}
)

// RemoveIndex - ...
func RemoveIndex(s []map[string]string, index int) []map[string]string {
	return append(s[:index], s[index+1:]...)
}

func greet(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	defer func() { delete(connections, conn.RemoteAddr()) }()
	connections[conn.RemoteAddr()] = conn
	var wsHandler = func(c *websocket.Conn) {
		notify()
		for {
			var rec map[string]string
			if err := c.ReadJSON(&rec); err != nil {
				return
			}
			if val, ok := rec["index"]; ok {
				n, _ := strconv.Atoi(val)
				data = RemoveIndex(data, n)
			} else {
				data = append(data, rec)
			}
			notify()
		}
	}
	wsHandler(conn)
}

func timer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	defer conn.Close()
	defer func() { delete(connections, conn.RemoteAddr()) }()
	connections[conn.RemoteAddr()] = conn
	ticker := time.NewTimer(time.Second)
	for {
		select {
		case t := <-ticker.C:
			if err := conn.WriteJSON(map[string]interface{}{"tick": t.Unix()}); err != nil {
				fmt.Println("ticker:", err)
				return
			}
			ticker.Reset(time.Second)
		case <-r.Context().Done():
			return
		}
	}
}

func notify() {
	for addr := range connections {
		go func(conn *websocket.Conn) {
			if err := conn.WriteJSON(map[string]interface{}{"table": data}); err != nil {
				fmt.Println("notify:", err)
				return
			}
		}(connections[addr])
	}
}

func main() {
	http.HandleFunc("/table", greet)
	http.HandleFunc("/timer", timer)
	http.ListenAndServe("0.0.0.0:9999", nil)
}
