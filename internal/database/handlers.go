package database

import (
	"context"
	"github.com/Hanekawa-chan/kanji-user/internal/services/models"
)

func (a *adapter) CreateUser(ctx context.Context, user *models.DBUser) error {
	var err error
	query := "insert into user (id, email, country) values(:id, :email, :country)"

	_, err = a.db.NamedExecContext(ctx, query, user)
	return err
}
