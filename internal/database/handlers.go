package database

import (
	"context"
	"github.com/kanji-team/user/internal/app"
)

func (a *adapter) CreateUser(ctx context.Context, user *app.User) error {
	var err error
	query := "insert into user (id, email, country) values(:id, :email, :country)"

	_, err = a.db.NamedExecContext(ctx, query, user)
	return err
}
