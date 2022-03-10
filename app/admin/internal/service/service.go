package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/tx7do/kratos-transport/broker"
	pb "kratos-realtimemap/api/admin/v1"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewAdminService)

type AdminService struct {
	pb.UnimplementedAdminServer

	log *log.Helper

	mb broker.Broker
	kb broker.Broker
}

func NewAdminService(
	logger log.Logger,
) *AdminService {
	l := log.NewHelper(log.With(logger, "module", "service/admin"))
	return &AdminService{
		log: l,
	}
}
