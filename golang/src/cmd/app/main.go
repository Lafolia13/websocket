package main

import (
  "fmt"
  "net/http"

  "github.com/gorilla/websocket"
)

var clients = make(map[*websocket.Conn]bool)
var broadcast = make(chan Message)
var upgrader = websocket.Upgrader{}

type Message struct {
  Message string `json:"message"`
}

func HandleClients(w http.ResponseWriter, r *http.Request) {
  go broadcastMessagesToClients()
  
  upgrader.CheckOrigin = func(r *http.Request) bool {return true}
  websocket, err := upgrader.Upgrade(w, r, nil)
  if err != nil {
    fmt.Println("error upgrading GET request to a websocket::", err)
  }

  defer websocket.Close()

  clients[websocket] = true

  
  for {
    var message Message
    
    err := websocket.ReadJSON(&message)
    if err != nil {
      fmt.Printf("error occurred white reading message: %v", err)
      delete(clients, websocket)
      break
    }
    fmt.Println(message)
    
    broadcast <- message
  }
}

func main() {
  http.HandleFunc("/chat", HandleClients)
  err := http.ListenAndServe(":8888", nil)
  if err != nil {
    fmt.Println("error starting http server::", err)
    return
  }
}

func broadcastMessagesToClients() {
  for {
    message := <- broadcast
    for client := range clients {
      err := client.WriteJSON(message)
      if err != nil {
        fmt.Printf("error occurred white writing message to client: %v", err)
        client.Close()
        delete(clients, client)
      }
    }
  }
}