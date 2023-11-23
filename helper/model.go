package helper

import (
	"technopartner/test/model/domain"
	"technopartner/test/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		ID:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {

	var categoryResponses []web.CategoryResponse
	for _, category := range categories {
		categoryResponses = append(categoryResponses, ToCategoryResponse(category))
	}

	return categoryResponses
}

func ToTransactionResponse(transaction domain.Transaction) web.TransactionResponse {
	return web.TransactionResponse{
		ID:          transaction.ID,
		Name:        transaction.Name,
		Description: transaction.Description,
		Date:        transaction.Date,
		Nominal:     transaction.Nominal,
		CategoryID:  transaction.CategoryID,
	}
}

func ToTransactionResponses(transactions []domain.Transaction) []web.TransactionResponse {

	var transactionResponses []web.TransactionResponse
	for _, transaction := range transactions {
		transactionResponses = append(transactionResponses, ToTransactionResponse(transaction))
	}

	return transactionResponses
}
