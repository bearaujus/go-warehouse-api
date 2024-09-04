package auth

import (
	"github.com/bearaujus/go-warehouse-api/internal/resource/user"
)

type MiddlewareAuth struct {
	rUser               user.UserResource
	authUserSecretKey   string
	authAllowedServices map[string]string
}

func NewAuthMiddleware(rUser user.UserResource, authUserSecretKey string, authAllowedServices map[string]string) *MiddlewareAuth {
	return &MiddlewareAuth{rUser: rUser, authUserSecretKey: authUserSecretKey, authAllowedServices: authAllowedServices}
}
