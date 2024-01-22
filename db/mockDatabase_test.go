package db

import (
	"receipt-processor-challenge/models"
	"testing"
)

func TestSetupDatabase(t *testing.T) {
	err := SetupDatabase()
	if err != nil {
		t.Errorf("SetupDatabase failed: %v", err)
	}

	// Check if the data map is initialized
	if getData() == nil {
		t.Error("Data map not initialized")
	}
}

func TestSaveReceipt(t *testing.T) {
	// Setup
	SetupDatabase()
	id := "123"
	receipt := models.Receipt{
		Retailer: "Test Retailer",
		// other fields omitted
	}

	// Test SaveReceipt
	err := SaveReceipt(id, receipt)
	if err != nil {
		t.Errorf("SaveReceipt failed: %v", err)
	}

	// Check if the receipt is saved in the data map
	if savedReceipt, ok := getData()[id]; !ok || savedReceipt.Retailer != receipt.Retailer {
		t.Error("SaveReceipt did not save the receipt correctly")
	}
}

func TestGetReceipt(t *testing.T) {
	// Setup
	SetupDatabase()
	id := "456"
	receipt := models.Receipt{
		Retailer: "Test Retailer",
		// other fields omitted
	}
	SaveReceipt(id, receipt)

	// Test GetReceipt with a valid id
	receivedReceipt, err := GetReceipt(id)
	if err != nil {
		t.Errorf("GetReceipt failed for valid id: %v", err)
	}

	// Check if the received receipt matches the saved receipt
	if receivedReceipt.Retailer != receipt.Retailer {
		t.Error("GetReceipt did not return the correct receipt for a valid id")
	}

	// Test GetReceipt with an invalid id
	_, err = GetReceipt("invalid_id")
	if err == nil || err.Error() != "No receipt found for that id" {
		t.Error("GetReceipt did not return an error for an invalid id")
	}
}
