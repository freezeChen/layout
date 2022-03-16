package biz

import (
	"context"
	"github.com/freezeChen/layout/internal/data"
	"github.com/freezeChen/layout/internal/model"
	"github.com/go-kratos/kratos/v2/log"
)

type GreeterUsecase struct {
	repo *data.GreeterRepo
	log  *log.Helper
}

func NewGreeterUsecase(repo *data.GreeterRepo, logger log.Logger) *GreeterUsecase {
	return &GreeterUsecase{repo: repo, log: log.NewHelper(logger)}
}

func (uc *GreeterUsecase) Create(ctx context.Context, g *model.Greeter) error {
	return uc.repo.CreateGreeter(ctx, g)
}

func (uc *GreeterUsecase) Update(ctx context.Context, g *model.Greeter) error {
	return uc.repo.UpdateGreeter(ctx, g)
}
