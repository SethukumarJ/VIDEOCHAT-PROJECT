package chat

type Hub struct {
	clients map[*Client]bool
	brockcast chan []byte
	register chan *Client
	unregister chan *Client
}