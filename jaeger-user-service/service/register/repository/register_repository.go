package repository

import (
	"context"
	"jarger-user-service/models"
	"jarger-user-service/service/register"

	"github.com/jmoiron/sqlx"
)

type registerRepository struct {
	db *sqlx.DB
}

func NewPsqlRegisterRepositoryImpl(db *sqlx.DB) register.RegisterRepository {
	return &registerRepository{
		db: db,
	}
}

func (r *registerRepository) FetchUserByUsername(ctx context.Context, username, sourse string) (*models.User, error) {
	var (
		sql     string
		account models.User
		err     error
	)
	sql = `
	SELECT *
	FROM account
	WHERE username = ? AND web_access = ?
	`
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	if err = r.db.GetContext(ctx, &account, sql, username, sourse); err != nil {
		return nil, err
	}

	return &account, nil
}
