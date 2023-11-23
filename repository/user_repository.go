package repository

import (
	"context"
	"database/sql"
	"technopartner/test/model/domain"
)

type UserRepository interface {
	Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error)
}
