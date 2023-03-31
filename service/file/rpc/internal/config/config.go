package config

import "github.com/zeromicro/go-zero/zrpc"

type Config struct {
	zrpc.RpcServerConf
	Cos                 Cos
	LocalVideoPath      string
	LocalCoverPath      string
	StaticFileServiceIP string
}

type Cos struct {
	BucketURL              string
	SecretID               string
	SecretKey              string
	VideoPrefix            string //cos保存视频前缀
	CoverPrefix            string //cos保存封面前缀
	VideoSuffix            string //cos保存视频后缀
	CoverSuffix            string //cos保存封面后缀
	VideoTranscodingSuffix string //转码后缀
}
