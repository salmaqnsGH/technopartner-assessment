package service

import (
	"context"
)

type HomeService interface {
	TotalSaldoCount(ctx context.Context) float64
	TotalSpendCount(ctx context.Context) float64
	TotalIncomeCount(ctx context.Context) float64
}
