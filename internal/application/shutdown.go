package application

import "context"

type Shutdown func(ctx context.Context) error
