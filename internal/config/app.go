package config

import (
	"Customer/internal/delivery/http/controller"

	"Customer/internal/gateways"
	"Customer/internal/repository"
	"Customer/internal/usecase"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/rabbitmq/amqp091-go"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"gorm.io/gorm"
)

// BootstrapConfig holds the configuration for bootstrapping the application.
type BootstrapConfig struct {
	DB       *gorm.DB
	App      *fiber.App
	Log      *logrus.Logger
	Validate *validator.Validate
	Config   *viper.Viper
	Producer *amqp091.Channel
}

// Bootstrap initializes the application components.
func Bootstrap(config *BootstrapConfig) {
	// Setup repositories
	customerRepository := repository.NewCustomerRepository(config.Log)

	// Setup producer
	// In your config package
customerProducer := gateways.NewCustomerProducer(config.Producer, config.Log)


	// Setup use case
	customerUseCase := usecase.NewCustomerUseCase(config.DB, config.Log, config.Validate, customerRepository, customerProducer)

	// Setup controller
	 controller.NewCustomerController(customerUseCase, config.Log)

	
}
