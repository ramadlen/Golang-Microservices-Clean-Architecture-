package routes

import (
	"github.com/gofiber/fiber/v2"

	"Customer/internal/delivery/http/controller"
	"Customer/internal/delivery/http/middleware"
)

type RouteConfig struct {
	App               *fiber.App
	CustomerController *controller.CustomerController
	AuthMiddleware    fiber.Handler
}

func (c *RouteConfig) Setup() {
	c.SetupAuthRoute()
}


func (c *RouteConfig) SetupAuthRoute() {
	c.App.Use(c.AuthMiddleware)
	c.App.Get("/api/customer",  c.AuthMiddleware, middleware.AuthMiddlewareKu, middleware.AuthLimiter, c.CustomerController.List)
	c.App.Post("/api/customer", c.AuthMiddleware, middleware.AuthMiddlewareKu,middleware.AuthLimiter,c.CustomerController.Create)
	c.App.Put("/api/customer/:customerId", c.AuthMiddleware, middleware.AuthMiddlewareKu,middleware.AuthLimiter,c.CustomerController.Update)
	c.App.Get("/api/customer/:customerId", c.AuthMiddleware, middleware.AuthMiddlewareKu,middleware.AuthLimiter,c.CustomerController.Get)
	c.App.Delete("/api/customer/:customerId", c.AuthMiddleware, middleware.AuthMiddlewareKu,middleware.AuthLimiter,c.CustomerController.Delete)

	
}
