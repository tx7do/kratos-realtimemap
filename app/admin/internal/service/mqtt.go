package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-realtimemap/api/hfp"
	"strings"
)

func (s *AdminService) SetMqttBroker(b broker.Broker) {
	s.mb = b
}

func (s *AdminService) TransitPostTelemetry(event broker.Event) error {
	fmt.Println("TransitPostTelemetry() Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))

	topicParts := strings.Split(event.Topic(), "/")

	var msg hfp.Event

	if err := json.Unmarshal(event.Message().Body, &event); err != nil {
		s.log.Error("Error unmarshalling json %v", err)
	} else {
		msg.OperatorId = topicParts[7]
		msg.VehicleId = topicParts[7] + "." + topicParts[8]
		//onEvent(&event)
	}

	return nil
}
