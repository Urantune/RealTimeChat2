# RealTime Chat Application

Real-time chat backend built with Go and Gin using WebSocket, JWT authentication, PostgreSQL and Redis.

---

## Tech

Language: Go 1.25  
Framework: Gin  
Realtime: WebSocket  
Database: PostgreSQL  
Cache / Presence: Redis  
Auth: JWT  

---

## Features

- User login with JWT
- Create and join chat rooms
- Real-time messaging with WebSocket
- Store and load chat history
- Track online / offline users
- Secure APIs with middleware

---

## API

POST http://localhost:8080/login  
{
  "email": "admin",
  "password": "123456"
}

GET http://localhost:8080/listRoom  
Header: Authorization: Bearer TOKEN

WebSocket connect  
ws://localhost:8080/ws?roomId=1  
Header: Authorization: Bearer TOKEN

Send message (WebSocket)  
{
  "type": "send_message",
  "message": "Hello everyone"
}

GET http://localhost:8080/chat/history?roomId=1  
Header: Authorization: Bearer TOKEN

---

## Project Structure

RealTimeChatApplication  
main.go  
headlers/  
middleware/  
services/  
repository/  
models/  
routes/  
utils/  

---

## Run

Start PostgreSQL and Redis

go mod tidy  
go run main.go  

Server:  
http://localhost:8080  

---

## Flow

1. User login to get JWT  
2. Client connects WebSocket with roomId  
3. Messages are broadcast in the room  
4. Redis tracks online users  
5. PostgreSQL stores messages  

---

## Use

- Real-time chat
- Team chat
- Live support
- Game chat
- Internal messaging
