package routes

import (
	"context"
	"encoding/base64"
	"errors"
	"simple-chat-api/db"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v3"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	Uid string `json:"uid"`
}

func SecureRouter(app fiber.App, router fiber.Router) fiber.Router {
	secureRouter := router.Group("/secure")

	secureRouter.Use(jwtware.New(jwtware.Config{
		KeyFunc: func(t *jwt.Token) (interface{}, error) {
			token := t.Raw
			claimsJson, err := base64.RawStdEncoding.DecodeString(
				strings.Split(token, ".")[1],
			)

			if err != nil {
				return nil, errors.New("BAD_JWT")
			}

			claims := new(TokenClaims)
			app.Config().JSONDecoder(claimsJson, claims)
			entity, err := db.GetUserByName(context.Background(), claims.Uid)

			if err != nil {
				return nil, errors.New("BAD_JWT")
			}

			return []byte(entity.Passhash), nil
		},
	}))

	return secureRouter
}
