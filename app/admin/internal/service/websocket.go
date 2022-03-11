package service

import (
	"fmt"
	"github.com/tx7do/kratos-transport/transport/websocket"
)

func (s *AdminService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *AdminService) OnWebsocketMessage(message *websocket.Message) (*websocket.Message, error) {
	fmt.Println(" Payload: ", string(message.Body))

	return nil, nil
}
