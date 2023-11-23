package service

import (
	"context"
	"technopartner/test/model/web"
)

type TransactionService interface {
	Create(ctx context.Context, req web.TransactionCreateRequest) web.TransactionResponse
	Update(ctx context.Context, req web.TransactionUpdateRequest) web.TransactionResponse
	Delete(ctx context.Context, TransactionID int)
	FindByID(ctx context.Context, TransactionID int) web.TransactionResponse
	FindAll(ctx context.Context) []web.TransactionResponse
}
