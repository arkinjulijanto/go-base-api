package jwt

import (
	"time"

	"github.com/arkinjulijanto/go-base-api/config"
	"github.com/arkinjulijanto/go-base-api/pkg/custom_error"
	"github.com/golang-jwt/jwt/v4"
)

type TokenClaims struct {
	jwt.RegisteredClaims
	Data interface{} `json:"data"`
}

func GenerateJWTToken(data interface{}) (string, error) {
	conf := config.GetEnv()

	expired := time.Duration(conf.JWT_EXPIRED)
	tokenExp := time.Now().Add(time.Minute * expired).Unix()

	claims := &TokenClaims{
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: conf.JWT_ISSUER,
			ExpiresAt: &jwt.NumericDate{
				Time: time.Unix(tokenExp, 0),
			},
			IssuedAt: &jwt.NumericDate{
				Time: time.Now(),
			},
		},
		Data: data,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(conf.JWT_SECRET))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateToken(token string) (*jwt.Token, error) {
	conf := config.GetEnv()

	return jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, custom_error.NewErrorWrapper(custom_error.CodeClientUnauthorized, "invalid token", nil, nil)
		}

		return []byte(conf.JWT_SECRET), nil
	})
}
