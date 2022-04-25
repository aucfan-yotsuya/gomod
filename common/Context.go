package common

import (
	"context"
	"time"
)

func Context(t time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), t)
}
