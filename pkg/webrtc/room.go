package webrtc

import (
	"log"
	"sync"

	"github.com/fasthttp/websocket"
)

func RoomConn(c *websocket.Conn, p *Peers) {

	var config webrtc.Configuration

	peerConnection, err := webrtc.NewPeerConnection(config)
	if err != nil {
		log.Print(err)
		return
	}
	newPeer := PeerConnectionState{
		PeerConnection: peerConnection,
		WebSocket:      &ThreadSafWriter{},
		Conn:           c,
		Mutex:          sync.Mutex{},
	}
}
