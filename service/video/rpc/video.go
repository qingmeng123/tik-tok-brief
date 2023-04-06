package main

import (
	"flag"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"tik-tok-brief/common/snowflake"

	"tik-tok-brief/service/video/rpc/internal/config"
	"tik-tok-brief/service/video/rpc/internal/server"
	"tik-tok-brief/service/video/rpc/internal/svc"
	"tik-tok-brief/service/video/rpc/proto/pb"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/video.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	err := snowflake.Init(c.MachineId)
	if err != nil {
		logx.Error("snowflake.Init err:", err)
		return
	}
	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		pb.RegisterVideoServer(grpcServer, server.NewVideoServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
