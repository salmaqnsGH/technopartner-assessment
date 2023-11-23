package repository

import (
	"context"
	"database/sql"
	"technopartner/test/model/domain"
)

type TransactionRepository interface {
	Save(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Update(ctx context.Context, tx *sql.Tx, transaction domain.Transaction) domain.Transaction
	Delete(ctx context.Context, tx *sql.Tx, transaction domain.Transaction)
	FindByID(ctx context.Context, tx *sql.Tx, transactionID int) (domain.Transaction, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Transaction
	TotalSaldoCount(ctx context.Context, tx *sql.Tx) float64
	TotalSpendCount(ctx context.Context, tx *sql.Tx) float64
	TotalIncomeCount(ctx context.Context, tx *sql.Tx) float64
}
