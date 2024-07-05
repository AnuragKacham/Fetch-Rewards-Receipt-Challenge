package handlers

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type Receipt struct {
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Item  `json:"items"`
	Total        float64 `json:"total,string"`
}

type Item struct {
	ShortDescription string  `json:"shortDescription"`
	Price            float64 `json:"price,string"`
}

var receiptStore = make(map[string]int)

func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	var receipt Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := uuid.New().String()
	points := calculatePoints(&receipt)

	receiptStore[id] = points

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"id": id})
}

func GetPoints(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	points, exists := receiptStore[id]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]int{"points": points})
}

func calculatePoints(receipt *Receipt) int {
	points := 0

	points += countAlphanumeric(receipt.Retailer)
	points += roundTotalPoints(receipt.Total)
	points += multipleOf25Points(receipt.Total)
	points += countItemPairs(receipt.Items)
	points += sumItemDescriptionPoints(receipt.Items)
	points += oddPurchaseDatePoints(receipt.PurchaseDate)
	points += purchaseTimePoints(receipt.PurchaseTime)

	return points
}

func countAlphanumeric(retailName string) int {
	return len(regexp.MustCompile(`[a-zA-Z0-9]`).FindAllString(retailName, -1))
}

func roundTotalPoints(totalAmount float64) int {
	if totalAmount == float64(int(totalAmount)) {
		return 50
	}
	return 0
}

func multipleOf25Points(totalAmount float64) int {
	if math.Mod(totalAmount, 0.25) == 0 {
		return 25
	}
	return 0
}

func countItemPairs(items []Item) int {
	return (len(items) / 2) * 5
}

func sumItemDescriptionPoints(items []Item) int {
	points := 0
	for _, item := range items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, _ := strconv.ParseFloat(fmt.Sprintf("%.2f", item.Price), 64)
			points += int(math.Ceil(price * 0.2))
		}
	}
	return points
}

func oddPurchaseDatePoints(dateStr string) int {
	date, _ := time.Parse("2006-01-02", dateStr)
	if date.Day()%2 != 0 {
		return 6
	}
	return 0
}

func purchaseTimePoints(timeStr string) int {
	time, _ := time.Parse("15:04", timeStr)
	if time.Hour() >= 14 && time.Hour() < 16 {
		return 10
	}
	return 0
}
