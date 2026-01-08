package headlers

import (
	"encoding/json"
	"net/http"
	"time"

	"RealTimeChatApplication/services"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ListRoom(c *gin.Context) {
	conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}

	done := make(chan struct{})

	go func() {
		defer close(done)
		for {
			if _, _, err := conn.ReadMessage(); err != nil {
				return
			}
		}
	}()

	roomService := services.NewChatRoomService()

	last := ""
	sendRooms := func() error {
		listRoom, err := roomService.GetAllChat()
		if err != nil {
			b, _ := json.Marshal(gin.H{"error": err.Error()})
			return conn.WriteMessage(websocket.TextMessage, b)
		}
		b, _ := json.Marshal(gin.H{"type": "room_list", "listRoom": listRoom})

		if string(b) == last {
			return nil
		}
		last = string(b)
		return conn.WriteMessage(websocket.TextMessage, b)
	}

	_ = sendRooms()

	ticker := time.NewTicker(1 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			_ = conn.Close()
			return
		case <-ticker.C:
			if err := sendRooms(); err != nil {
				_ = conn.Close()
				return
			}
		}
	}
}
