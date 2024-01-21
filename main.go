package main

import (
	handlers "receipt-processor-challenge/controllers"
	"receipt-processor-challenge/db"

	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func main() {
	router := gin.Default()

	db.SetupDatabase()

	group1 := router.Group("/receipts")
	{
		group1.POST("/process", handlers.ProcessReceipt)
		group1.GET("/:id/points", handlers.GetPoints)
	}

	err := router.Run(":8080")
	if err != nil {
		log.Error(err)
	}
}
