package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"Customer/internal/config"
	"Customer/internal/delivery/messaging"
)

func main() {
	viperConfig := config.NewViper()
	logger := config.NewLogger(viperConfig)
	logger.Info("Starting worker service")
	ctx, cancel := context.WithCancel(context.Background())

	
	go RunCustomerConsumer(logger, viperConfig, ctx)

	terminateSignals := make(chan os.Signal, 1)
	signal.Notify(terminateSignals, syscall.SIGINT, syscall.SIGKILL, syscall.SIGTERM)

	stop := false
	for !stop {
		select {
		case s := <-terminateSignals:
			logger.Info("Got one of stop signals, shutting down worker gracefully, SIGNAL NAME :", s)
			cancel()
			stop = true
		}
	}

	time.Sleep(5 * time.Second) // wait for all consumers to finish processing
}

func RunCustomerConsumer(logger *logrus.Logger, viperConfig *viper.Viper, ctx context.Context) {
	logger.Info("setup customer consumer")
	
	// Membuat konsumen RabbitMQ
	channelOK, err := config.NewRabbitMQConsumer(viper.GetViper(), logger)
	if err != nil {
		log.Fatalf("Failed to create RabbitMQ consumer: %v", err)
	}
	defer channelOK.Close()
	contactHandler := messaging.NewCustomerConsumer(logger)
	messaging.ConsumeTopic(ctx, channelOK, "customers", logger, contactHandler.Consume)

}
