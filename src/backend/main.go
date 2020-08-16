package main

import (
	"fmt"
	"net"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/gorilla/websocket"
)

// ConnectionsPool - ...
type ConnectionsPool struct {
	mu   *sync.RWMutex
	Cons map[net.Addr]*websocket.Conn
}

// Add - ...
func (cp *ConnectionsPool) Add(c *websocket.Conn) {
	cp.mu.Lock()
	cp.Cons[c.RemoteAddr()] = c
	cp.mu.Unlock()
}

// Remove - ...
func (cp *ConnectionsPool) Remove(addr net.Addr) {
	cp.mu.Lock()
	if conn, ok := cp.Cons[addr]; ok {
		conn.Close()
		delete(cp.Cons, addr)
	}
	cp.mu.Unlock()
}

// Get - ...
func (cp *ConnectionsPool) Get(addr net.Addr) *websocket.Conn {
	cp.mu.RLock()
	defer cp.mu.RUnlock()
	return cp.Cons[addr]
}

// NewConnectionsPool - ...
func NewConnectionsPool() *ConnectionsPool {
	return &ConnectionsPool{
		mu:   new(sync.RWMutex),
		Cons: make(map[net.Addr]*websocket.Conn)}
}

var (
	connections *ConnectionsPool
	data        = make([]map[string]string, 0)
	upgrade     = websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }, ReadBufferSize: 1024, WriteBufferSize: 1024}
)

func init() {
	connections = NewConnectionsPool()
}

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
	connections.Add(conn)
	var wsHandler = func(addr net.Addr) {
		defer connections.Remove(addr)
		notify()
		for {
			select {
			default:
				var rec map[string]string
				if err := connections.Get(addr).ReadJSON(&rec); websocket.IsCloseError(err, websocket.CloseGoingAway, websocket.CloseNoStatusReceived) {
					return
				}
				if val, ok := rec["index"]; ok {
					n, _ := strconv.Atoi(val)
					data = RemoveIndex(data, n)
				} else {
					data = append(data, rec)
				}
				notify()
			case <-r.Context().Done():
				return
			}
		}
	}
	wsHandler(conn.RemoteAddr())
}

func timer(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("upgrade:", err)
		return
	}
	connections.Add(conn)
	var wsHandler = func(addr net.Addr) {
		defer connections.Remove(addr)
		ticker := time.NewTimer(time.Second)
		for {
			select {
			case t := <-ticker.C:
				c := connections.Get(addr)
				if err := c.WriteJSON(map[string]interface{}{"tick": t.Unix()}); err != nil {
					return
				}
				ticker.Reset(time.Second)
			case <-r.Context().Done():
				ticker.Stop()
				return
			}
		}
	}
	wsHandler(conn.RemoteAddr())
}

func notify() {
	for addr := range connections.Cons {
		go func(addr net.Addr) {
			if err := connections.Get(addr).WriteJSON(map[string]interface{}{"table": data}); err != nil {
				fmt.Println("notify:", err)
				return
			}
		}(addr)
	}
}

func main() {
	http.Handle("/", http.FileServer(http.Dir("static")))
	http.HandleFunc("/table", greet)
	http.HandleFunc("/timer", timer)
	http.ListenAndServe("0.0.0.0:9999", nil)
}
