package logic

import (
	"context"

	"datacenter/taizhang/rpc/internal/svc"
	"datacenter/taizhang/rpc/taizhang"

	"github.com/tal-tech/go-zero/core/logx"
)

type GetTaizhangLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetTaizhangLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTaizhangLogic {
	return &GetTaizhangLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetTaizhangLogic) GetTaizhang(in *taizhang.TaizhangReq) (*taizhang.TaizhangResp, error) {
	// todo: add your logic here and delete this line

	return &taizhang.TaizhangResp{}, nil
}
