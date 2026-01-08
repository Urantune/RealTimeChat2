package headlers

import (
	"encoding/json"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type Client struct {
	Conn   *websocket.Conn
	RoomID int
	UserID int
	Send   chan []byte
}

type broadcast struct {
	RoomID int
	Data   []byte
}

type Hub struct {
	rooms  map[int]map[*Client]struct{}
	online map[int]map[int]int
	join   chan *Client
	leave  chan *Client
	bcast  chan broadcast
}

var hub = NewHub()

func NewHub() *Hub {
	h := &Hub{
		rooms:  make(map[int]map[*Client]struct{}),
		online: make(map[int]map[int]int),
		join:   make(chan *Client),
		leave:  make(chan *Client),
		bcast:  make(chan broadcast, 32),
	}
	go h.Run()
	return h
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.join:
			if h.rooms[c.RoomID] == nil {
				h.rooms[c.RoomID] = make(map[*Client]struct{})
			}
			h.rooms[c.RoomID][c] = struct{}{}

			if h.online[c.RoomID] == nil {
				h.online[c.RoomID] = make(map[int]int)
			}
			h.online[c.RoomID][c.UserID]++

			h.broadcastPresence(c.RoomID)

		case c := <-h.leave:
			if m := h.rooms[c.RoomID]; m != nil {
				delete(m, c)
				close(c.Send)
				if len(m) == 0 {
					delete(h.rooms, c.RoomID)
				}
			}

			if om := h.online[c.RoomID]; om != nil {
				om[c.UserID]--
				if om[c.UserID] <= 0 {
					delete(om, c.UserID)
				}
				if len(om) == 0 {
					delete(h.online, c.RoomID)
				}
			}

			h.broadcastPresence(c.RoomID)

		case msg := <-h.bcast:
			if m := h.rooms[msg.RoomID]; m != nil {
				for c := range m {
					select {
					case c.Send <- msg.Data:
					default:
						delete(m, c)
						close(c.Send)
					}
				}
			}
		}
	}
}

func (h *Hub) broadcastPresence(roomID int) {
	onlineUsers := make([]int, 0)
	if m := h.online[roomID]; m != nil {
		for uid, cnt := range m {
			if cnt > 0 {
				onlineUsers = append(onlineUsers, uid)
			}
		}
	}

	payload, _ := json.Marshal(gin.H{
		"type":   "presence",
		"roomId": roomID,
		"online": onlineUsers,
	})

	if clients := h.rooms[roomID]; clients != nil {
		for c := range clients {
			select {
			case c.Send <- payload:
			default:
			}
		}
	}
}
