package adapter

import (
	"context"
)

type EntityCheckRequest struct {
	Payload interface{}
}

type RequestHandler func(request EntityCheckRequest)

type RequestManagerAdapter interface {
	SetRequestHandler(handler RequestHandler)
	Start(ctx context.Context) error
}
