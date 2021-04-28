package logic

import (
	"context"

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

func (l *GameLogic) Game(req types.ListReq) (*types.ListResp, error) {
	// todo: add your logic here and delete this line

	return &types.ListResp{}, nil
}
