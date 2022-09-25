//go:build wireinject
// +build wireinject

package bootstrap

import (
	"github.com/google/wire"
	"github.com/kainonly/ip2region-sync/api"
	"github.com/kainonly/ip2region-sync/common"
)

func NewAPI() (*api.API, error) {
	wire.Build(
		wire.Struct(new(api.API), "*"),
		wire.Struct(new(common.Inject), "*"),
		Provides,
	)
	return &api.API{}, nil
}
