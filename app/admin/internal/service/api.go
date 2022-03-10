package service

import (
	"context"
	"fmt"
	"google.golang.org/protobuf/types/known/emptypb"
	v1 "kratos-realtimemap/api/admin/v1"
	"kratos-realtimemap/app/admin/internal/pkg/data"
	"sort"
)

func (s *AdminService) GetOrganizations(ctx context.Context, _ *emptypb.Empty) (*v1.GetOrganizationsReply, error) {
	reply := &v1.GetOrganizationsReply{Organizations: make([]*v1.Organization, 0, len(data.AllOrganizations))}

	for _, org := range data.AllOrganizations {
		if len(org.Geofences) > 0 {
			reply.Organizations = append(reply.Organizations, &v1.Organization{
				Id:   org.Id,
				Name: org.Name,
			})
		}
	}

	sort.Slice(reply.Organizations, func(i, j int) bool {
		return reply.Organizations[i].Name < reply.Organizations[j].Name
	})

	return reply, nil
}

func (s *AdminService) GetGeofences(ctx context.Context, req *v1.GetGeofencesReq) (*v1.GetGetGeofencesReply, error) {
	if org, ok := data.AllOrganizations[req.Id]; ok {

		return mapOrganization(org), nil
	} else {
		return nil, v1.ErrorResourceNotFound(fmt.Sprintf("Organization %s not found", req.Id))
	}
}

func mapOrganization(org *data.Organization) *v1.GetGetGeofencesReply {
	geofences := make([]*v1.Geofence, 0, 0)

	for _, geofence := range org.Geofences {
		geofences = append(geofences,
			&v1.Geofence{
				Name:           geofence.Name,
				Longitude:      geofence.CentralPoint.Lng(),
				Latitude:       geofence.CentralPoint.Lat(),
				RadiusInMeters: geofence.RadiusInMeters,
				VehiclesInZone: getMapKeys(geofence.VehiclesInZone),
			})
	}

	sort.Slice(geofences, func(i, j int) bool {
		return geofences[i].Name < geofences[j].Name
	})

	return &v1.GetGetGeofencesReply{
		Id:        org.Id,
		Name:      org.Name,
		Geofences: geofences,
	}
}

func getMapKeys(m map[string]struct{}) []string {
	keys := make([]string, len(m))

	i := 0
	for k := range m {
		keys[i] = k
		i++
	}

	return keys
}

func (s *AdminService) GetPositionsHistory(ctx context.Context, req *v1.GetPositionsHistoryReq) (*v1.GetPositionsHistoryReply, error) {
	his, ok := s.positionHistory[req.Id]
	if !ok {
		return nil, v1.ErrorResourceNotFound(fmt.Sprintf("%s positions history not found", req.Id))
	}
	return &v1.GetPositionsHistoryReply{Positions: his}, nil
}
