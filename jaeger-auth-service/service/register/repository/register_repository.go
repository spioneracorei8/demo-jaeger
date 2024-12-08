package repository

import (
	"context"
	"fmt"
	"jaeger-auth-service/models"
	"jaeger-auth-service/service/register"

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

func (r *registerRepository) FetchAccountByUsername(ctx context.Context, username, source string) (*models.Account, error) {
	var (
		sql     string
		account models.Account
		err     error
	)
	fmt.Println("username", username)
	fmt.Println("source", source)
	sql = `
	SELECT *
	FROM account
	WHERE username = ? AND web_access = ?
	`
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	if err = r.db.GetContext(ctx, &account, sql, username, source); err != nil {
		return nil, err
	}

	return &account, nil
}
