package authutil

import (
	"errors"
	"fmt"
	"github.com/bearaujus/go-warehouse-api/internal/model"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"time"
)

func GenerateAuthToken(data string, signKey string, ttl time.Duration) (string, error) {
	tokenIssuedAt := time.Now()
	tokenExpiresAt := tokenIssuedAt.Add(ttl)
	tokenClaims := jwt.RegisteredClaims{
		Issuer:    data,
		ExpiresAt: jwt.NewNumericDate(tokenExpiresAt),
		IssuedAt:  jwt.NewNumericDate(tokenIssuedAt),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, tokenClaims)
	tokenString, err := token.SignedString([]byte(signKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func GenerateAndSetAuthTokenForHTTPRequestHeader(req *http.Request, data string, signKey string, ttl time.Duration) error {
	token, err := GenerateAuthToken(data, signKey, ttl)
	if err != nil {
		return err
	}
	req.Header.Set(model.AuthHTTPHeaderKey, fmt.Sprintf("Bearer %v", token))
	return nil
}

func ReadAuthTokenData(rawToken string, signKey string) (string, error) {
	token, err := jwt.Parse(rawToken, func(t *jwt.Token) (interface{}, error) { return []byte(signKey), nil })
	if err != nil {
		return "", err
	}

	tokenClaims, ok := token.Claims.(jwt.Claims)
	if tokenClaims == nil || !ok || !token.Valid {
		return "", errors.New("invalid auth token")
	}

	id, err := tokenClaims.GetIssuer()
	if err != nil {
		return "", errors.New("invalid auth token")
	}

	return id, nil
}

func ReadAuthTokenDataFromHTTPRequest(rCtx *app.RequestContext, signKey string) (string, error) {
	rawToken := rCtx.Request.Header.Get(model.AuthHTTPHeaderKey)
	if rawToken == "" {
		return "", errors.New("auth token is not present")
	}

	if strings.HasPrefix(rawToken, "Bearer ") {
		rawTokenSplit := strings.Split(rawToken, "Bearer ")
		if len(rawTokenSplit) != 2 {
			return "", errors.New("invalid auth token")
		}
		rawToken = rawTokenSplit[1]
	}

	return ReadAuthTokenData(rawToken, signKey)
}
