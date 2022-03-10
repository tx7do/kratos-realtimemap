package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/mqtt"

	"kratos-realtimemap/app/transport/mqtt/internal/conf"
	"kratos-realtimemap/app/transport/mqtt/internal/service"
)

// NewMQTTServer create a mqtt server.
func NewMQTTServer(c *conf.Server, _ log.Logger, s *service.TransportMqttService) *mqtt.Server {
	ctx := context.Background()

	srv := mqtt.NewServer(
		broker.Addrs(c.Mqtt.Addr),
		broker.OptionContext(ctx),
	)

	s.SetMqttBroker(srv)

	_ = srv.RegisterSubscriber("/hfp/v2/journey/ongoing/vp/bus/#",
		s.TransitPostTelemetry,
		broker.SubscribeContext(ctx),
	)

	return srv
}
