package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-realtimemap/api/admin/v1"
)

func (s *AdminService) SetWebsocketServer(ws *websocket.Server) {
	s.ws = ws
}

func (s *AdminService) OnWebsocketMessage(message *websocket.Message) (*websocket.Message, error) {
	fmt.Println(" Payload: ", string(message.Body))

	var proto v1.WebsocketProto

	if err := json.Unmarshal(message.Body, &proto); err != nil {
		s.log.Error("Error unmarshalling proto json %v", err)
		return nil, nil
	}

	switch proto.EventId {
	case "viewport":
		var msg v1.Viewport
		if err := json.Unmarshal([]byte(proto.Payload), &msg); err != nil {
			s.log.Error("Error unmarshalling payload json %v", err)
			return nil, nil
		}
	}

	return nil, nil
}

func (s *AdminService) BroadcastToWebsocketClient(eventId string, payload interface{}) {
	bufPayload, _ := json.Marshal(&payload)

	var proto v1.WebsocketProto
	proto.EventId = eventId
	proto.Payload = string(bufPayload)

	bufProto, _ := json.Marshal(&proto)

	var msg websocket.Message
	msg.Body = bufProto

	s.ws.Broadcast(&msg)
}
