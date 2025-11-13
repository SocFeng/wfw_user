// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"fmt"
	"github.com/redis/go-redis/v9"
	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"net/http"
	"wfw_user/sql/model"

	//"github.com/wumitech-com/cancer-rpc/rpc/cancer"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
	//"github.com/zeromicro/go-zero/zrpc"
	"wfw_user/internal/config"
)

type ServiceContext struct {
	Config    config.Config
	Databases *databases
	Redis     *redis.Client
	Models    *models
	//Rpcs         *rpcs
	UFileRequest *ufsdk.UFileRequest
}

type databases struct {
	Users sqlx.SqlConn
}
type models struct {
	WfwUser *WfwUsersModel
}

type WfwUsersModel struct {
	TUsers model.TUsersModel
}

//type rpcs struct {
//	ProxyCenter cancerapi.CancerAPIClient
//}

func NewServiceContext(c config.Config) *ServiceContext {
	dbUser := sqlx.NewMysql(c.Databases.Users)

	httpClient := &http.Client{}
	uFileRequest, err := ufsdk.NewFileRequest(c.US3s.FileUpload.ToUfsdkConfig(), httpClient)
	if err != nil {
		panic(err)
	}
	return &ServiceContext{

		// 上下文配置
		Config: c,
		// 上下纹数据库-数据源
		Databases: &databases{
			Users: dbUser,
		},
		// 上下文redis
		Redis: redis.NewClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%d", c.Redis.Host, c.Redis.Port),
			DB:   int(c.Redis.Db),
		}),
		// 上下文模型
		Models: &models{
			WfwUser: &WfwUsersModel{
				TUsers: model.NewTUsersModel(dbUser),
			},
		},
		//上下文文件上传
		UFileRequest: uFileRequest,
	}
}
