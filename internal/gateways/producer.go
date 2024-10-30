package gateways

import (
	"Customer/internal/model"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// ProducerRabbitMQ struct untuk RabbitMQ
type ProducerRabbitMQ[T model.Event] struct {
	Channel *amqp091.Channel
	Topic   string
	Log     *logrus.Logger
}

// NewProducerRabbitMQ untuk inisialisasi producer RabbitMQ
func NewProducerRabbitMQ[T model.Event](channel *amqp091.Channel, topic string, log *logrus.Logger) *ProducerRabbitMQ[T] {
	return &ProducerRabbitMQ[T]{Channel: channel, Topic: topic, Log: log}
}

// GetTopic mengembalikan nama topik
func (p *ProducerRabbitMQ[T]) GetTopic() string {
	return p.Topic
}

// Send mengirimkan event ke RabbitMQ
func (p *ProducerRabbitMQ[T]) Send(event T) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal event")
		return err
	}

	err = p.Channel.Publish(
		"",               // Default exchange
		p.Topic,         // Routing key (nama antrian)
		false,           // Mandatory
		false,           // Immediate
		amqp091.Publishing{
			ContentType: "application/json",
			Body:        value,
			Headers:     amqp091.Table{"event_id": event.GetId()},
		},
	)
	if err != nil {
		p.Log.WithError(err).Error("failed to publish message")
		return err
	}

	p.Log.Infof("Message sent to %s: %s", p.Topic, string(value))
	return nil
}
