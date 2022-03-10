package service

import (
	"encoding/json"
	"fmt"
	"github.com/tx7do/kratos-transport/broker"
	v1 "kratos-realtimemap/api/admin/v1"
	"kratos-realtimemap/api/hfp"
	"kratos-realtimemap/app/admin/internal/pkg/data"
	"strings"
)

const MaxPositionHistory = 200

func (s *AdminService) SetMqttBroker(b broker.Broker) {
	s.mb = b
}

func (s *AdminService) TransitPostTelemetry(event broker.Event) error {
	//fmt.Println("Topic: ", event.Topic(), " Payload: ", string(event.Message().Body))

	//0/1       /2        /3             /4              /5           /6               /7            /8               /9         /10            /11        /12          /13         /14             /15       /16
	// /<prefix>/<version>/<journey_type>/<temporal_type>/<event_type>/<transport_mode>/<operator_id>/<vehicle_number>/<route_id>/<direction_id>/<headsign>/<start_time>/<next_stop>/<geohash_level>/<geohash>/<sid>/#
	topicParts := strings.Split(event.Topic(), "/")

	var msg hfp.Event

	if err := json.Unmarshal(event.Message().Body, &msg); err != nil {
		s.log.Error("Error unmarshalling json %v", err)
	} else {

		vehicleNumber := topicParts[8]
		operatorId := topicParts[7]
		eventType := topicParts[5]     // vp, due, arr, dep, ars, pde, pas, wait, doo, doc, tlr, tla, da, dout, ba, bout, vja, vjout
		transportMode := topicParts[6] // bus, tram, train, ferry, metro, ubus

		msg.OperatorId = operatorId
		msg.VehicleId = operatorId + "." + vehicleNumber

		position := MapToPosition(&msg)
		if position != nil {
			s.UpdatePositionHistory(position)
			s.UpdateGeofence(position)
		}

		fmt.Println("事件类型: ", eventType, " 交通工具类型: ", transportMode, " 司机ID: ", msg.OperatorId, " 车辆ID: ", msg.VehicleId)

		//fmt.Printf("%v\n", msg.VehiclePosition)
	}

	return nil
}

func (s *AdminService) UpdateGeofence(position *v1.Position) {
	for _, org := range data.AllOrganizations {
		if len(org.Geofences) > 0 {
		}

		for _, geofence := range org.Geofences {
			_, vehicleIsInZone := geofence.VehiclesInZone[position.VehicleId]
			if geofence.IncludesPosition(position.Latitude, position.Longitude) {
				if !vehicleIsInZone {
					geofence.VehiclesInZone[position.VehicleId] = struct{}{}
				}
			} else {
				if vehicleIsInZone {
					delete(geofence.VehiclesInZone, position.VehicleId)
				}
			}
		}
	}
}

func (s *AdminService) UpdatePositionHistory(position *v1.Position) {
	_, ok := s.positionHistory[position.VehicleId]
	if !ok {
		his := make([]*v1.Position, 0, MaxPositionHistory)
		s.positionHistory[position.VehicleId] = his
	}

	if len(s.positionHistory[position.VehicleId]) > MaxPositionHistory {
		s.positionHistory[position.VehicleId] = s.positionHistory[position.VehicleId][1:]
	}
	s.positionHistory[position.VehicleId] = append(s.positionHistory[position.VehicleId], position)
}

func MapToPosition(e *hfp.Event) *v1.Position {
	var payload *hfp.Payload

	if e.VehiclePosition != nil {
		payload = e.VehiclePosition
	} else if e.DoorOpen != nil {
		payload = e.DoorOpen
	} else if e.DoorClosed != nil {
		payload = e.DoorClosed
	} else {
		return nil
	}

	if !payload.IsValid() {
		return nil
	}

	return &v1.Position{
		VehicleId: e.VehicleId,
		OrgId:     e.OperatorId,
		Latitude:  *payload.Latitude,
		Longitude: *payload.Longitude,
		Heading:   *payload.Heading,
		Timestamp: (*payload.Timestamp).UnixMilli(),
		Speed:     *payload.Speed,
	}
}
