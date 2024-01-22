package controllers

import (
	"net/http"
	"receipt-processor-challenge/db"
	"receipt-processor-challenge/models"
	"receipt-processor-challenge/tools"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ProcessReceipt(c *gin.Context) {
	// parse body into params object
	var params models.ProcessReceiptParams
	err := c.ShouldBindJSON(&params)

	// error if json does not match Receipt model
	if err != nil {
		c.JSON(http.StatusBadRequest, models.ErrorResponse{Error: "The receipt is invalid"})
		return
	}

	// TODO: it would be wise to perform some validation/sanitation on the incoming request as well

	// generate unique id for this receipt
	id := uuid.New().String()

	db.SaveReceipt(id, params.Receipt)

	c.JSON(http.StatusOK, models.ProcessReceiptReponse{
		ID: id,
	})
}

func GetPoints(c *gin.Context) {
	// extract id from url
	id := c.Param("id")

	// TODO: it would be wise to perform some validation/sanitation on the incoming request as well

	// get receipt from db, error if not found
	receipt, err := db.GetReceipt(id)
	if err != nil {
		c.JSON(http.StatusNotFound, models.ErrorResponse{
			Error: err.Error(),
		})
		return
	}

	points := tools.GetPoints(receipt)

	c.JSON(http.StatusOK, models.GetPointsResponse{
		Points: points,
	})
}
