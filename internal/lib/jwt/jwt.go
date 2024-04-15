package jwt

import (
	"github.com/golang-jwt/jwt/v5"
	"sso-learn-project-grpc/internal/domain/models"
	"time"
)

func NewToken(user models.User, app models.App, durationTime time.Duration) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["uid"] = user.ID
	claims["app_id"] = app.ID
	claims["exp"] = time.Now().Add(durationTime).Unix()
	claims["email"] = user.Email

	tokenString, err := token.SignedString([]byte(app.Secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
