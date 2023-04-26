package cmd

import (
	"context"

	"github.com/gogf/gf/contrib/registry/polaris/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
	"google.golang.org/grpc"

	"github.com/houseme/goframe-polaris-demo/internal/controller/user"
)

var (
	// Main is the main command.
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			conf := config.NewDefaultConfiguration([]string{"183.47.111.80:8091"})
			conf.Consumer.LocalCache.SetPersistDir("./manifest/logs/polaris/backup")
			if err := api.SetLoggersDir("./manifest/logs/polaris/log"); err != nil {
				g.Log().Fatal(context.Background(), err)
			}
			grpcx.Resolver.Register(polaris.NewWithConfig(conf, polaris.WithTTL(10)))

			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			user.Register(s)
			s.Run()
			return nil
		},
	}
)
