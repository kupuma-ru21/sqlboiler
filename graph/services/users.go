package services

import (
	"context"
	"sqlboiler/graph/db"
	"sqlboiler/graph/model"

	"github.com/volatiletech/sqlboiler/v4/boil"
	"github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type userService struct {
	exec boil.ContextExecutor
}

func (u *userService) GetUserByName(ctx context.Context, name string) (*model.User, error) {
	user, err := db.Users(
		qm.Select(db.UserTableColumns.ID, db.UserTableColumns.Username),
		db.UserWhere.Username.EQ(name),
	).One(ctx, u.exec)
	if err != nil {
		return nil, err
	}
	return &model.User{
		ID:   string(rune(user.ID)),
		Name: user.Username,
	}, nil
}
