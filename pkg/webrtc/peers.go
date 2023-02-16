package webrtc

import (
	"sync"
	"videochat/pkg/chat"

	"github.com/gofiber/websocket/v2"
)

type Room struct {
	Peers *Peers
	Hub   *chat.Hub
}

type Peers struct {
	ListLock    sync.RWMutex
	Connections []PeerConnectionState
	TrackLocals map[string]*webrtc.TrackLocalStaticRTP
}

type PeerConnectionState struct {
	PeerConnection *webrtc.PeerConnection
	websocket      *ThreadSafeWriter
}

type ThreadSafeWriter struct {
	Conn *websocket.Conn
	Mutex sync.Mutex
}

func (t *ThreadSafeWriter) writeJSON (v interface{}) error{
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack (t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP{
	p.ListLock.Lock()
	defer func(){
		p.ListLock.Lock()
		p.SignalPeerConnections()
	
	}()

	trackLocal, err := webrtc.NewTrackLocalStaticRTP(t.Codec().RTPCodecCapability, t.ID(), t.StreamID())
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	p.TrackLocals[t.ID()] = p.TrackLocal
	return trackLocal
}

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP){

}

func(p *Peers)SignalPeerConnections(){

}


func (p *Peers) DispatchKeyFrame() {

}

type websocketMessage struct {
	Event string `json:"event"`
	Data string `json:"event"`
}
