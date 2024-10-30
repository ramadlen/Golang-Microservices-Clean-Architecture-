package usecase

import (
	"Customer/internal/entity"
	"Customer/internal/gateways"

	"Customer/internal/model"
	"Customer/internal/model/converter"
	"Customer/internal/repository"
	"context"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"gorm.io/gorm"
)

type CustomerUseCase struct {
	DB                *gorm.DB
	Log               *logrus.Logger
	Validate          *validator.Validate
	CustomerRepository *repository.CustomerRepository
	CustomerProducer   *gateways.CustomerProducer
}

func NewCustomerUseCase(db *gorm.DB, logger *logrus.Logger, validate *validator.Validate,
	customerRepository *repository.CustomerRepository, customerProducer *gateways.CustomerProducer) *CustomerUseCase {
	return &CustomerUseCase{
		DB:                db,
		Log:               logger,
		Validate:          validate,
		CustomerRepository: customerRepository,
		CustomerProducer:   customerProducer,
	}
}

func (c *CustomerUseCase) Create(ctx context.Context, request *model.CreateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer := &entity.Customer{
		CustomerID:   request.CustomerID,
		Nama_Lengkap: request.Nama_Lengkap,
		Alamat:  request.Alamat,
		Email:     request.Email,
		NoTelepon:     request.NoTelepon,
		TanggalLahir: request.TanggalLahir,
		TanggalBergabung: request.TanggalBergabung,
	}

	if err := c.CustomerRepository.Create(tx, customer); err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error creating customer")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CustomerToEvent(customer)
	if err := c.CustomerProducer.Send(event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Update(ctx context.Context, request *model.UpdateCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	customer := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, customer, request.CustomerID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	customer.Nama_Lengkap = request.Nama_Lengkap
	customer.Alamat = request.Alamat
	customer.Email = request.Email
	customer.NoTelepon = request.NoTelepon
	customer.TanggalLahir = request.TanggalLahir
	customer.TanggalBergabung = request.TanggalBergabung

	if err := c.CustomerRepository.Update(tx, customer); err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return nil, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return nil, fiber.ErrInternalServerError
	}

	event := converter.CustomerToEvent(customer)
	if err := c.CustomerProducer.Send(event); err != nil {
		c.Log.WithError(err).Error("error publishing contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(customer), nil
}

func (c *CustomerUseCase) Get(ctx context.Context, request *model.GetCustomerRequest) (*model.CustomerResponse, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, fiber.ErrBadRequest
	}

	contact := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, contact, request.CustomerID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrNotFound
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return nil, fiber.ErrInternalServerError
	}

	return converter.CustomerToResponse(contact), nil
}

func (c *CustomerUseCase) Delete(ctx context.Context, request *model.DeleteCustomerRequest) error {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return fiber.ErrBadRequest
	}

	contact := new(entity.Customer)
	if err := c.CustomerRepository.FindById(tx, contact, request.CustomerID); err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return fiber.ErrNotFound
	}

	if err := c.CustomerRepository.Delete(tx, contact); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return fiber.ErrInternalServerError
	}

	return nil
}

func (c *CustomerUseCase) Search(ctx context.Context, request *model.SearchCustomerRequest) ([]model.CustomerResponse, int64, error) {
	tx := c.DB.WithContext(ctx).Begin()
	defer tx.Rollback()

	if err := c.Validate.Struct(request); err != nil {
		c.Log.WithError(err).Error("error validating request body")
		return nil, 0, fiber.ErrBadRequest
	}

	customers, total, err := c.CustomerRepository.Search(tx, request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	if err := tx.Commit().Error; err != nil {
		c.Log.WithError(err).Error("error getting contacts")
		return nil, 0, fiber.ErrInternalServerError
	}

	responses := make([]model.CustomerResponse, len(customers))
	for i, contact := range customers {
		responses[i] = *converter.CustomerToResponse(&contact)
	}

	return responses, total, nil
}
