package data

import (
	"context"
	"github.com/freezeChen/layout/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type GreeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) *GreeterRepo {
	return &GreeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *GreeterRepo) CreateGreeter(ctx context.Context, g *model.Greeter) error {
	return nil
}

func (r *GreeterRepo) UpdateGreeter(ctx context.Context, g *model.Greeter) error {
	return nil
}
