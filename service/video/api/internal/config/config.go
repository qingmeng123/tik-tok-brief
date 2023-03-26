package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	FileRPC  zrpc.RpcClientConf
	VideoRPC zrpc.RpcClientConf
	UserRPC  zrpc.RpcClientConf
	JWTAuth  struct {
		AccessSecret string
	}
	MaxVideoSize int64
}
