package service

import (
	"context"
	"database/sql"
	"technopartner/test/exception"
	"technopartner/test/helper"
	"technopartner/test/model/domain"
	"technopartner/test/model/web"
	"technopartner/test/repository"

	"github.com/go-playground/validator/v10"
	"golang.org/x/crypto/bcrypt"
)

type UserServiceImpl struct {
	UserRepository repository.UserRepository
	DB             *sql.DB
	Validate       *validator.Validate
}

func NewUserService(userRepository repository.UserRepository, DB *sql.DB, validate *validator.Validate) UserService {
	return &UserServiceImpl{
		UserRepository: userRepository,
		DB:             DB,
		Validate:       validate,
	}
}

func (service *UserServiceImpl) Create(ctx context.Context, req web.UserCreateRequest) web.UserResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	passwordHash, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.MinCost)
	helper.PanicIfError(err)

	user := domain.User{
		Username: req.Username,
		Password: string(passwordHash),
	}

	user = service.UserRepository.Save(ctx, tx, user)

	return web.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}

func (service *UserServiceImpl) Login(ctx context.Context, req web.UserLoginRequest) web.UserResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	user, err := service.UserRepository.FindByUsername(ctx, tx, req.Username)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password))
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return web.UserResponse{
		ID:       user.ID,
		Username: user.Username,
	}
}
