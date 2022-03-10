package hfp

import "time"

type Payload struct {
	Longitude *float64   `json:"long"`
	Latitude  *float64   `json:"lat"`
	Heading   *int32     `json:"hdg"`
	DoorState *int32     `json:"drst"`
	Timestamp *time.Time `json:"tst"`
	Speed     *float64   `json:"spd"`
}

func (p *Payload) HasValidPosition() bool {
	return p != nil && p.Latitude != nil && p.Longitude != nil && p.Heading != nil && p.Timestamp != nil && p.Speed != nil && p.DoorState != nil
}

type Event struct {
	VehiclePosition *Payload `json:"VP"`
	DoorOpen        *Payload `json:"DOO"`
	DoorClosed      *Payload `json:"DOC"`
	VehicleId       string
	OperatorId      string
}
