package main

import (
	"encoding/json"
	"github.com/M-kos/wb_level0/internal/config"
	"github.com/M-kos/wb_level0/internal/logger"
	"github.com/M-kos/wb_level0/internal/producer"
	"time"
)

func main() {
	conf := config.New()
	log := logger.NewLogger(conf)

	kp, err := producer.NewKafkaProducer(conf)
	if err != nil {
		log.Error("failed to init producer", "error", err)
		return
	}

	defer func() {
		err := kp.Close()
		if err != nil {
			log.Error("failed to close kafka producer", "error", err)
		}
	}()

	for _, order := range orders() {
		data, err := json.Marshal(order)
		if err != nil {
			log.Error("failed to marshal order", order["order_uid"].(string), err)
			continue
		}

		key := []byte(order["order_uid"].(string))
		if err := kp.SendMessage(key, data); err != nil {
			log.Error("failed to send order", order["order_uid"].(string), err)
			continue
		}

		log.Info("sent order", "order_uid", order["order_uid"])
		time.Sleep(500 * time.Millisecond)
	}

	log.Info("messages sent")
}

func orders() []map[string]interface{} {
	orders := []map[string]interface{}{
		{
			"order_uid":    "b563feb7b2b84b6test",
			"track_number": "WBILMTESTTRACK",
			"entry":        "WBIL",
			"delivery": map[string]string{
				"name":    "Test Testov",
				"phone":   "+9720000000",
				"zip":     "2639809",
				"city":    "Kiryat Mozkin",
				"address": "Ploshad Mira 15",
				"region":  "Kraiot",
				"email":   "test@gmail.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "b563feb7b2b84b6test",
				"request_id":    "",
				"currency":      "USD",
				"provider":      "wbpay",
				"amount":        1817,
				"payment_dt":    1637907727,
				"bank":          "alpha",
				"delivery_cost": 1500,
				"goods_total":   317,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      9934930,
					"track_number": "WBILMTESTTRACK",
					"price":        453,
					"rid":          "ab4219087a764ae0btest",
					"name":         "Mascaras",
					"sale":         30,
					"size":         "0",
					"total_price":  317,
					"nm_id":        2389212,
					"brand":        "Vivienne Sabo",
					"status":       202,
				},
			},
			"locale":             "en",
			"internal_signature": "",
			"customer_id":        "test",
			"delivery_service":   "meest",
			"shardkey":           "9",
			"sm_id":              99,
			"date_created":       "2021-11-26T06:22:19Z",
			"oof_shard":          "1",
		},
		{
			"order_uid":    "order-001",
			"track_number": "WBTRACK001",
			"entry":        "WBIL",
			"delivery": map[string]interface{}{
				"name":    "John Doe",
				"phone":   "+7900000001",
				"zip":     "101000",
				"city":    "Moscow",
				"address": "Red Square 1",
				"region":  "Moscow",
				"email":   "john@example.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "txn-001",
				"currency":      "USD",
				"provider":      "wbpay",
				"amount":        1500,
				"payment_dt":    time.Now().Unix(),
				"bank":          "alpha",
				"delivery_cost": 300,
				"goods_total":   1200,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      10001,
					"track_number": "WBTRACK001",
					"price":        1200,
					"rid":          "rid-10001",
					"name":         "T-shirt",
					"sale":         10,
					"size":         "L",
					"total_price":  1080,
					"nm_id":        11111,
					"brand":        "Nike",
					"status":       202,
				},
			},
			"locale":             "en",
			"internal_signature": "",
			"customer_id":        "cust-001",
			"delivery_service":   "meest",
			"shardkey":           "1",
			"sm_id":              99,
			"date_created":       time.Now().UTC().Format(time.RFC3339),
			"oof_shard":          "1",
		},
		{
			"order_uid":    "order-002",
			"track_number": "WBTRACK002",
			"entry":        "WBIL",
			"delivery": map[string]interface{}{
				"name":    "Anna Petrova",
				"phone":   "+7900000002",
				"zip":     "190000",
				"city":    "Saint Petersburg",
				"address": "Nevsky Prospekt 10",
				"region":  "Leningrad",
				"email":   "anna@example.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "txn-002",
				"currency":      "USD",
				"provider":      "wbpay",
				"amount":        2500,
				"payment_dt":    time.Now().Unix(),
				"bank":          "tinkoff",
				"delivery_cost": 400,
				"goods_total":   2100,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      10002,
					"track_number": "WBTRACK002",
					"price":        2100,
					"rid":          "rid-10002",
					"name":         "Sneakers",
					"sale":         15,
					"size":         "38",
					"total_price":  1785,
					"nm_id":        11112,
					"brand":        "Adidas",
					"status":       202,
				},
			},
			"locale":             "en",
			"internal_signature": "",
			"customer_id":        "cust-002",
			"delivery_service":   "meest",
			"shardkey":           "2",
			"sm_id":              99,
			"date_created":       time.Now().UTC().Format(time.RFC3339),
			"oof_shard":          "2",
		},
		{
			"order_uid":    "order-003",
			"track_number": "WBTRACK003",
			"entry":        "WBIL",
			"delivery": map[string]interface{}{
				"name":    "Ali Khan",
				"phone":   "+7900000003",
				"zip":     "050000",
				"city":    "Almaty",
				"address": "Abay Ave 5",
				"region":  "Almaty",
				"email":   "ali@example.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "txn-003",
				"currency":      "USD",
				"provider":      "wbpay",
				"amount":        1800,
				"payment_dt":    time.Now().Unix(),
				"bank":          "revolut",
				"delivery_cost": 200,
				"goods_total":   1600,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      10003,
					"track_number": "WBTRACK003",
					"price":        1600,
					"rid":          "rid-10003",
					"name":         "Laptop Bag",
					"sale":         5,
					"size":         "U",
					"total_price":  1520,
					"nm_id":        11113,
					"brand":        "Samsonite",
					"status":       202,
				},
			},
			"locale":             "en",
			"internal_signature": "",
			"customer_id":        "cust-003",
			"delivery_service":   "meest",
			"shardkey":           "3",
			"sm_id":              99,
			"date_created":       time.Now().UTC().Format(time.RFC3339),
			"oof_shard":          "3",
		},
		{
			"order_uid":    "order-004",
			"track_number": "WBTRACK004",
			"entry":        "WBIL",
			"delivery": map[string]interface{}{
				"name":    "Satoshi Tanaka",
				"phone":   "+819000000004",
				"zip":     "1500001",
				"city":    "Tokyo",
				"address": "Shibuya 1-1",
				"region":  "Tokyo",
				"email":   "satoshi@example.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "txn-004",
				"currency":      "JPY",
				"provider":      "wbpay",
				"amount":        210000,
				"payment_dt":    time.Now().Unix(),
				"bank":          "mizuho",
				"delivery_cost": 15000,
				"goods_total":   195000,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      10004,
					"track_number": "WBTRACK004",
					"price":        195000,
					"rid":          "rid-10004",
					"name":         "Headphones",
					"sale":         10,
					"size":         "U",
					"total_price":  175500,
					"nm_id":        11114,
					"brand":        "Sony",
					"status":       202,
				},
			},
			"locale":             "jp",
			"internal_signature": "",
			"customer_id":        "cust-004",
			"delivery_service":   "meest",
			"shardkey":           "4",
			"sm_id":              99,
			"date_created":       time.Now().UTC().Format(time.RFC3339),
			"oof_shard":          "4",
		},
		{
			"order_uid":    "order-005",
			"track_number": "WBTRACK005",
			"entry":        "WBIL",
			"delivery": map[string]interface{}{
				"name":    "Carlos Gomez",
				"phone":   "+34900000005",
				"zip":     "28001",
				"city":    "Madrid",
				"address": "Gran Via 3",
				"region":  "Madrid",
				"email":   "carlos@example.com",
			},
			"payment": map[string]interface{}{
				"transaction":   "txn-005",
				"currency":      "EUR",
				"provider":      "wbpay",
				"amount":        1300,
				"payment_dt":    time.Now().Unix(),
				"bank":          "bbva",
				"delivery_cost": 200,
				"goods_total":   1100,
				"custom_fee":    0,
			},
			"items": []map[string]interface{}{
				{
					"chrt_id":      10005,
					"track_number": "WBTRACK005",
					"price":        1100,
					"rid":          "rid-10005",
					"name":         "Watch",
					"sale":         20,
					"size":         "U",
					"total_price":  880,
					"nm_id":        11115,
					"brand":        "Casio",
					"status":       202,
				},
			},
			"locale":             "es",
			"internal_signature": "",
			"customer_id":        "cust-005",
			"delivery_service":   "meest",
			"shardkey":           "5",
			"sm_id":              99,
			"date_created":       time.Now().UTC().Format(time.RFC3339),
			"oof_shard":          "5",
		},
	}

	return orders
}
