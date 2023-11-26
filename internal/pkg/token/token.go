package token

import (
	"errors"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/hifat/con-q-api/internal/app/config"
	"github.com/hifat/con-q-api/internal/app/domain/authDomain"
)

type TokenType int64

const (
	REFRESH TokenType = iota
	ACCESS
)

const (
	refresh_name = "refresh-token"
	access_name  = "access-token"
)

var ErrNotFoundTokenType = errors.New("not found token type")

func (t TokenType) name() (string, error) {
	tokenTypes := map[TokenType]string{
		REFRESH: "refresh-token",
		ACCESS:  "access-token",
	}
	if _, ok := tokenTypes[t]; !ok {
		return "", ErrNotFoundTokenType
	}

	return tokenTypes[t], nil
}

func (t TokenType) duration(cfg config.AuthConfig) (time.Duration, error) {
	durations := map[TokenType]time.Duration{
		REFRESH: cfg.RefreshTokenDuration,
		ACCESS:  cfg.AccessTokenDuration,
	}
	if _, ok := durations[t]; !ok {
		return 0, ErrNotFoundTokenType
	}

	return durations[t], nil
}

func (t TokenType) secret(cfg config.AuthConfig) (string, error) {
	secrets := map[TokenType]string{
		REFRESH: cfg.RefreshTokenSecret,
		ACCESS:  cfg.AccessTokenSecret,
	}
	if _, ok := secrets[t]; !ok {
		return "", ErrNotFoundTokenType
	}

	return secrets[t], nil
}

func Signed(cfg config.AppConfig, condiment string, tokenType TokenType, passport authDomain.Passport) (string, error) {
	subject, err := tokenType.name()
	if err != nil {
		return "", err
	}

	duration, err := tokenType.duration(cfg.Auth)
	if err != nil {
		return "", err
	}

	secret, err := tokenType.secret(cfg.Auth)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, AuthClaims{
		Passport: passport,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer:    cfg.Env.AppName,
			Subject:   subject,
			Audience:  []string{"*"},
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(duration)),
			NotBefore: jwt.NewNumericDate(time.Now()),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	})

	return token.SignedString(secret + condiment)
}

func Claims(cfg config.AppConfig, condiment string, tokenType TokenType, tokenString string) (*AuthClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		secret, err := tokenType.secret(cfg.Auth)
		if err != nil {
			return nil, err
		}

		return secret + condiment, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(AuthClaims); ok {
		return &claims, nil
	}

	return nil, err
}
