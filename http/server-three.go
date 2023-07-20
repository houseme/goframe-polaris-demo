package main

import (
	"context"

	"github.com/gogf/gf/contrib/registry/polaris/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gipv4"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
)

func main() {
	conf := config.NewDefaultConfiguration([]string{"127.0.0.1:8091"})
	conf.Consumer.LocalCache.SetPersistDir("./manifest/logs/polaris/backup")
	if err := api.SetLoggersDir("./manifest/logs/polaris/log"); err != nil {
		g.Log().Fatal(context.Background(), err)
	}

	// TTL egt 2*time.Second
	gsvc.SetRegistry(polaris.NewWithConfig(conf, polaris.WithTTL(10)))

	s := g.Server(`hello-world.svc`)
	s.BindHandler("/", func(r *ghttp.Request) {
		g.Log().Info(r.Context(), `request received three`)
		var IntranetIpArray, err = gipv4.GetIntranetIpArray()
		if err != nil {
			g.Log().Error(r.Context(), `request received three`, err)
		}
		g.Log().Info(r.Context(), `request received three IntranetIpArray: `, IntranetIpArray)
		r.Response.Write(`Hello world three`)
	})
	s.Run()
}
