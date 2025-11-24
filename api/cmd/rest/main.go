package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/database"
)

type ScheduleDTO struct {
	UID       uuid.UUID `json:"uuid"`
	Email     string    `json:"email" binding:"required" example:"test001@gmail.com"`
	Message   string    `json:"message" binding:"required" example:"Teste de envio temporizado"`
	TriggerAt string    `json:"UTC_trigger_at" binding:"required" example:"2025-11-18T15:28:00Z"`
}

func main() {

	cfg, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	cacheInMemoDB, err := database.NewInMemory(cfg.Cache.ToInMemoryDB())
	if err != nil {
		log.Fatal("cannot connect in cacheInMemoDB: ", err)
	}

	triggerInMemoDB, err := database.NewInMemory(cfg.Trigger.ToInMemoryDB())
	if err != nil {
		log.Fatal("cannot connect in triggerInMemoDB: ", err)
	}

	//----------------

	r := gin.Default()
	v1 := r.Group("/v1")

	r.GET("liveness", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "ok",
		})
	})

	v1.POST("schedules", func(ctx *gin.Context) {

		var scheduleReq ScheduleDTO
		if err := ctx.ShouldBindBodyWith(&scheduleReq, binding.JSON); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "nok",
			})

			return
		}

		scheduleReq.UID = uuid.New()
		ctxReq := context.Background()
		key := fmt.Sprintf("schedule:%s", scheduleReq.UID.String())

		// CACHED MESSAGE DATA
		expiration, _ := cacheInMemoDB.GetDefaultExpiration(ctxReq)
		err = cacheInMemoDB.Set(
			ctxReq,
			key,
			scheduleReq,
			expiration,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "nok",
			})

			return
		}

		// ONLY KEY TO TRIGGER SEND MESSAGE
		triggerAt, _ := convertUTCTriggerAtToTimeDuration(scheduleReq.TriggerAt)
		err = triggerInMemoDB.Set(
			ctxReq,
			key,
			nil,
			*triggerAt,
		)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"message": "nok",
			})

			return
		}

		ctx.JSON(http.StatusAccepted, gin.H{
			"message":     "ok",
			"scheduleUID": scheduleReq.UID.String(),
		})
	})

	port := fmt.Sprintf(":%s", cfg.API.Port)
	r.Run(port)
}

func convertUTCTriggerAtToTimeDuration(UTCData string) (*time.Duration, error) {
	targetTime, err := time.Parse(time.RFC3339, UTCData)
	if err != nil {
		return nil, fmt.Errorf("erro ao parsear data UTC: %w", err)
	}

	now := time.Now().UTC()
	duration := targetTime.Sub(now)

	if duration <= 0 {
		return nil, errors.New("data UTC já passou")
	}

	fmt.Println("TTL duration:", duration)
	fmt.Println("TTL seconds:", int64(duration.Seconds()))

	return &duration, nil
}
