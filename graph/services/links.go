package services

import (
	"context"
	"sqlboiler/graph/db"
	"sqlboiler/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
)

type linkService struct {
	exec boil.ContextExecutor
}

func (u *linkService) CreateLink(ctx context.Context, input model.NewLink) (*model.Link, error) {
	id := uuid.New()
	newLink := db.Link{
		ID:      id.String(),
		Title:   null.StringFrom(input.Title),
		Address: null.StringFrom(input.Address),
	}
	err := newLink.Insert(ctx, u.exec, boil.Infer())
	if err != nil {
		return nil, err
	}

	return &model.Link{
		ID:      newLink.ID,
		Title:   input.Title,
		Address: input.Address,
	}, nil
}
