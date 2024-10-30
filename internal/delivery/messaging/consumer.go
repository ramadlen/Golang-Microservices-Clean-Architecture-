package messaging

import (
	"Customer/internal/model"
	"context"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// ConsumerHandler tipe untuk menangani pesan yang diterima
type ConsumerHandler func(message amqp091.Delivery) error

// ConsumeTopic mengkonsumsi pesan dari topik (antrian) di RabbitMQ
func ConsumeTopic(ctx context.Context, channel *amqp091.Channel, queue string, log *logrus.Logger, handler ConsumerHandler) {
	msgs, err := channel.Consume(
		queue,             // Nama antrian
		"",                // Consumer tag (bisa kosong untuk auto-generated)
		false,             // Auto-ack
		false,             // Exclusive
		false,             // No-local
		false,             // No-wait
		nil,               // Args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	go func() {
		for d := range msgs {
			if err := handler(d); err != nil {
				log.Errorf("Failed to process message: %v", err)
				// Optionally requeue the message if processing fails
				if err := d.Nack(false, true); err != nil {
					log.Errorf("Failed to requeue message: %v", err)
				}
			} else {
				if err := d.Ack(false); err != nil {
					log.Errorf("Failed to acknowledge message: %v", err)
				}
			}
		}
	}()

	// Tunggu hingga konteks dibatalkan
	<-ctx.Done()
	log.Infof("Closing consumer for queue: %s", queue)
}

// Example handler function
func ExampleHandler(message amqp091.Delivery) error {
	var event model.CustomerEvent
	if err := json.Unmarshal(message.Body, &event); err != nil {
		return err
	}
	// Proses event
	return nil
}
