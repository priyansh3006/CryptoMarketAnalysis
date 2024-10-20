// internal/models/orderbook.go
package models

import (
	"sync"
	"time"
)

type OrderBookEntry struct {
	Price     float64   `json:"price"`
	Quantity  float64   `json:"quantity"`
	Timestamp time.Time `json:"timestamp"`
}

type OrderBook struct {
	Symbol       string           `json:"symbol"`
	LastUpdateID int64            `json:"last_update_id"`
	Bids         []OrderBookEntry `json:"bids"`
	Asks         []OrderBookEntry `json:"asks"`
	Timestamp    time.Time        `json:"timestamp"`
	mutex        sync.RWMutex
}

type OrderBookUpdate struct {
	Symbol    string           `json:"symbol"`
	UpdateID  int64            `json:"update_id"`
	Bids      []OrderBookEntry `json:"bids"`
	Asks      []OrderBookEntry `json:"asks"`
	Timestamp time.Time        `json:"timestamp"`
}

// OrderBookSnapshot represents a point-in-time copy of the order book
type OrderBookSnapshot struct {
	Symbol    string           `json:"symbol"`
	Sequence  int64            `json:"sequence"`
	Bids      []OrderBookEntry `json:"bids"`
	Asks      []OrderBookEntry `json:"asks"`
	Timestamp time.Time        `json:"timestamp"`
}
