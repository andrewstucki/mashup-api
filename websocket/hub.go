package websocket

type hub struct {
	// Registered connections.
	connections map[*connection]bool
 
	// Register requests from the connections.
	register chan *connection
 
	// Unregister requests from connections.
	unregister chan *connection
}
 
var Hub = hub{
  register:    make(chan *connection),
	unregister:  make(chan *connection),
	connections: make(map[*connection]bool),
}
 
func (h *hub) Run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			delete(h.connections, c)
			close(c.send)
		}
	}
}

func (h *hub) SendMessage(clientKey string, message *Message) {
  for c := range h.connections {
    if c.token == clientKey {
      c.send <- message
    }
  }
}