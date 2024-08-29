package service

import (
	"context"
	"review-o/internal/biz"
	"review-o/internal/data/model"

	pb "review-o/api/appeal/v1"
)

type AppealService struct {
	pb.UnimplementedAppealServer
	uc *biz.ReAppeal
}

func NewAppealService(usecase *biz.ReAppeal) *AppealService {
	return &AppealService{
		uc: usecase,
	}
}

func (s *AppealService) OpCreateAppeal(ctx context.Context, req *pb.OpCreateAppealRequest) (*pb.OpCreateAppealReply, error) {
	info := &model.ReviewAppealInfo{
		Status:    req.Status,
		AppealID:  req.AppealId,
		OpRemarks: req.OpRemark,
		OpUser:    req.OpUser,
	}
	if err := s.uc.CreateReAppeal(ctx, info); err != nil {
		return nil, err
	}
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
