package logic

import (
	"context"

	"github.com/wlevene/loginservice/internal/svc"
	"github.com/wlevene/loginservice/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type EchoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// EchoHandler
func NewEchoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EchoLogic {
	return &EchoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *EchoLogic) Echo(req *types.EchoReq) (resp *types.Response, err error) {
	// todo: add your logic here and delete this line

	return
}
