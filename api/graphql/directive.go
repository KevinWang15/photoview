package api

import (
	"context"
	"errors"

	"github.com/99designs/gqlgen/graphql"
	"github.com/photoview/photoview/api/graphql/auth"
)

func IsAdmin(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := auth.UserFromContext(ctx)
	if user == nil || user.Admin == false {
		return nil, errors.New("user must be admin")
	}

	if auth.IsGuestUserFromContext(ctx) {
		return nil, errors.New("logged in as guest users")
	}
	return next(ctx)
}

func IsAuthorized(ctx context.Context, obj interface{}, next graphql.Resolver) (res interface{}, err error) {
	user := auth.UserFromContext(ctx)
	if user == nil {
		return nil, auth.ErrUnauthorized
	}

	return next(ctx)
}
