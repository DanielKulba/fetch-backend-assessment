package models

type ProcessReceiptParams struct {
	Receipt
}

type Receipt struct {
	Retailer     string `json:"retailer" binding:"required"`
	PurchaseDate string `json:"purchaseDate" binding:"required"`
	PurchaseTime string `json:"purchaseTime" binding:"required"`
	Items        []Item `json:"items" binding:"required"`
	Total        string `json:"total" binding:"required"`
}

type Item struct {
	ShortDescription string `json:"shortDescription" binding:"required"`
	Price            string `json:"price" binding:"required"`
}

type ProcessReceiptReponse struct {
	ID string `json:"id"`
}

type GetPointsResponse struct {
	Points int `json:"points"`
}

type ErrorResponse struct {
	Error string `json:"description"`
}
