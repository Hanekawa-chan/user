package database

import (
	"context"
	"user/internal/app"
	"user/internal/database/models"
)

func (a *adapter) CreateUser(ctx context.Context, user *app.User) error {
	var err error
	query := `insert into users (id, email, name) values($1, $2, $3)`
	dbUser := models.FromDomain(user)

	_, err = a.db.ExecContext(ctx, query, dbUser.Id, dbUser.Email, dbUser.Name)
	return err
}
