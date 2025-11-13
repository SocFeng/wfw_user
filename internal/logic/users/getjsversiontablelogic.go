// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package users

import (
	"context"
	"database/sql"
	"fmt"
	"wfw_user/internal/svc"
	"wfw_user/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetJsVersionTableLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewGetJsVersionTableLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetJsVersionTableLogic {
	return &GetJsVersionTableLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetJsVersionTableLogic) GetJsVersionTable() (resp *types.GetJsVersionTableResp, err error) {
	// 写在逻辑中的数据查询
	type Tuser struct {
		Id       int64          `db:"id" json:"id"`
		Name     string         `db:"name" json:"name"`
		Age      int64          `db:"age" json:"age"`
		Gender   sql.NullString `db:"gender"  json:"gender"`
		CreateAt sql.NullTime   `db:"create_at" json:"create_at"`
		UpdateAt sql.NullTime   `db:"update_at" json:"update_at"`
		Num      sql.NullInt64  `db:"num" json:"num"`
		Data     sql.NullString `db:"data" json:"data"`
	}
	tl := make([]*Tuser, 0)
	sqlStr := `SELECT * FROM t_users`

	fmt.Println(l.svcCtx.Databases.Users)
	_ = l.svcCtx.Databases.Users.QueryRows(&tl, sqlStr)

	for _, t := range tl {
		fmt.Println("===", t.Id)
		fmt.Println("===", t.Name)
		fmt.Println("===", t.Age)
		fmt.Println("===", t.Gender)
		fmt.Println("===", t.CreateAt)
		fmt.Println("===", t.UpdateAt)
		fmt.Println("===", t.Num)
		fmt.Println("===", t.Data)
		fmt.Println("===========================")
	}

	fmt.Println(" === ", tl)

	//自定义查询方法
	_, _ = l.svcCtx.Models.WfwUser.TUsers.GetAllData(l.ctx)

	// 返回数据结果的封装
	resp = &types.GetJsVersionTableResp{
		types.BaseResponse{Code: 200, Message: "Yes"},
		[]types.JsVersionTable{
			{Version: "1.0.0", Branch: "Hello", Note: "hello world", Id: 110, RN: 12},
		},
	}

	return
}
