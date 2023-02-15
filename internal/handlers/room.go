package handlers

import(
	"fmt"
	"os"
	"time"
	w "videochat/pkg/webrtc"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	return c.Redirect(fmt.Sprintf("/room/%s",guuid.New().String()))
}

func Room (c *fiber.Ctx) error{

	uuid := c.Params("uuid")
	if uuid == "" {
		c.Status(400)
		return nil
	}

	uuid,suuid,_ := c.createOrGetRoom(uuid)
}


func RoomWebsocket(c *websocket.Conn) {
	uuid := c.Params("uuid")
	if uuid == "" {
		return
	}

	_,_, room := createOrGetRoom(uuid)

	w.RoomConn

}	


func createOrGetRoom(uuid string) (string, string,w.Room) {
	
}

func RoomViewwerWebsocket(c *websocket.Conn) {

}

func roomViewwerConn(c *websocket.Conn, p *w.Peers){

}

type websocketMessage struct {
	Event string `json:"event"`
	Data string `json:"data"`
}