package controller

import (
	"Customer/internal/model"
	"Customer/internal/usecase"

	"github.com/gofiber/fiber/v2"
	"github.com/sirupsen/logrus"

	"math"
)

type CustomerController struct {
	UseCase *usecase.CustomerUseCase
	Log     *logrus.Logger
}

func NewCustomerController(useCase *usecase.CustomerUseCase, log *logrus.Logger) *CustomerController {
	return &CustomerController{
		UseCase: useCase,
		Log:     log,
	}
}

func (c *CustomerController) Create(ctx *fiber.Ctx) error {


	request := new(model.CreateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}


	response, err := c.UseCase.Create(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error creating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) List(ctx *fiber.Ctx) error {
	

	request := &model.SearchCustomerRequest{
		CustomerID: ctx.Query("id"),
		Nama_Lengkap:   ctx.Query("nama_lengkap", ""),
		Email:  ctx.Query("email", ""),
		NoTelepon:  ctx.Query("no_telepon", ""),
		TanggalLahir: ctx.Query("tanggal_lahir",""),
		TanggalBergabung: ctx.Query("tanggal_bergabung",""),
		Page:   ctx.QueryInt("page", 1),
		Size:   ctx.QueryInt("size", 10),
	}

	responses, total, err := c.UseCase.Search(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error searching contact")
		return err
	}

	paging := &model.PageMetadata{
		Page:      request.Page,
		Size:      request.Size,
		TotalItem: total,
		TotalPage: int64(math.Ceil(float64(total) / float64(request.Size))),
	}

	return ctx.JSON(model.WebResponse[[]model.CustomerResponse]{
		Data:   responses,
		Paging: paging,
	})
}

func (c *CustomerController) Get(ctx *fiber.Ctx) error {

	request := &model.GetCustomerRequest{
		CustomerID:     ctx.Params("contactId"),
	}

	response, err := c.UseCase.Get(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error getting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Update(ctx *fiber.Ctx) error {

	request := new(model.UpdateCustomerRequest)
	if err := ctx.BodyParser(request); err != nil {
		c.Log.WithError(err).Error("error parsing request body")
		return fiber.ErrBadRequest
	}

	request.CustomerID = ctx.Params("contactId")

	response, err := c.UseCase.Update(ctx.UserContext(), request)
	if err != nil {
		c.Log.WithError(err).Error("error updating contact")
		return err
	}

	return ctx.JSON(model.WebResponse[*model.CustomerResponse]{Data: response})
}

func (c *CustomerController) Delete(ctx *fiber.Ctx) error {
	customerId := ctx.Params("customerId")

	request := &model.DeleteCustomerRequest{
		CustomerID:     customerId,
	}

	if err := c.UseCase.Delete(ctx.UserContext(), request); err != nil {
		c.Log.WithError(err).Error("error deleting contact")
		return err
	}

	return ctx.JSON(model.WebResponse[bool]{Data: true})
}
