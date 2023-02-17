package webrtc

import (
	"log"
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
	Conn  *websocket.Conn
	Mutex sync.Mutex
}

func (t *ThreadSafeWriter) writeJSON(v interface{}) error {
	t.Mutex.Lock()
	defer t.Mutex.Unlock()
	return t.Conn.WriteJSON(v)
}

func (p *Peers) AddTrack(t *webrtc.TrackRemote) *webrtc.TrackLocalStaticRTP {
	p.ListLock.Lock()
	defer func() {
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

func (p *Peers) RemoveTrack(t *webrtc.TrackLocalStaticRTP) {

	p.ListLock.Lock()
	defer func() {
		p.ListLock.Unlock()
		p.SignalPeerConnections()
	}()

	delete(p.TrackLocals, t.ID())

}

func (p *Peers) SignalPeerConnections() {
	p.ListLock.Lock()
	defer func() {
		p.ListLock.Unlock()
		p.DispatchKeyFrame()
	}()

	attyemptSync := func() (tryAgain bool) {
		for i := range p.Connections {
			if p.Connections[i].PeerConnection.ConnectionState() == webrtc.PeerConnectionStateClosed {
				p.Connections = append(p.Connections[:i], p.Connections[i+1:]...)
				log.Println("a", p.Connections)
				return true
			}
			existingSenders := map[string]bool{}
				for _,sender := range p.Connections[i].PeerConnection.GetSenders(){
					if sender.Track() == nil {
						continue
					}

					existingSenders[sender.Track().ID()] = true
					if _, ok := p.TrackLocals[sender.Track().ID()] !ok {
						if err := p.Connections[i].PeerConnection.RemoveTrack(sender); err  != nil {
							return true
						}
					}
				}

				for _, receiver := range p.connections[i].Peerconnection.GetRecievers(){
					
					if receiver.Track() == nil {
						continue
					}

					existingSendore[receiver.Track().ID()] = true
				}


				for trackID := range p.TrackLocals{
					if _, ok := existingSenders[trackID]; !ok {
						if _,err := p.Connections[i].PeerConnection.AddTrack(p.TrackLocals[trackID]); err != nil {
							return true
						}
					}
				}
			}
		}
	}


func (p *Peers) DispatchKeyFrame() {

}

type websocketMessage struct {
	Event string `json:"event"`
	Data  string `json:"event"`
}
