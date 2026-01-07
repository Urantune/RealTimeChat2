package headlers

import (
	"net/http"

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
	defer conn.Close()

	roomService := services.NewChatRoomService()
	listRoom, err := roomService.GetAllChat()
	if err != nil {
		_ = conn.WriteJSON(gin.H{"error": err.Error()})
		return
	}

	_ = conn.WriteJSON(gin.H{"listRoom": listRoom})

}
