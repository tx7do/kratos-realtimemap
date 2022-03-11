package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tx7do/kratos-transport/broker"
	"github.com/tx7do/kratos-transport/transport/websocket"
	v1 "kratos-realtimemap/api/admin/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

type AdminService struct {
	v1.UnimplementedAdminServer

	log *log.Helper

	mb broker.Broker
	kb broker.Broker
	ws *websocket.Server

	positionHistory map[string][]*v1.Position
}

func NewAdminService(logger log.Logger) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{
		log:             l,
		positionHistory: make(map[string][]*v1.Position, 0),
	}
}
