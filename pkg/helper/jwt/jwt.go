package jwthelper

import (
	"time"

	"github.com/ArthurTirta/monogo/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/google/uuid"
)

func GenerateAccessToken(cfg *config.JWTConfig, userID uuid.UUID) (string, int64, error) {
	exp := time.Now().Add(time.Duration(cfg.AccessTokenExpiryInHours) * time.Hour).Unix()
	claims := jwt.MapClaims{
		"sub": userID.String(),
		"iss": cfg.TokenIssuer,
		"aud": cfg.TokenAudience,
		"exp": exp,
		"iat": time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signed, err := token.SignedString([]byte(cfg.SecretKey))
	if err != nil {
		return "", 0, err
	}

	return signed, exp, nil
}

func ValidateToken(cfg *config.JWTConfig, tokenStr string) (*jwt.Token, error) {
	return jwt.Parse(tokenStr, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrTokenUnverifiable
		}
		return []byte(cfg.SecretKey), nil
	})
}
