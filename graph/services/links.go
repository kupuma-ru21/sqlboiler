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

func (u *linkService) CreateLink(ctx context.Context, input model.CreateLinkInput) (*model.Link, error) {
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

func (u *linkService) CreateLinks(ctx context.Context, input []*model.CreateLinkInput) ([]*model.Link, error) {
	var resLinks []*model.Link
	for _, i := range input {
		id := uuid.New()
		newLink := db.Link{
			ID:      id.String(),
			Title:   null.StringFrom(i.Title),
			Address: null.StringFrom(i.Address),
		}
		err := newLink.Insert(ctx, u.exec, boil.Infer())
		if err != nil {
			return nil, err
		}

		resLink := &model.Link{
			ID:      newLink.ID,
			Title:   newLink.Title.String,
			Address: newLink.Address.String,
		}
		resLinks = append(resLinks, resLink)
	}

	return resLinks, nil
}

func (u *linkService) UpdateLink(ctx context.Context, input model.UpdateLinkInput) (*model.Link, error) {
	dbLink, err := db.Links(
		qm.Select(
			db.LinkColumns.ID,
			db.LinkColumns.Title,
			db.LinkColumns.Address,
		),
		db.LinkWhere.ID.EQ(input.ID),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	dbLink.Title = null.StringFrom(input.Title)
	dbLink.Address = null.StringFrom(input.Address)
	if _, err := dbLink.Update(ctx, u.exec, boil.Infer()); err != nil {
		return nil, err
	}

	return &model.Link{
		ID:      dbLink.ID,
		Title:   input.Title,
		Address: input.Address,
	}, nil
}

func (u *linkService) UpdateLinks(ctx context.Context, input []*model.UpdateLinkInput) ([]*model.Link, error) {
	var resLinks []*model.Link
	for _, i := range input {
		dbLink, err := db.Links(
			qm.Select(
				db.LinkColumns.ID,
				db.LinkColumns.Title,
				db.LinkColumns.Address,
			),
			db.LinkWhere.ID.EQ(i.ID),
		).One(ctx, u.exec)
		if err != nil {
			return nil, err
		}

		dbLink.Title = null.StringFrom(i.Title)
		dbLink.Address = null.StringFrom(i.Address)
		if _, err := dbLink.Update(ctx, u.exec, boil.Infer()); err != nil {
			return nil, err
		}
		resLink := &model.Link{
			ID:      dbLink.ID,
			Title:   i.Title,
			Address: i.Address,
		}
		resLinks = append(resLinks, resLink)
	}

	return resLinks, nil
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

func (u *linkService) GetLinkByTitle(ctx context.Context, title string) (*model.Link, error) {
	dbLink, err := db.Links(
		qm.Select(
			db.LinkColumns.ID,
			db.LinkColumns.Title,
			db.LinkColumns.Address,
		),
		db.LinkWhere.Title.EQ(null.StringFrom(title)),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}

	return &model.Link{
		ID:      dbLink.ID,
		Title:   title,
		Address: dbLink.Address.String,
	}, nil
}
