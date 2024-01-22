package tools

import (
	"errors"
	"math"
	"receipt-processor-challenge/models"
	"strconv"
	"strings"
	"unicode"

	log "github.com/sirupsen/logrus"
)

func countAlphanumeric(s string) int {
	total := 0

	for _, char := range s {
		if unicode.IsLetter(char) || unicode.IsDigit(char) {
			total++
		}
	}
	return total
}

func isDivisibleBy(amount string, divisor float64) bool {
	f, err := strconv.ParseFloat(amount, 32)

	if err != nil {
		log.Error(errors.New("could not parse float from amount"))
	}

	return math.Mod(f, divisor) == 0
}

func perTwoItems(items []models.Item) int {
	total := 0

	for i, item := range items {
		// ensure that items are not empty
		if i%2 == 1 && item.ShortDescription != "" {
			total++
		}
	}
	return total
}

func descriptionMultiplier(items []models.Item) int {
	total := 0

	for _, item := range items {
		if item.ShortDescription == "" {
			continue
		}
		if len(strings.Trim(item.ShortDescription, " "))%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 32)
			if err != nil {
				log.Error(err)
				return 0
			}

			total += int(math.Ceil(0.2 * price))
		}
	}
	return total
}

func checkPurchaseDate(date string) bool {
	// extract month from
	date = date[8:]
	intDate, _ := strconv.ParseInt(date, 10, 64)
	return intDate%2 == 1
}

func checkPurchaseTime(time string) bool {
	time = time[:2]
	intTime, _ := strconv.ParseInt(time, 10, 64)
	return intTime >= 14 && intTime < 16
}

func GetPoints(receipt models.Receipt) int {
	points := 0

	// implement rules

	// One point for every alphanumeric character in the retailer name
	points += countAlphanumeric(receipt.Retailer)

	// 50 points if the total is a round dollar amount with no cents
	if isDivisibleBy(receipt.Total, 1) {
		points += 50
	}

	// 25 points if the total is a multiple of 0.25
	if isDivisibleBy(receipt.Total, .25) {
		points += 25
	}

	// 5 points for every two items on the receipt
	points += 5 * perTwoItems(receipt.Items)

	// If the trimmed length of the item description is a multiple of 3, multiply the price
	// by 0.2 and round up to the nearest integer. The result is the number of points earned
	points += descriptionMultiplier(receipt.Items)

	// 6 points if the day in the purchase date is odd
	if checkPurchaseDate(receipt.PurchaseDate) {
		points += 6
	}

	// 10 points if the time of purchase is after 2:00pm and before 4:00pm
	if checkPurchaseTime(receipt.PurchaseTime) {
		points += 10
	}

	return points
}
