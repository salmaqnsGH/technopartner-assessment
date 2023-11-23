package repository

import (
	"context"
	"database/sql"
	"errors"
	"technopartner/test/helper"
	"technopartner/test/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}

func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users(username,password) VALUES ($1,$2) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, user.Username, user.Password)
	err := row.Scan(&user.ID)
	helper.PanicIfError(err)

	return user
}

func (repository *UserRepositoryImpl) FindByUsername(ctx context.Context, tx *sql.Tx, username string) (domain.User, error) {
	SQL := "SELECT id,username,password FROM users WHERE username = $1"
	rows, err := tx.QueryContext(ctx, SQL, username)
	helper.PanicIfError(err)

	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.ID, &user.Username, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user is not found")
	}
}
