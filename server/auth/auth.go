package auth

import (
	"errors"

	"github.com/form3tech-oss/jwt-go"
	"github.com/labstack/echo/v4"
)

var (
	err = errors.New("could not get user from context")
)

// GetUserIDFromContext will return the UserID from context or an error if it could not get it
func GetUserIDFromContext(c echo.Context) (string, error) {
	token := c.Request().Context().Value("user")
	// if we can not pull the user from context
	if token == nil {
		return "", err
	}

	// convert the context value to a jwt token and then turn the claims into a golang map
	t := token.(*jwt.Token)
	claims := t.Claims.(jwt.MapClaims)

	// get the user id from the claims
	id, ok := claims["sub"]
	if !ok {
		return "", errors.New("could not get user id from context")
	}

	// return the user ID and no error if we could pull it from the claims
	return id.(string), nil
}
