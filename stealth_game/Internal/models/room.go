package models

import (
	"sync"

	"github.com/gorilla/websocket"
)

type Room struct {
	Clients map[*websocket.Conn]*Player
	Mutex   sync.Mutex
}

func NewRoom() *Room {
	return &Room{
		Clients: make(map[*websocket.Conn]*Player),
	}
}

// プレイヤー追加用メソッド
func (r *Room) AddPlayer(conn *websocket.Conn, player *Player) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	r.Clients[conn] = player
}

// プレイヤー削除用メソッド
func (r *Room) RemovePlayer(conn *websocket.Conn) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()
	delete(r.Clients, conn)
}

func (r *Room) BroadcastMessage(message map[string]interface{}) {
	r.Mutex.Lock()
	defer r.Mutex.Unlock()

	for conn := range r.Clients {
		err := conn.WriteJSON(message)
		if err != nil {
			conn.Close()
			delete(r.Clients, conn)
		}
	}
}
