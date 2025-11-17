// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package users

import (
	"context"

	"wfw_user/internal/svc"
	"wfw_user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AboutPostTestJsonReqLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewAboutPostTestJsonReqLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AboutPostTestJsonReqLogic {
	return &AboutPostTestJsonReqLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AboutPostTestJsonReqLogic) AboutPostTestJsonReq(req *types.AboutTestJsonReq) (resp *types.AboutTestAllResp, err error) {
	// todo: add your logic here and delete this line

	resp = &types.AboutTestAllResp{
		BaseResponse: types.BaseResponse{
			Code:    200,
			Message: "Hello",
		},
		AboutString: "This is aboutString",
	}
	return
}
