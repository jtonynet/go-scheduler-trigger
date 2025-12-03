package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/google/uuid"
	"github.com/jtonynet/go-scheduler-trigger/api/config"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/database"
	"github.com/jtonynet/go-scheduler-trigger/api/internal/email"

	"github.com/redis/go-redis/v9"
)

type ScheduleDTO struct {
	UID       uuid.UUID `json:"uuid"`
	Email     string    `json:"email" binding:"required" example:"test001@gmail.com"`
	Message   string    `json:"message" binding:"required" example:"Teste de envio temporizado"`
	TriggerAt string    `json:"UTC_trigger_at" binding:"required" example:"2025-11-18T15:28:00Z"`
}

func main() {
	cfg2, err := config.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}

	mail := email.New(cfg2.MailNotification)

	ctx := context.Background()

	cfg := redis.Options{
		Addr:     fmt.Sprintf("%s:%s", os.Getenv("PUBSUB_HOST"), os.Getenv("PUBSUB_PORT")),
		Password: os.Getenv("PUBSUB_PASSWORD"),
		DB:       mustInt(os.Getenv("PUBSUB_DB")),
		Protocol: mustInt(os.Getenv("PUBSUB_PROTOCOL")),
	}

	rdb := redis.NewClient(&cfg)

	// Teste de conexão
	if err := rdb.Ping(ctx).Err(); err != nil {
		log.Fatalf("Não conectou ao Redis: %v", err)
	}

	log.Printf("Conectado ao Redis -> %s:%s (DB: %s)",
		os.Getenv("PUBSUB_HOST"),
		os.Getenv("PUBSUB_PORT"),
		os.Getenv("PUBSUB_DB"),
	)

	cacheInMemoDB, err := database.NewInMemory(cfg2.Cache.ToInMemoryDB())
	if err != nil {
		log.Fatal("cannot connect in cacheInMemoDB: ", err)
	}

	// Canal de expiração
	channel := fmt.Sprintf("__keyevent@%s__:expired", os.Getenv("PUBSUB_DB"))
	log.Printf("Escutando expirações no canal: %s", channel)

	// Loop resiliente
	for {
		if err := listenTriggers(ctx, rdb, cacheInMemoDB, channel, mail); err != nil {
			log.Printf("Erro no listener, reconectando em 2s... err=%v", err)
			time.Sleep(2 * time.Second)
		}
	}
}

func listenTriggers(ctx context.Context, rdb *redis.Client, cacheInMemoDB database.InMemory, channel string, mail *email.Mail) error {
	pubsub := rdb.Subscribe(ctx, channel)

	_, err := pubsub.Receive(ctx)
	if err != nil {
		return fmt.Errorf("erro ao inscrever no canal: %w", err)
	}

	ch := pubsub.Channel()

	for msg := range ch {
		key := msg.Payload // ex: "schedule:8f7a..."

		log.Printf("Chave expirada: %s", key)

		//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
		// Tentar recuperar o valor original
		value, err := cacheInMemoDB.Get(ctx, key)
		if err == redis.Nil {
			log.Printf("Shadow key não encontrada: %s (já removida?)", key)
			continue
		} else if err != nil {
			log.Printf("Erro ao recuperar shadow key (%s): %v", key, err)
			continue
		}

		log.Printf("Payload Recuperado: %s", value)

		// PROCESSAR EVENTO DE NEGÓCIO AQUI --------------------
		processExpiration(key, value, mail)
		// ------------------------------------------------------

		// Remover shadow key após processar
		if err := cacheInMemoDB.Delete(ctx, key); err != nil {
			log.Printf("Erro ao remover shadow key %s: %v", key, err)
		} else {
			log.Printf("Shadow key removida: %s", key)
		}
		//+++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
	}

	return fmt.Errorf("canal fechado inesperadamente")
}

// Lógica de negócio quando uma chave expira
func processExpiration(key, value string, mail *email.Mail) {
	log.Printf("Processando expiração da chave [%s] com payload: %s", key, value)
	// Aqui vai sua lógica de trigger, scheduler, etc.

	var scheduleDTO ScheduleDTO
	json.Unmarshal([]byte(value), &scheduleDTO)
	mail.Send(
		scheduleDTO.Email,
		"TESTE!",
		scheduleDTO.Message,
	)
}

func mustInt(v string) int {
	var x int
	fmt.Sscan(v, &x)
	return x
}
