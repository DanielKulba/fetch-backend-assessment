package db

import (
	"errors"
	"receipt-processor-challenge/models"
)

// key: unique id, value: Receipt model
var data map[string]models.Receipt

// initializes the map used to store values
func SetupDatabase() error {
	data = make(map[string]models.Receipt)
	return nil
}

// returns the receipt with the given id
// error if there is no value for the given key
func GetReceipt(id string) (models.Receipt, error) {
	receipt := data[id]

	if receipt.Retailer == "" {
		return models.Receipt{}, errors.New("No receipt found for that id")
	}

	data[id] = receipt
	return receipt, nil
}

// saves the given id-receipt pair
func SaveReceipt(id string, receipt models.Receipt) error {
	data[id] = receipt
	return nil
}
