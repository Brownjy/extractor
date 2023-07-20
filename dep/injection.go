package dep

import (
	"context"
	"extractor/conf/grafana"
	"extractor/conf/serialize"
	"extractor/mongo"
	"github.com/dtynn/dix"
	"go.uber.org/fx"
	"honnef.co/go/tools/config"
)

func Extractor(ctx context.Context, logger fx.Printer, target ...interface{}) dix.Option {
	return dix.Options(
		dix.Override(new(GlobalContext), ctx),
		dix.Override(new(*config.Config), serialize.Load),
		dix.Override(new(*mongo.DB), mongo.NewMongoDB),
		dix.Override(new(*grafana.Grafana), grafana.New),
	)
}
