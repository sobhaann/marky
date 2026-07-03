package auth

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/sobhaann/markdown-notetaking-api/internal/env"
)

type auth struct {
	authConfig env.AuthConfig
}

func NewAuth(env env.Config) *auth {
	return &auth{
		authConfig: env.Auth,
	}
}

func (a *auth) GenerateJWT(userRID uuid.UUID) (string, error) {
	utcNow := time.Now().UTC()
	claims := &jwt.RegisteredClaims{
		Issuer:    a.authConfig.JWTIssuer,
		IssuedAt:  jwt.NewNumericDate(utcNow),
		ExpiresAt: jwt.NewNumericDate(utcNow.Add(a.authConfig.JWTExpiresAt)),
		Subject:   userRID.String(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(a.authConfig.JWTSecret))
	if err != nil {
		//TODO: implement the logger
		return "", err
	}
	return signedToken, nil
}

func (a *auth) ValidateJWT(tokenString string) (uuid.UUID, error) {
	token, err := jwt.ParseWithClaims(tokenString, jwt.RegisteredClaims{}, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.New("unexpected signing method for jwt token")
		}

		return []byte(a.authConfig.JWTSecret), nil
	})
	if err != nil {
		return uuid.Nil, err
	}

	if !token.Valid {
		return uuid.Nil, errors.New("invalid token")
	}

	claimsSubject, err := token.Claims.GetSubject()
	if err != nil {
		return uuid.Nil, err
	}
	userRID, err := uuid.Parse(claimsSubject)
	if err != nil {
		return uuid.Nil, err
	}

	return userRID, nil
}

// TODO: i think this function is to basic
func (a *auth) GenerateRefreshToken() string {
	token := make([]byte, 64)
	rand.Read(token)
	return hex.EncodeToString(token)
}

func (a *auth) GetTokenFromHTTPHeader(headers http.Header) (string, error) {
	authorizationHeader := headers.Get("Authorization")
	authValue, found := strings.CutPrefix(authorizationHeader, "Bearer ")
	if !found {
		return "", errors.New("Bearer token not found")
	}
	return authValue, nil
}
