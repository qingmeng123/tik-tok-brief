package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Cos Cos
}

type Cos struct {
	BucketURL string
	SecretID  string
	SecretKey string
}
