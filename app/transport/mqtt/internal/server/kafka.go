package server

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/kafka"

	"kratos-realtimemap/app/transport/mqtt/internal/conf"
	"kratos-realtimemap/app/transport/mqtt/internal/service"
)

// NewKafkaServer create a kafka server.
func NewKafkaServer(c *conf.Server, _ log.Logger, s *service.TransportMqttService) *kafka.Server {
	ctx := context.Background()

	srv := kafka.NewServer(
		broker.Addrs(c.Kafka.Addrs...),
		broker.OptionContext(ctx),
	)

	s.SetKafkaBroker(srv)

	return srv
}
