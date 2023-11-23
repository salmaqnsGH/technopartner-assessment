package service

import (
	"context"
	"database/sql"
	"technopartner/test/exception"
	"technopartner/test/helper"
	"technopartner/test/repository"

	"github.com/go-playground/validator/v10"
)

type HomeServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewHomeService(transactionRepository repository.TransactionRepository, DB *sql.DB, validate *validator.Validate) HomeService {
	return &HomeServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *HomeServiceImpl) TotalSaldoCount(ctx context.Context) float64 {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	totalSaldo := service.TransactionRepository.TotalSaldoCount(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return totalSaldo
}

func (service *HomeServiceImpl) TotalSpendCount(ctx context.Context) float64 {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	totalSpend := -service.TransactionRepository.TotalSpendCount(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return totalSpend
}

func (service *HomeServiceImpl) TotalIncomeCount(ctx context.Context) float64 {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	totalIncome := service.TransactionRepository.TotalIncomeCount(ctx, tx)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return totalIncome
}
