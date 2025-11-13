// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package users

import (
	"context"
	"fmt"

	"wfw_user/internal/svc"
	"wfw_user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetTestUserOneLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetTestUserOneLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetTestUserOneLogic {
	return &GetTestUserOneLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetTestUserOneLogic) GetTestUserOne(req *types.TestUserReqOne) (resp *types.GetJsVersionTableResp, err error) {

	fmt.Println(" === ", req)
	fmt.Println(" === req.UserId: ", req.UserId)
	fmt.Println(" === req.UserGender: ", req.UserGender)
	fmt.Println(" === req.UserName: ", req.UserName)
	resp = &types.GetJsVersionTableResp{
		types.BaseResponse{Code: 200, Message: "Yes"},
		[]types.JsVersionTable{
			{Version: "1.0.0", Branch: "Hello", Note: "hello world", Id: 110, RN: 12},
		},
	}
	return
}
