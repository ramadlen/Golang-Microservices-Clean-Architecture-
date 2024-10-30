package config

import (
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// NewRabbitMQConsumer membuat konsumen RabbitMQ
func NewRabbitMQConsumer(config *viper.Viper, log *logrus.Logger) (*amqp091.Channel, error) {
	// rabbitMQURL := config.GetString("rabbitmq.bootstrap.servers")
	rabbitMQURL := config.GetString("rabbitmq.bootstrap.servers" )
	// consumerTag := config.GetString("rabbitmq.group.id") // Ambil group.id untuk consumer
queueName := config.GetString("rabbitmq.queue.tag")
	// Membuat properti koneksi
	properties := amqp091.NewConnectionProperties()

	// Membuat konfigurasi koneksi
	rabbitmqConfig := amqp091.Config{
		Properties: properties,
	}

	// Membuat koneksi ke RabbitMQ
	conn, err := amqp091.DialConfig(rabbitMQURL, rabbitmqConfig)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Membuat channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, err
	}

	// Mengkonsumsi pesan
	
	msgs, err := channel.Consume(
		"customers-clean-architecture", // Nama antrian
		queueName,                    // Consumer tag
		true,                          // Auto-ack (false untuk manual ack)
		false,                          // Exclusive
		false,                          // No-local
		false,                          // No-wait
		nil,                            // Args
	)
	if err != nil {
		log.Fatalf("Failed to register consumer: %v", err)
		return nil, err
	}

	go func() {
		for msg := range msgs {
			log.Infof("Received message: %s", msg.Body)
			// Proses pesan di sini
			msg.Ack(false) // Mengakui pesan setelah diproses
		}
	}()

	return channel, nil

	
}

// NewRabbitMQProducer membuat producer RabbitMQ
func NewRabbitMQProducer(config *viper.Viper, log *logrus.Logger) (*amqp091.Channel, error) {
	rabbitMQURL := config.GetString("rabbitmq.bootstrap.servers")

	// Membuat koneksi ke RabbitMQ
	conn, err := amqp091.Dial(rabbitMQURL)
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
		return nil, err
	}
	defer conn.Close()

	// Membuat channel
	channel, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
		return nil, err
	}

	// Mendeklarasikan antrian
	_, err = channel.QueueDeclare(
		"customers-clean-architecture", // Nama antrian
		true,                           // Durable
		false,                          // Delete when unused
		false,                          // Exclusive
		false,                          // No-wait
		nil,                            // Args
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
		return nil, err
	}

	return channel, nil
}

// SendMessage mengirimkan pesan ke antrian RabbitMQ
func SendMessage(channel *amqp091.Channel, message string, log *logrus.Logger) error {
	err := channel.Publish(
		"",                                 // Default exchange
		"customers-clean-architecture",    // Routing key (antrian)
		false,                              // Mandatory
		false,                              // Immediate
		amqp091.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		},
	)
	if err != nil {
		log.Errorf("Failed to publish a message: %v", err)
		return err
	}
	log.Infof("Message sent: %s", message)
	return nil
}




