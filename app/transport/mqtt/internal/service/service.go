package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tx7do/kratos-transport/broker"
	v1 "kratos-realtimemap/api/transport/mqtt/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewTransportMqttService)

type TransportMqttService struct {
	v1.UnimplementedTransportMqttServer

	log *log.Helper
	mb  broker.Broker
	kb  broker.Broker
}

func NewTransportMqttService(logger log.Logger) *TransportMqttService {
	l := log.NewHelper(log.With(logger, "module", "service/transport-mqtt"))
	return &TransportMqttService{
		log: l,
	}
}
