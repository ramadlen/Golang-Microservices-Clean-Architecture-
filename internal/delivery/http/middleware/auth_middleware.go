package middleware

import (
	"Customer/internal/config"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/limiter"
)

// AuthMiddleware checks the authorization header for a valid token.
func AuthMiddlewareKu(ctx *fiber.Ctx) error {
	authHeader := ctx.Get("Authorization") // Corrected spelling to "Authorization"
	if authHeader == "" {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Authorization header missing")
	}

	// Extracting token from the header
	tokenString := authHeader[len("Bearer "):] // Added space after "Bearer"
	if len(tokenString) == 0 {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Token not found")
	}

	// Verify the token
	token, err := config.VerifyToken(tokenString, config.Secret)
	if err != nil {
		return ctx.Status(fiber.StatusUnauthorized).SendString("Invalid token")
	}

	// Store the user ID or relevant information from the token
	ctx.Locals("userID", token) // Adjust as needed based on your token structure
	return ctx.Next()
}

// AuthLimiter applies rate limiting to requests.
func AuthLimiter(ctx *fiber.Ctx) error {
	return limiter.New(limiter.Config{
		Max:         60, // Maximum requests
		Expiration:  40 * time.Second, // Maximum time frame
		LimitReached: func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusTooManyRequests).SendStatus(429)
		},
	})(ctx) // Apply the limiter middleware
}
