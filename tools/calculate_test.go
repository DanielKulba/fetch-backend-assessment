package tools

import (
	"receipt-processor-challenge/models"
	"testing"
)

func TestCountAlphanumeric(t *testing.T) {
	input := "M&M Corner Market"
	expected := 14
	result := countAlphanumeric(input)

	if result != expected {
		t.Errorf("CountAlphanumeric(%s) = %d, expected %d", input, result, expected)
	}
}

func TestIsDivisibleBy(t *testing.T) {
	amount := "10.00"

	// Test case where it is divisible
	divisor := 1.0
	if !isDivisibleBy(amount, divisor) {
		t.Errorf("IsDivisibleBy(%s, %f) = false, expected true", amount, divisor)
	}

	// Test case where it is not divisible
	divisor = 0.3
	if isDivisibleBy(amount, divisor) {
		t.Errorf("IsDivisibleBy(%s, %f) = true, expected false", amount, divisor)
	}
}

func TestPerTwoItems(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Item1", Price: "10.00"},
		{ShortDescription: "Item2", Price: "15.00"},
		{ShortDescription: "Item3", Price: "5.00"},
		{ShortDescription: "Item3", Price: "7.00"},
		{ShortDescription: "Item3", Price: "9.00"},
	}

	expected := 2
	result := perTwoItems(items)

	if result != expected {
		t.Errorf("PerTwoItems(%v) = %d, expected %d", items, result, expected)
	}
}

func TestDescriptionMultiplier(t *testing.T) {
	items := []models.Item{
		{ShortDescription: "Ite", Price: "10.00"},
		{ShortDescription: "Item12", Price: "1"},
		{ShortDescription: "Item3", Price: "5.00"},
	}

	expected := 3
	result := descriptionMultiplier(items)

	if result != expected {
		t.Errorf("DescriptionMultiplier(%v) = %d, expected %d", items, result, expected)
	}
}

func TestCheckPurchaseDate(t *testing.T) {
	date := "2022-01-21"

	// Test case where the day is odd
	if !checkPurchaseDate(date) {
		t.Errorf("CheckPurchaseDate(%s) = false, expected true", date)
	}

	// Test case where the day is even
	date = "2022-01-22"
	if checkPurchaseDate(date) {
		t.Errorf("CheckPurchaseDate(%s) = true, expected false", date)
	}
}

func TestCheckPurchaseTime(t *testing.T) {
	// Test case where the time is after 2:00pm and before 4:00pm
	time := "15:30"
	if !checkPurchaseTime(time) {
		t.Errorf("CheckPurchaseTime(%s) = false, expected true", time)
	}

	// Test case where the time is before 2:00pm
	time = "13:30"
	if checkPurchaseTime(time) {
		t.Errorf("CheckPurchaseTime(%s) = true, expected false", time)
	}

	// Test case where the time is after 4:00pm
	time = "16:30"
	if checkPurchaseTime(time) {
		t.Errorf("CheckPurchaseTime(%s) = true, expected false", time)
	}

}

func TestGetPoints(t *testing.T) {

	receipt := models.Receipt{
		Retailer:     "Target",
		PurchaseDate: "2022-01-01",
		PurchaseTime: "13:01",
		Items: []models.Item{
			{ShortDescription: "Mountain Dew 12PK", Price: "6.49"},
			{ShortDescription: "Emils Cheese Pizza", Price: "12.25"},
			{ShortDescription: "Knorr Creamy Chicken", Price: "1.26"},
			{ShortDescription: "Doritos Nacho Cheese", Price: "3.35"},
			{ShortDescription: "Klarbrunn 12-PK 12 FL OZ", Price: "12.00"},
		},
		Total: "35.35",
	}

	expected := 28
	result := GetPoints(receipt)

	if result != expected {
		t.Errorf("GetPoints(%v) = %d, expected %d", receipt, result, expected)
	}
}
