package jwt

import (
	"time"

	jwtlib "github.com/golang-jwt/jwt/v5"
)

func GenerateToken(
	email string,
	secret string,
) (string, error) {

	claims := jwtlib.MapClaims{
		"email": email,
		"exp": time.Now().
			Add(24 * time.Hour).
			Unix(),
	}

	token := jwtlib.NewWithClaims(
		jwtlib.SigningMethodHS256,
		claims,
	)

	return token.SignedString(
		[]byte(secret),
	)
}
