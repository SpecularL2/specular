//go:build wireinject
// +build wireinject

package di

import (
	"github.com/google/wire"

	"github.com/specularL2/specular/services/sidecar/internal/service/config"
)

func SetupApplication() (*Application, func(), error) {
	panic(wire.Build(wire.NewSet(
		CommonProvider,
		ConfigProvider,
		SystemConfigProvider,
		L1StateProvider,
		DisseminatorProvider,
		ValidatorProvider,
		wire.Struct(new(Application), "*"))),
	)
}

func SetupApplicationForIntegrationTests(cfg *config.Config) (*TestApplication, func(), error) {
	panic(wire.Build(wire.NewSet(
		CommonProvider,
		SystemConfigProvider,
		L1StateProvider,
		DisseminatorProvider,
		ValidatorProvider,
		wire.Struct(new(Application), "*"),
		wire.Struct(new(TestApplication), "*"))),
	)
}
