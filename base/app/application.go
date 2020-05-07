package app

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"

	"github.com/JJMeg/web_template/base/appconfig"
	"github.com/JJMeg/web_template/base/config"
	"github.com/JJMeg/web_template/base/logger"
	"github.com/JJMeg/web_template/base/runmode"
)

type Application struct {
	*gin.Engine
	Mode runmode.RunMode

	cfg    *appconfig.AppConfig
	logger *logrus.Logger
}

func NewApplication(mode runmode.RunMode, srcPath string, cfg interface{}) *Application {
	if err := config.Load(string(mode), srcPath, &cfg); err != nil {
		panic(err.Error())
	}

	var appCfg *appconfig.AppConfig
	if err := config.Load(string(mode), srcPath, &appCfg); err != nil {
		panic(err.Error())
	}

	appLogger, err := logger.NewLogger(appCfg.Logger)
	if err != nil {
		panic(err.Error())
	}

	//	set gin with logger
	if !appCfg.Logger.IsStdout() {
		gin.DisableConsoleColor()
	}

	gin.DefaultWriter = appLogger.Out
	gin.SetMode(mode.ParseGinMode())

	//	set engine
	engine := gin.Default()
	appLogger.Infof("Initialized %s in %s mode...", appCfg.Name, mode)

	return &Application{
		Engine: engine,
		Mode:   mode,
		cfg:    appCfg,
		logger: appLogger,
	}
}

func (app *Application) Run() {
	//	set http server config
	s := http.Server{
		Addr:           app.cfg.Server.Host,
		Handler:        app.Engine,
		ReadTimeout:    time.Duration(app.cfg.Server.RequestTimeout) * time.Second,
		WriteTimeout:   time.Duration(app.cfg.Server.ResponseTimeout) * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	app.logger.Infof("Listening on %s...", app.cfg.Server.Host)

	if err := s.ListenAndServe(); err != nil {
		panic(err.Error())
	}
}

func (app *Application) Logger() *logrus.Logger {
	return app.logger
}
