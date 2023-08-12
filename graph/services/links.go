package services

import (
	"context"
	"sqlboiler/graph/db"
	"sqlboiler/graph/model"

	"github.com/google/uuid"
	"github.com/volatiletech/null/v8"
	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
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

func (u *linkService) GetLinks(ctx context.Context) ([]*model.Link, error) {
	dbLinks, err := db.Links(
		qm.Select(
			db.LinkColumns.ID,
			db.LinkColumns.Title,
			db.LinkColumns.Address,
		)).All(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	var links []*model.Link
	for _, dbLink := range dbLinks {
		link := &model.Link{
			ID:      dbLink.ID,
			Title:   dbLink.Title.String,
			Address: dbLink.Address.String,
		}
		links = append(links, link)
	}

	return links, nil
}
