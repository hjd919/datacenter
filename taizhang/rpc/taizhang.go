package main

import (
	"flag"
	"fmt"

	"datacenter/taizhang/rpc/internal/config"
	"datacenter/taizhang/rpc/internal/server"
	"datacenter/taizhang/rpc/internal/svc"
	"datacenter/taizhang/rpc/taizhang"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/zrpc"
	"google.golang.org/grpc"
)

var configFile = flag.String("f", "etc/taizhang.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)
	srv := server.NewTaizhangServer(ctx)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		taizhang.RegisterTaizhangServer(grpcServer, srv)
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
