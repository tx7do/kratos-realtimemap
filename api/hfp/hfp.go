package hfp

import "time"

type Payload struct {
	Longitude *float64   `json:"long"` // 经度(WGS84)
	Latitude  *float64   `json:"lat"`  // 纬度(WGS84)
	Heading   *int32     `json:"hdg"`  // 朝向角度[0, 360]
	DoorState *int32     `json:"drst"` // 门状态 0:所有门都已关闭 1:有门打开
	Timestamp *time.Time `json:"tst"`  // 时间戳
	Speed     *float64   `json:"spd"`  // 车速(m/s)
	Odometer  *int32     `json:"odo"`  // 里程(m)
}

func (p *Payload) IsValid() bool {
	return p != nil && p.Latitude != nil && p.Longitude != nil && p.Heading != nil && p.Timestamp != nil && p.Speed != nil && p.DoorState != nil
}

type Event struct {
	VehicleId  string // 车辆ID
	OperatorId string // 司机ID

	VehiclePosition *Payload `json:"VP"`  // 坐标
	DoorOpen        *Payload `json:"DOO"` // 开门
	DoorClosed      *Payload `json:"DOC"` // 关门
}
