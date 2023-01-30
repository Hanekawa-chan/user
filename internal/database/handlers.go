package database

import (
	"context"
	"github.com/kanji-team/user/internal/app"
	"github.com/kanji-team/user/internal/database/models"
)

func (a *adapter) CreateUser(ctx context.Context, user *app.User) error {
	var err error
	query := `insert into "user" (id, email, name) values($1, $2, $3)`
	dbUser := models.FromDomain(user)

	_, err = a.db.ExecContext(ctx, query, dbUser.Id, dbUser.Email, dbUser.Name)
	return err
}
