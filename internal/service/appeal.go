package service

import (
	"context"

	pb "review-o//api/appeal/v1"
)

type AppealService struct {
	pb.UnimplementedAppealServer
}

func NewAppealService() *AppealService {
	return &AppealService{}
}

func (s *AppealService) OpCreateAppeal(ctx context.Context, req *pb.OpCreateAppealRequest) (*pb.OpCreateAppealReply, error) {
	return &pb.OpCreateAppealReply{}, nil
}
func (s *AppealService) UpdateAppeal(ctx context.Context, req *pb.UpdateAppealRequest) (*pb.UpdateAppealReply, error) {
	return &pb.UpdateAppealReply{}, nil
}
func (s *AppealService) DeleteAppeal(ctx context.Context, req *pb.DeleteAppealRequest) (*pb.DeleteAppealReply, error) {
	return &pb.DeleteAppealReply{}, nil
}
func (s *AppealService) GetAppeal(ctx context.Context, req *pb.GetAppealRequest) (*pb.GetAppealReply, error) {
	return &pb.GetAppealReply{}, nil
}
func (s *AppealService) ListAppeal(ctx context.Context, req *pb.ListAppealRequest) (*pb.ListAppealReply, error) {
	return &pb.ListAppealReply{}, nil
}
