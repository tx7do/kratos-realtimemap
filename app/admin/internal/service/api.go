package service

import (
	"context"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-realtimemap/api/admin/v1"
)

func (s *AdminService) GetOrganizations(ctx context.Context, _ *emptypb.Empty) (*v1.GetOrganizationsReply, error) {
	return nil, nil
}

func (s *AdminService) GetGeofences(ctx context.Context, req *v1.GetGeofencesReq) (*v1.GetGetGeofencesReply, error) {
	return nil, nil
}

func (s *AdminService) GetPositionsHistory(ctx context.Context, req *v1.GetPositionsHistoryReq) (*v1.GetPositionsHistoryReply, error) {
	return nil, nil
}
