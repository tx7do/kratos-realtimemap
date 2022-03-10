package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	"kratos-realtimemap/api/hfp"
	"strings"
)

func (s *TransportMqttService) SetMqttBroker(b broker.Broker) {
	s.mb = b
}

func (s *TransportMqttService) TransitPostTelemetry(event broker.Event) error {
	fmt.Println("TransitPostTelemetry() Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))

	var msg hfp.Event

	//0/1       /2        /3             /4              /5           /6               /7            /8               /9         /10            /11        /12          /13         /14             /15       /16
	// /<prefix>/<version>/<journey_type>/<temporal_type>/<event_type>/<transport_mode>/<operator_id>/<vehicle_number>/<route_id>/<direction_id>/<headsign>/<start_time>/<next_stop>/<geohash_level>/<geohash>/<sid>/#
	topicParts := strings.Split(event.Topic(), "/")

	if err := json.Unmarshal(event.Message().Body, &event); err != nil {
		s.log.Error("Error unmarshalling json %v", err)
	} else {
		msg.OperatorId = topicParts[7]
		msg.VehicleId = topicParts[7] + "." + topicParts[8]
	}

	return nil
}
