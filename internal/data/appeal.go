package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	v1 "review-o/api/review/v1"
	"review-o/internal/biz"
	"review-o/internal/data/model"
)

type reAppealRepo struct {
	data   *Data
	logger *log.Helper
}

func NewreAppealRepo(data *Data, logger log.Logger) biz.AppealRepo {
	return &reAppealRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}
}

func (ra *reAppealRepo) CreateReAppeal(ctx context.Context, info *model.ReviewAppealInfo) error {
	//rpc 调用review-service的程序
	request := &v1.OpCreateAppealRequest{
		Status:   info.Status,
		AppealId: info.AppealID,
		OpRemark: info.OpRemarks,
		OpUser:   info.OpUser,
	}
	_, err := ra.data.rpcClinet.OpReAppeal(ctx, request)
	if err != nil {
		return err
	}
	return nil
}
