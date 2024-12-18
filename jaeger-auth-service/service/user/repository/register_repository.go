package repository

import (
	"context"
	"jaeger-auth-service/constant"
	"jaeger-auth-service/models"
	"jaeger-auth-service/service/user"

	"github.com/jmoiron/sqlx"
)

type userRepository struct {
	db *sqlx.DB
}

func NewPsqlUserRepositoryImpl(db *sqlx.DB) user.UserRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) FetchAccountByUsername(ctx context.Context, username, source string) (*models.Account, error) {
	var (
		sql     string
		account models.Account
		err     error
	)
	sql = `
	SELECT *
	FROM account
	WHERE username = ? AND web_access = ?
	`
	sql = sqlx.Rebind(sqlx.DOLLAR, sql)
	if err = r.db.GetContext(ctx, &account, sql, username, source); err != nil {
		if err.Error() == constant.SQL_NO_REC {
			return nil, nil
		}
		return nil, err
	}

	return &account, nil
}
