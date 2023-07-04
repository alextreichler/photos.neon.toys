package context

import (
	"context"

	"github.com/straightbuggin/photos.neon.toys/models"
)

type key string

const (
	userKey key = "user"
)

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, userKey, user)
}

func User(ctx context.Context) *models.User {
	val := ctx.Value(userKey)
	user, ok := val.(*models.User)
	if !ok {
		// the most likely case is that nothgin was ever stored in the context,
		// so it does not have a type of *models.user
		return nil
	}
	return user
}
