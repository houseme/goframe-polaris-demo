package main

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/contrib/registry/polaris/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gsel"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/polarismesh/polaris-go/api"
	"github.com/polarismesh/polaris-go/pkg/config"
)

func main() {
	conf := config.NewDefaultConfiguration([]string{"127.0.0.1:8091"})
	conf.Consumer.LocalCache.SetPersistDir("./manifest/logs/polaris/backup")
	if err := api.SetLoggersDir("./manifest/logs/polaris/log"); err != nil {
		g.Log().Fatal(context.Background(), err)
	}

	gsvc.SetRegistry(polaris.NewWithConfig(conf, polaris.WithTTL(10)))
	gsel.SetBuilder(gsel.NewBuilderRoundRobin())

	for i := 0; i < 1000; i++ {
		res, err := g.Client().Get(gctx.New(), `http://hello-world.svc/`)
		if err != nil {
			panic(err)
		}
		fmt.Println(res.ReadAllString(), " id: ", i, " time: ", time.Now().Format("2006-01-02 15:04:05"), " code: ", res.StatusCode)
		res.Close()
		time.Sleep(time.Second)
	}
}
