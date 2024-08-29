package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"review-o/internal/data/model"
)

type AppealRepo interface {
	CreateReAppeal(context.Context, *model.ReviewAppealInfo) error
}

type ReAppeal struct {
	repo AppealRepo
	log  *log.Helper
}

func NewReAppeal(appeal AppealRepo, logger log.Logger) *ReAppeal {
	return &ReAppeal{repo: appeal, log: log.NewHelper(logger)}
}

func (ra *ReAppeal) CreateReAppeal(ctx context.Context, info *model.ReviewAppealInfo) error {
	if err := ra.repo.CreateReAppeal(ctx, info); err != nil {
		return err
	}
	return nil
}
