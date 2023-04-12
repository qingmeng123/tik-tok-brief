package main

import (
	"flag"
	"fmt"
	_ "github.com/dtm-labs/driver-gozero"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/rest"
	"tik-tok-brief/service/user/api/internal/config"
	"tik-tok-brief/service/user/api/internal/handler"
	"tik-tok-brief/service/user/api/internal/svc"
)

var configFile = flag.String("f", "etc/user-api.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	server := rest.MustNewServer(c.RestConf, rest.WithCors())
	defer server.Stop()

	ctx := svc.NewServiceContext(c)
	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
