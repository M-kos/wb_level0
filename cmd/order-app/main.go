package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/consumer"
	"github.com/M-kos/wb_level0/internal/db"
	"github.com/M-kos/wb_level0/internal/handlers"
	"github.com/M-kos/wb_level0/internal/logger"
	"github.com/M-kos/wb_level0/internal/repositories"
	"github.com/M-kos/wb_level0/internal/services"
	"log/slog"
	"net/http"
	"os/signal"
	"syscall"
	"time"
)

const (
	shutdownTimeout = 10 * time.Second
)

func main() {
	if err := run(); err != nil {
		slog.Error(err.Error())
	}
}

func run() error {
	conf := config.New()
	log := logger.NewLogger(conf)

	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer cancel()

	postgresDb, err := db.NewDB(ctx, conf, log)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	orderRepository := repositories.NewOrderRepository(postgresDb)

	cache := services.NewOrderCache(conf, orderRepository)
	err = cache.LoadCache(ctx)
	if err != nil {
		log.Error(err.Error())
	}

	orderService := services.NewOrderService(orderRepository, cache)

	kafkaConsumer, err := consumer.NewKafkaConsumer(conf, log)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	kafkaOrderHandler := handlers.NewKafkaOrderHandler(orderService)
	kafkaHandler, err := consumer.NewConsumerHandler(kafkaOrderHandler, conf, log)
	if err != nil {
		log.Error(err.Error())
		return err
	}

	go func() {
		err := kafkaConsumer.RunConsume(ctx, kafkaHandler)
		if err != nil {
			log.Error(err.Error())
		}
	}()

	router := http.NewServeMux()
	handlers.NewOrderHandler(router, conf, log, orderService)

	server := http.Server{
		Addr:    fmt.Sprintf(":%d", conf.Port),
		Handler: router,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			log.Error("listen and serve: %v", err)
			cancel()
		}
	}()

	log.Info("server starting", slog.Int("port", conf.Port))

	<-ctx.Done()
	log.Info("shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	if err := server.Shutdown(shutdownCtx); err != nil {
		log.Error("shutdown server error: %s", err.Error())
	}

	if err := kafkaConsumer.Close(); err != nil {
		log.Error("shutdown consumer error: %s", err.Error())
	}

	postgresDb.Pool.Close()

	return nil
}
