package routes

import (
	"simple-chat-api/db"
	"simple-chat-api/utils/security"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
)

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type AuthError struct {
	ErrorType string `json:"errorType"`
}

type TokenResponse struct {
	Token string `json:"token"`
}

func AuthRoutes(router fiber.Router) {
	router.Post("/register", func(c *fiber.Ctx) error {
		creds := new(Credentials)
		c.BodyParser(creds)

		if !security.ValidateAgainstRequirements(creds.Password) {
			return c.Status(fiber.StatusBadRequest).
				JSON(AuthError{ErrorType: "PASSWORD_REQUIREMENTS"})
		}

		entity := db.UserEntity{
			Username: creds.Username,
			Passhash: security.HashPassword(creds.Password),
		}

		err := db.CreateUser(c.Context(), entity)

		if err == nil {
			return c.JSON(Credentials{
				Username: creds.Username,
				Password: "*****",
			})
		} else {
			return c.Status(fiber.StatusBadRequest).
				JSON(AuthError{ErrorType: "USERNAME_TAKEN"})
		}
	})

	router.Post("/login", func(c *fiber.Ctx) error {
		creds := new(Credentials)
		c.BodyParser(creds)
		entity, err := db.GetUserByName(c.Context(), creds.Username)

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(AuthError{ErrorType: "BAD_AUTH"})
		}

		if !security.ValidateAgainstHash(creds.Password, entity.Passhash) {
			return c.Status(fiber.StatusUnauthorized).
				JSON(AuthError{ErrorType: "BAD_AUTH"})
		}

		token, err := jwt.NewWithClaims(
			jwt.SigningMethodHS256,
			jwt.MapClaims{"uid": entity.Username},
		).SignedString([]byte(entity.Passhash))

		if err != nil {
			return c.Status(fiber.StatusUnauthorized).
				JSON(AuthError{ErrorType: "JWT_FAIL"})
		}

		return c.JSON(TokenResponse{Token: token})
	})
}
