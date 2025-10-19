package main

import (
	"context"
	"fmt"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/handlers"
	"github.com/M-kos/wb_level0/internal/logger"
	"github.com/M-kos/wb_level0/internal/providers"
	"github.com/M-kos/wb_level0/internal/repositories"
	"github.com/M-kos/wb_level0/internal/services"
	"log/slog"
	"net/http"
)

func main() {
	conf := config.New()
	log := logger.NewLogger(conf)

	ctx, cancel := context.WithCancel(context.Background())

	postgresDb, err := db.NewDB(ctx, conf, log)
	if err != nil {
		log.Error(err.Error())
		cancel()
		return
	}

	defer func() {
		postgresDb.Pool.Close()
	}()

	orderRepository, err := repositories.NewOrderRepository(postgresDb)
	if err != nil {
		log.Error(err.Error())
		cancel()
		return
	}

	cache := services.NewOrderCache(conf, orderRepository)
	err = cache.LoadCache(ctx)
	if err != nil {
		log.Error(err.Error())
	}

	orderService := services.NewOrderService(orderRepository, cache)

	kafkaConsumer := providers.NewKafkaConsumer(conf, log, orderService)
	go kafkaConsumer.Consume(ctx)

	router := http.NewServeMux()

	handlers.NewOrderHndler(router, conf, log, orderService)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}

	log.Info("server starting", slog.Int("port", conf.Port))

	if err = server.ListenAndServe(); err != nil {
		log.Error(err.Error())
		cancel()
		return
	}
}
