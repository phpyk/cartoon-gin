package dao

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
	"errors"
	"cartoon-gin/DB"
)

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

func CreateOrder(order *Order, userId, orderType, packageId int) error {
	order.OrderSn = generateOrderSn(userId,orderType)
	order.UserId = userId
	order.OrderType = orderType
	order.ConfigId = packageId

	var amount float64
	var label string
	if orderType == OrderTypeVip {
		vipConfigRow := GetChargeVipConfigRow(packageId)
		order.MonthCou = vipConfigRow.MonthCou
		order.ExtraDays = vipConfigRow.Extra
		amount = vipConfigRow.Price
		label = vipConfigRow.Label
	}else {
		coinConfigRow := GetChargeCoinConfigRow(packageId)
		order.CoinAmount = coinConfigRow.Amount
		if coinConfigRow.IsDouble == 1 && IsFirstBy(userId,orderType,packageId) {
			order.ExtraCoinAmount = coinConfigRow.Amount
		}else {
			order.ExtraCoinAmount = coinConfigRow.Extra
		}
		amount = coinConfigRow.Price
		label = coinConfigRow.Label
	}
	order.OrderAmount = amount
	order.PayAmount = amount
	order.Price = amount
	order.Label = label
	db, _ := DB.OpenCartoon()
	db.NewRecord(order)
	db.Create(order)
	db.First(order)
	if order.ID > 0 {
		return nil
	}else {
		return errors.New("create order faild")
	}
}

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

func IsFirstBy(userId, orderType, packageId int) bool {
	db, _ := DB.OpenCartoon()
	var orderCount int
	db.Table("orders").Where("user_id = ?", userId).Where("order_type = ?",orderType).Where("config_id = ?",packageId).Where("order_status = ?",OrderStatusPaid).Count(&orderCount)
	return orderCount > 0
}

func generateOrderSn(userId, orderType int) string {
	var prefix string
	if orderType == OrderTypeVip {
		prefix = "BV"
	}else {
		prefix = "BC"
	}
	timeStampString := time.Now().Format("200601021504")
	userIdString := fmt.Sprintf("%06d",strconv.Itoa(userId))
	randomString := strconv.Itoa(rand.Intn(9999 - 1000) + 1000)
	return prefix + timeStampString + userIdString + randomString
}