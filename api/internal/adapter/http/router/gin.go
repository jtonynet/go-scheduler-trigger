package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/jtonynet/go-scheduler-trigger/api/docs"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/jtonynet/go-scheduler-trigger/api/bootstrap"
	"github.com/jtonynet/go-scheduler-trigger/api/config"

	ginHandler "github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/http/handler"
	ginMiddleware "github.com/jtonynet/go-scheduler-trigger/api/internal/adapter/http/middleware"
)

type Router interface {
	HandleRequests(ctx context.Context) error
}

type Gin struct {
	cfg config.API
	app bootstrap.REST
}

func NewGin(cfg config.API, app bootstrap.REST) (Router, error) {
	return Gin{
		cfg,
		app,
	}, nil
}

func (gr Gin) HandleRequests(_ context.Context) error {
	docs.SwaggerInfo.BasePath = "/"

	r := gin.Default()
	r.Use(ginMiddleware.ConfigInject(gr.cfg))
	r.GET("liveness", ginHandler.Liveness)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	v1 := r.Group("/v1")
	v1.Use(ginMiddleware.AppInject(gr.app))
	v1.POST("schedules", ginHandler.SchedulerTriggerCreate)

	return r.Run(fmt.Sprintf(":%s", gr.cfg.Port))
}
