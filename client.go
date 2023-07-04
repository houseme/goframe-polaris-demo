package main

import (
	"fmt"

	"github.com/gogf/gf/contrib/registry/polaris/v2"
	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"

	v1 "github.com/houseme/goframe-polaris-demo/api/user/v1"
)

func main() {
	var (
		ctx = grpcx.Ctx.NewOutgoing(gctx.New(), g.Map{
			"UserId":   "1000",
			"UserName": "john",
		})
		conf = config.NewDefaultConfiguration([]string{"192.168.110.30:8091"})
	)
	conf.Consumer.LocalCache.SetPersistDir("./manifest/logs/polaris/backup")
	if err := api.SetLoggersDir("./manifest/logs/polaris/log"); err != nil {
		g.Log().Fatal(ctx, err)
	}
	grpcx.Resolver.Register(polaris.NewWithConfig(conf, polaris.WithTTL(5)))
	gsel.SetBuilder(gsel.NewBuilderRoundRobin())
	var (
		conn   = grpcx.Client.MustNewGrpcClientConn("GoFrame-Polaris-Demo", grpcx.Balancer.WithRandom())
		client = v1.NewUserClient(conn)
	)

	g.Log().Infof(ctx, `outgoing data: %v`, grpcx.Ctx.OutgoingMap(ctx).Map())
	res, err := client.Create(ctx, &v1.CreateReq{Nickname: "houseme", Passport: "1234568" + grand.Digits(6), Password: "123456"})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", res)
	fmt.Println("Response:", res)
	r, err := client.GetOne(ctx, &v1.GetOneReq{Id: res.GetId()})
	if err != nil {
		g.Log().Error(ctx, err)
		return
	}
	g.Log().Debug(ctx, "Response:", r, " user string:", r.String(), " createTime:", r.GetUser().GetCreateAt().AsTime().String())
	fmt.Println("Response:", r)
}
