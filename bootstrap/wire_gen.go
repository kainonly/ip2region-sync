// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package bootstrap

import (
	"github.com/kainonly/ip2region-sync/api"
	"github.com/kainonly/ip2region-sync/common"
)

// Injectors from wire.go:

func NewAPI() (*api.API, error) {
	values, err := LoadValues()
	if err != nil {
		return nil, err
	}
	db, err := UseGorm(values)
	if err != nil {
		return nil, err
	}
	inject := &common.Inject{
		Values: values,
		Db:     db,
	}
	apiAPI := &api.API{
		Inject: inject,
	}
	return apiAPI, nil
}
