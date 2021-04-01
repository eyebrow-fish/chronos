package chronos

import (
	"context"
	"errors"
)

type JobFunc func(ctx context.Context) error

func Job(name, cron string, f JobFunc) error {
	return errors.New("unimplemeneted")
}
