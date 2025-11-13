package model

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

var _ TUsersModel = (*customTUsersModel)(nil)

type (
	// TUsersModel is an interface to be customized, add more methods here,
	// and implement the added methods in customTUsersModel.
	TUsersModel interface {
		tUsersModel
		GetAllData(ctx context.Context) (any, error)
		withSession(session sqlx.Session) TUsersModel
	}

	customTUsersModel struct {
		*defaultTUsersModel
	}
)

// NewTUsersModel returns a model for the database table.
func NewTUsersModel(conn sqlx.SqlConn) TUsersModel {
	return &customTUsersModel{
		defaultTUsersModel: newTUsersModel(conn),
	}
}

func (m *customTUsersModel) withSession(session sqlx.Session) TUsersModel {
	return NewTUsersModel(sqlx.NewSqlConnFromSession(session))
}

// 自定义数据查询接口
func (m *customTUsersModel) GetAllData(ctx context.Context) (any, error) {
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
	_ = m.conn.QueryRows(&tl, sqlStr)

	for _, t := range tl {
		fmt.Println("*******", t.Id)
		fmt.Println("*******", t.Name)
		fmt.Println("*******", t.Age)
		fmt.Println("*******", t.Gender)
		fmt.Println("*******", t.CreateAt)
		fmt.Println("*******", t.UpdateAt)
		fmt.Println("*******", t.Num)
		fmt.Println("*******", t.Data)
		fmt.Println("===========================")
	}

	return nil, nil
}
