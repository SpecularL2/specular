package di

import (
	"context"
	"os"
	"os/signal"

	"github.com/cockroachdb/errors"
	"golang.org/x/sync/errgroup"

	"github.com/specularL2/specular/services/sidecar/internal/service/config"
	"github.com/specularL2/specular/services/sidecar/rollup/rpc/eth"
	"github.com/specularL2/specular/services/sidecar/rollup/services"
	"github.com/specularL2/specular/services/sidecar/rollup/services/disseminator"
	"github.com/specularL2/specular/services/sidecar/rollup/services/validator"

	"github.com/sirupsen/logrus"
)

type WaitGroup interface {
	Add(int)
	Done()
	Wait()
}

type Application struct {
	ctx               context.Context
	log               *logrus.Logger
	config            *config.Config
	systemConfig      *services.SystemConfig
	l1State           *eth.EthState
	batchDisseminator *disseminator.BatchDisseminator
	validator         *validator.Validator
}

func (app *Application) Run() error {
	var _, cancel = context.WithCancel(app.ctx)
	var err error

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	defer func() {
		signal.Stop(c)
		cancel()
	}()

	errGroup, _ := errgroup.WithContext(app.ctx)

	if app.systemConfig.Disseminator().GetIsEnabled() {
		app.log.Info("Starting disseminator...")
		err := app.batchDisseminator.Start(app.ctx, errGroup)
		if err != nil {
			return err
		}
	}

	if app.systemConfig.Validator().GetIsEnabled() {
		app.log.Info("Starting validator...")
		err := app.validator.Start(app.ctx, errGroup)
		if err != nil {
			return err
		}
	}

	if err := errGroup.Wait(); err != nil {
		return errors.Newf("service failed while running: %w", err)
	}
	app.log.Info("app stopped")

	return err
}

func (app *Application) ShutdownAndCleanup() {
	app.log.Info("app shutting down")
}

func (app *Application) GetLogger() *logrus.Logger {
	return app.log
}

func (app *Application) GetContext() context.Context {
	return app.ctx
}

func (app *Application) GetConfig() *config.Config {
	return app.config
}

type TestApplication struct {
	*Application

	Ctx    context.Context
	Log    *logrus.Logger
	Config *config.Config
}
