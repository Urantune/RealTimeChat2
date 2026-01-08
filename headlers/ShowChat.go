package headlers

import (
	"RealTimeChatApplication/services"
	"RealTimeChatApplication/utils"
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var upgrade = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func ShowChat(c *gin.Context) {
	roomIdStr := c.Query("roomId")
	if roomIdStr == "" {
		c.AbortWithStatusJSON(400, gin.H{"error": "missing roomId"})
		return
	}
	roomId, err := strconv.Atoi(roomIdStr)
	if err != nil {
		c.AbortWithStatusJSON(400, gin.H{"error": "roomId must be number"})
		return
	}

	conne, err := upgrade.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		return
	}
	defer conne.Close()

	claimAny, ok := c.Get("user")
	if !ok {
		_ = conne.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"missing user in context"}`))
		return
	}
	claim, ok := claimAny.(*utils.Claims)
	if !ok {
		_ = conne.WriteMessage(websocket.TextMessage, []byte(`{"type":"error","error":"invalid user claim"}`))
		return
	}

	userService := services.NewUserService()
	u, err := userService.GetUserByUserName(claim.Username)
	if err != nil {
		b, _ := json.Marshal(gin.H{"type": "error", "error": err.Error()})
		_ = conne.WriteMessage(websocket.TextMessage, b)
		return
	}
	userId := int(u.ID)

	client := &Client{
		Conn:   conne,
		RoomID: roomId,
		UserID: userId,
		Send:   make(chan []byte, 32),
	}
	hub.join <- client
	defer func() { hub.leave <- client }()

	go func() {
		for msg := range client.Send {
			_ = conne.WriteMessage(websocket.TextMessage, msg)
		}
	}()

	chatService := services.NewChatMessengeService()

	sendChat := func() {
		listChat, err := chatService.GetChatByRoomIdCached(roomId)
		if err != nil {
			b, _ := json.Marshal(gin.H{"type": "error", "error": err.Error()})
			select {
			case client.Send <- b:
			default:
			}
			return
		}

		b, _ := json.Marshal(gin.H{
			"type":     "chat_list",
			"roomId":   roomId,
			"messages": listChat,
		})

		select {
		case client.Send <- b:
		default:
		}
	}

	sendChat()

	type InMsg struct {
		Type    string `json:"type"`
		Message string `json:"message"`
	}

	windowStart := time.Now()
	msgCount := 0

	for {
		_, data, err := conne.ReadMessage()
		if err != nil {
			return
		}

		var in InMsg
		if err := json.Unmarshal(data, &in); err != nil {
			continue
		}

		if in.Type == "get_list" {
			sendChat()
			continue
		}

		if in.Type != "" && in.Type != "send_message" {
			continue
		}

		if in.Message == "" {
			continue
		}

		now := time.Now()
		if now.Sub(windowStart) > 2*time.Second {
			windowStart = now
			msgCount = 0
		}
		if msgCount >= 5 {
			b, _ := json.Marshal(gin.H{"type": "error", "error": "rate limited"})
			select {
			case client.Send <- b:
			default:
			}
			continue
		}
		msgCount++

		_ = chatService.AddMessage(roomId, userId, in.Message)

		out, _ := json.Marshal(gin.H{
			"type":    "new_message",
			"roomId":  roomId,
			"userId":  userId,
			"message": in.Message,
		})
		hub.bcast <- broadcast{RoomID: roomId, Data: out}
	}
}
