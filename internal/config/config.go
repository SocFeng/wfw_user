package config

import (
	"time"

	ufsdk "github.com/ufilesdk-dev/ufile-gosdk"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	LocalEnv bool
	rest.RestConf
	Auth struct {
		AccessSecret string
		AccessExpire int64
	}
	Redis struct {
		Host string
		Port int64
		Db   int64
	}

	Databases struct {
		Users string
	}
	US3s struct {
		FileUpload US3Config
	}
	Rpcs struct {
		Proxycenter       zrpc.RpcClientConf
		PhoneBackendApi   zrpc.RpcClientConf
		PhoneBackendApiHk zrpc.RpcClientConf
		CancerRpc         zrpc.RpcClientConf
		CancerRpcHk       zrpc.RpcClientConf
		SignalingInternal zrpc.RpcClientConf
	}
	Executors struct {
		LoadRoleAuthorizations struct {
			Interval time.Duration `json:",default=10m"`
		}
	}
}

type US3Config struct {
	PublicKey       string
	PrivateKey      string
	BucketHost      string
	BucketName      string
	FileHost        string
	VerifyUploadMD5 bool
	Endpoint        string
}

func (a US3Config) ToUfsdkConfig() *ufsdk.Config {
	return &ufsdk.Config{
		PublicKey:       a.PublicKey,
		PrivateKey:      a.PrivateKey,
		BucketHost:      a.BucketHost,
		BucketName:      a.BucketName,
		FileHost:        a.FileHost,
		VerifyUploadMD5: a.VerifyUploadMD5,
		Endpoint:        a.Endpoint,
	}
}
