package repository

import (
	"context"
	"database/sql"
	"errors"
	"technopartner/test/helper"
	"technopartner/test/model/domain"
)

type TransactionRepositoryImpl struct {
}

func NewTransactionRepository() TransactionRepository {
	return &TransactionRepositoryImpl{}
}

func (repository *TransactionRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "INSERT INTO transactions(name,description,nominal,date,category_id) VALUES ($1,$2,$3,$4,$5) RETURNING id"
	row := tx.QueryRowContext(ctx, SQL, transaction.Name, transaction.Description, transaction.Nominal, transaction.Date, transaction.CategoryID)
	err := row.Scan(&transaction.ID)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction {
	SQL := "UPDATE transactions SET name=$1, description=$2, nominal=$3, date=$4, category_id=$5 WHERE id = $6"

	_, err := tx.ExecContext(ctx, SQL, transaction.Name, transaction.Description, transaction.Nominal, transaction.Date, transaction.CategoryID, transaction.ID)
	helper.PanicIfError(err)

	return transaction
}

func (repository *TransactionRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) {
	SQL := "DELETE FROM transactions WHERE id = $1"

	_, err := tx.ExecContext(ctx, SQL, transaction.ID)
	helper.PanicIfError(err)
}

func (repository *TransactionRepositoryImpl) FindByID(ctx context.Context, tx *sql.Tx, transactionID int) (domain.Transaction, error) {
	SQL := "SELECT id,name,description,nominal,date,category_id FROM transactions WHERE id = $1"
	rows, err := tx.QueryContext(ctx, SQL, transactionID)
	helper.PanicIfError(err)

	defer rows.Close()

	transaction := domain.Transaction{}
	if rows.Next() {
		err := rows.Scan(&transaction.ID, &transaction.Name, &transaction.Description, &transaction.Nominal, &transaction.Date, &transaction.CategoryID)
		helper.PanicIfError(err)
		return transaction, nil
	} else {
		return transaction, errors.New("transaction is not found")
	}
}

func (repository *TransactionRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction {
	SQL := "SELECT id,name,description,nominal,date,category_id FROM transactions"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)

	defer rows.Close()

	var transactions []domain.Transaction
	for rows.Next() {
		transaction := domain.Transaction{}
		err := rows.Scan(&transaction.ID, &transaction.Name, &transaction.Description, &transaction.Nominal, &transaction.Date, &transaction.CategoryID)
		helper.PanicIfError(err)

		transactions = append(transactions, transaction)
	}

	return transactions
}
