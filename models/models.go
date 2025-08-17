package models

import "time"

// User represents a platform user.
type User struct {
	UID            string    `json:"uid"`
	Email          string    `json:"email"`
	Plan           string    `json:"plan"` // e.g., "free", "premium"
	StripeCustomerID string    `json:"stripeCustomerId,omitempty"`
	CreatedAt      time.Time `json:"createdAt"`
	UpdatedAt      time.Time `json:"updatedAt"`
}

// Account represents a linked broker account (Master or Slave).
type Account struct {
	ID         string    `json:"id"` // Unique ID for the account in our system
	UserID     string    `json:"userId"`
	Broker     string    `json:"broker"`
	Platform   string    `json:"platform"` // e.g., "mt4", "mt5", "ctrader", "dxtrade"
	Login      string    `json:"login,omitempty"`      // Account login (can be omitted for some API key methods)
	APIKey     string    `json:"apiKey,omitempty"`     // API key or token (consider encryption)
	IsMaster   bool      `json:"isMaster"`
	Status     string    `json:"status"` // e.g., "linked", "disconnected", "error"
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
}

// Mapping represents a master-slave copying relationship.
type Mapping struct {
	ID              string    `json:"id"` // Unique ID for the mapping
	MasterAccountID string    `json:"masterAccountId"`
	SlaveAccountID  string    `json:"slaveAccountId"`
	RiskMode        string    `json:"riskMode"` // e.g., "same", "multiplier", "fixed"
	LotMultiplier   float64   `json:"lotMultiplier,omitempty"` // Used for multiplier mode
	FixedLot        float64   `json:"fixedLot,omitempty"`      // Used for fixed mode
	IsActive        bool      `json:"isActive"`
	CreatedAt       time.Time `json:"createdAt"`
	UpdatedAt       time.Time `json:"updatedAt"`
}

// TradeRecord represents a copied trade execution.
type TradeRecord struct {
	ID        string    `json:"id"` // Unique ID for the trade record
	MappingID string    `json:"mappingId"`
	OrderID   string    `json:"orderId"`   // Original order ID from the master platform
	SlaveOrderID string    `json:"slaveOrderId,omitempty"` // Order ID on the slave platform
	Type      string    `json:"type"`    // e.g., "buy", "sell"
	Symbol    string    `json:"symbol"`
	LotSize   float64   `json:"lotSize"`
	Price     float64   `json:"price"`
	Timestamp time.Time `json:"timestamp"` // Time of the trade event
	Status    string    `json:"status"`  // e.g., "executed", "failed"
	Error     string    `json:"error,omitempty"` // Error message if execution failed
}