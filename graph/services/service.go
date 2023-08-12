package services

import (
	"context"
	"sqlboiler/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
)

type UserService interface {
	GetUserByName(ctx context.Context, name string) (*model.User, error)
}

type LinkService interface {
	CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error)
	GetLinks(ctx context.Context) ([]*model.Link, error)
}

type Services interface {
	UserService
	LinkService
}

type services struct {
	*userService
	*linkService
}

func New(exec boil.ContextExecutor) Services {
	return &services{
		userService: &userService{exec: exec},
		linkService: &linkService{exec: exec},
	}
}
