package router

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
	// "github.com/swaggo/swag/example/basic/docs"

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
	r := gin.Default()
	v1 := r.Group("/v1")

	v1.Use(ginMiddleware.ConfigInject(gr.cfg))
	v1.Use(ginMiddleware.AppInject(gr.app))

	v1.GET("liveness", ginHandler.Liveness)
	v1.POST("schedules", ginHandler.SchedulerTriggerCreate)
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	port := fmt.Sprintf(":%s", gr.cfg.Port)
	r.Run(port)

	return nil
}
