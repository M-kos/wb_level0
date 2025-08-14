package models

import "time"

type Order struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	Delivery          Delivery
	Payment           Payment
	Items             []Item
	Locale            Locale
	InternalSignature string
	CustomerID        string
	DeliveryService   DeliveryService
	Shardkey          string
	SmID              int
	DateCreated       time.Time
	OofShard          string
}

type OrderDBModel struct {
	OrderUID          string
	TrackNumber       string
	Entry             string
	DeliveryID        int
	PaymentID         int
	LocaleID          int
	InternalSignature string
	CustomerID        string
	DeliveryServiceID int
	Shardkey          string
	SmID              int
	DateCreated       int
	OofShard          string
}
