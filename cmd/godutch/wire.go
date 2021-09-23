//go:build wireinject
// +build wireinject

package main

import (
	"github.com/blackhorseya/godutch/internal/app/godutch"
	"github.com/blackhorseya/godutch/internal/app/godutch/api"
	"github.com/blackhorseya/godutch/internal/pkg/app"
	"github.com/blackhorseya/godutch/internal/pkg/entity/config"
	"github.com/blackhorseya/godutch/internal/pkg/infra/databases"
	"github.com/blackhorseya/godutch/internal/pkg/infra/idgen"
	"github.com/blackhorseya/godutch/internal/pkg/infra/log"
	"github.com/blackhorseya/godutch/internal/pkg/infra/token"
	"github.com/blackhorseya/godutch/internal/pkg/infra/transports/http"
	"github.com/google/wire"
)

var providerSet = wire.NewSet(
	godutch.ProviderSet,
	log.ProviderSet,
	idgen.ProviderSet,
	config.ProviderSet,
	http.ProviderSet,
	databases.ProviderSet,
	token.ProviderSet,
	api.ProviderSet,
)

// CreateApp serve caller to create an injector
func CreateApp(path string, nodeID int64) (*app.Application, error) {
	panic(wire.Build(providerSet))
}
