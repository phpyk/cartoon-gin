package dao

import "cartoon-gin/DB"

type Order struct {
	MyGormModel
	OrderSn         string  `json:"order_sn"`
	UserId          int     `json:"user_id"`
	OrderType       int     `json:"order_type"`
	Currency        string  `json:"currency"`
	OrderAmount     float64 `json:"order_amount"`
	PayAmount       float64 `json:"pay_amount"`
	PayType         int     `json:"pay_type"`
	PayTime         int     `json:"pay_time"`
	OrderStatus     int     `json:"order_status"`
	ConfigId        int     `json:"config_id"`
	Price           float64 `json:"price"`
	Label           string  `json:"label"`
	CoinAmount      int     `json:"coin_amount"`
	ExtraCoinAmount int     `json:"extra_coin_amount"`
	MonthCou        int     `json:"month_cou"`
	ExtraDays       int     `json:"extra_days"`
}

const (
	OrderTypeVip  = 1
	OrderTypeCoin = 2

	OrderStatusUnpay = 0
	OrderStatusPaid  = 1
)

func HasBought(userId, orderType, configId int) bool {
	db, _ := DB.OpenCartoon()
	var count int
	db.Table("orders").
		Where("user_id = ?", userId).
		Where("order_type = ?", orderType).
		Where("config_id = ?", configId).
		Where("order_status = ?", OrderStatusPaid).Count(&count)
	return count > 0
}
