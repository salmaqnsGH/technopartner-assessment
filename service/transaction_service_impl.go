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
)

type TransactionServiceImpl struct {
	TransactionRepository repository.TransactionRepository
	DB                    *sql.DB
	Validate              *validator.Validate
}

func NewTransactionService(transactionRepository repository.TransactionRepository, DB *sql.DB, validate *validator.Validate) TransactionService {
	return &TransactionServiceImpl{
		TransactionRepository: transactionRepository,
		DB:                    DB,
		Validate:              validate,
	}
}

func (service *TransactionServiceImpl) Create(ctx context.Context, req web.TransactionCreateRequest) web.TransactionResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transaction := domain.Transaction{
		Name:        req.Name,
		Description: req.Description,
		Date:        req.Date,
		Nominal:     req.Nominal,
		CategoryID:  req.CategoryID,
	}

	transaction = service.TransactionRepository.Save(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Update(ctx context.Context, req web.TransactionUpdateRequest) web.TransactionResponse {
	err := service.Validate.Struct(req)
	helper.PanicIfError(err)

	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindByID(ctx, tx, req.ID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	transaction.Name = req.Name
	transaction.Description = req.Description
	transaction.Nominal = req.Nominal
	transaction.Date = req.Date
	transaction.CategoryID = req.CategoryID

	transaction = service.TransactionRepository.Update(ctx, tx, transaction)

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) Delete(ctx context.Context, transactionID int) {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindByID(ctx, tx, transactionID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	service.TransactionRepository.Delete(ctx, tx, transaction)
}

func (service *TransactionServiceImpl) FindByID(ctx context.Context, transactionID int) web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transaction, err := service.TransactionRepository.FindByID(ctx, tx, transactionID)
	if err != nil {
		panic(exception.NewNotFoundError(err.Error()))
	}

	return helper.ToTransactionResponse(transaction)
}

func (service *TransactionServiceImpl) FindAll(ctx context.Context, startDate string, endDate string) []web.TransactionResponse {
	tx, err := service.DB.Begin()
	helper.PanicIfError(err)

	defer helper.CommitOrRollback(tx)

	transactions := service.TransactionRepository.FindAll(ctx, tx, startDate, endDate)

	return helper.ToTransactionResponses(transactions)
}
