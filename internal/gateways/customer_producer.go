package gateways

import (
	"Customer/internal/model"
	"encoding/json"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
)

// CustomerProducer struct untuk RabbitMQ
type CustomerProducer struct {
	Channel *amqp091.Channel
	Topic   string
	Log     *logrus.Logger
}

// NewCustomerProducer untuk inisialisasi producer pengguna RabbitMQ
func NewCustomerProducer(channel *amqp091.Channel, log *logrus.Logger) *CustomerProducer {
	return &CustomerProducer{
		Channel: channel,
		Topic:   "users",
		Log:     log,
	}
}

// Send mengirimkan event pengguna ke RabbitMQ
func (p *CustomerProducer) Send(event *model.CustomerEvent) error {
	value, err := json.Marshal(event)
	if err != nil {
		p.Log.WithError(err).Error("failed to marshal user event")
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
		p.Log.WithError(err).Error("failed to publish user event")
		return err
	}

	p.Log.Infof("User event sent to %s: %s", p.Topic, string(value))
	return nil
}
