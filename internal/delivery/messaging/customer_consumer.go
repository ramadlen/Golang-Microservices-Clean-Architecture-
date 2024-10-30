package messaging

import (
	"Customer/internal/model"
	"encoding/json"

	amqp "github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// ContactConsumer struct untuk menangani konsumsi pesan
type CustomerConsumer struct {
	Log *logrus.Logger
}

// NewContactConsumer untuk inisialisasi ContactConsumer
func NewCustomerConsumer(log *logrus.Logger) *CustomerConsumer {
	return &CustomerConsumer{
		Log: log,
	}
}

// Consume menangani pesan yang diterima dari RabbitMQ
func (c *CustomerConsumer) Consume(message amqp.Delivery) error {
	contactEvent := new(model.CustomerEvent)
	if err := json.Unmarshal(message.Body, contactEvent); err != nil {
		c.Log.WithError(err).Error("error unmarshalling Contact event")
		return err
	}

	// TODO: Proses event
	c.Log.Infof("Received message from queue 'contacts': %v", contactEvent)

	// Mengakui pesan setelah diproses
	if err := message.Ack(false); err != nil {
		c.Log.WithError(err).Error("failed to acknowledge message")
		return err
	}

	return nil
}
