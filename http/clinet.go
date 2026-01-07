package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {

	b, _ := json.Marshal(map[string]string{"username": "urantune"})
	res, _ := http.Post("http://localhost:8080/login", "application/json", bytes.NewReader(b))
	var r map[string]string
	json.NewDecoder(res.Body).Decode(&r)
	token := r["token"]

	h := http.Header{}
	h.Set("Authorization", token)
	c, _, _ := websocket.DefaultDialer.Dial("ws://localhost:8080/listRoom", h)

	_, msg, _ := c.ReadMessage()
	fmt.Println(string(msg))
}
