package websocket

import (
  "code.google.com/p/go.net/websocket"
  
  "github.com/mashup-cms/mashup-api/model"

  "log"
)

type connection struct {
	websocket *websocket.Conn
	token string
	send chan *Message
}

func (conn *connection) reader() {
  var message Message
  for {
    log.Printf("waiting in read")
    if err := websocket.JSON.Receive(conn.websocket, &message); err != nil {
      break
    }
    log.Printf("%s", message)
  }
  conn.websocket.Close()
}

func (conn *connection) writer() {
  for message := range conn.send {
    err := websocket.JSON.Send(conn.websocket, message)
    if err != nil {
      break
    }
  }
  conn.websocket.Close()
}
  
func Handler() websocket.Handler {
  return websocket.Handler(sockServer)
}

func sockServer(ws *websocket.Conn) {
  queryKeys := ws.Request().URL.Query()
  if mashupKey,ok := queryKeys["key"]; ok {
    log.Printf("%s", mashupKey[0])
    token := model.FindAccessToken(mashupKey[0])
    if token != nil {
      client := &connection{ws, mashupKey[0], make(chan *Message, 256)}
      Hub.register <- client
      defer func() { Hub.unregister <- client }()
      go client.writer()
      client.reader()
    }
  }
}