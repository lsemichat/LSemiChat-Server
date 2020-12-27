package lcontext

import (
	"context"
	"errors"
)

type key string

const (
	userIDKey key = "userID"
)

func SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, userIDKey, userID)
}

func GetUserIDFromContext(ctx context.Context) (string, error) {
	if ctx.Value(userIDKey) == nil {
		return "", errors.New("failed to get userid from context")
	}
	return ctx.Value(userIDKey).(string), nil
}
