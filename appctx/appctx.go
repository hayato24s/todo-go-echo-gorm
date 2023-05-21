package appctx

import (
	"context"

	"github.com/google/uuid"
)

type UserIDKey struct{}

func GetUserID(ctx context.Context) (uuid.UUID, bool) {
	userID, ok := ctx.Value(UserIDKey{}).(uuid.UUID)
	return userID, ok
}

func SetUserID(ctx context.Context, userID uuid.UUID) context.Context {
	return context.WithValue(ctx, UserIDKey{}, userID)
}
