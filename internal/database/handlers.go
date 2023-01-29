package database

import (
	"context"
	"github.com/kanji-team/user/internal/app"
	"github.com/kanji-team/user/internal/database/models"
)

func (a *adapter) CreateUser(ctx context.Context, user *app.User) error {
	var err error
	query := "insert into user (id, email, name) values(:id, :email, :name)"

	_, err = a.db.NamedExecContext(ctx, query, models.FromDomain(user))
	return err
}
