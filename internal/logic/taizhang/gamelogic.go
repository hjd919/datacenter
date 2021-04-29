package logic

import (
	"context"
	"datacenter/taizhang/rpc/taizhang"

	"datacenter/internal/svc"
	"datacenter/internal/types"

	"github.com/tal-tech/go-zero/core/logx"
)

type GameLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGameLogic(ctx context.Context, svcCtx *svc.ServiceContext) GameLogic {
	return GameLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GameLogic) Game(req types.ListReq) (res *types.ListResp, err error) {
	taizhangInfo, err := l.svcCtx.TaizhangRpc.GetTaizhang(l.ctx, &taizhang.TaizhangReq{
		Id: 1,
	})
	if err != nil {
		return
	}

	return &types.ListResp{
		Id:   taizhangInfo.Id,
		Name: taizhangInfo.Id,
	}, nil
}
