package dao

import "cartoon-gin/DB"

type ChargeCoinConfig struct {
	MyGormModel
	Label             string  `json:"label"`
	Price             float64 `json:"price"`
	Amount            int     `json:"amount"`
	Extra             int     `json:"extra"`
	IsRecommend       int     `json:"is_recommend"`
	IsDouble          int     `json:"is_double"`
	ApplepayProductId string  `json:"-"`
	PackageType       string  `json:"-"`
	ClientType        int     `json:"-"`
}

func GetChargeCoinConfigs(packageType string) []ChargeCoinConfig {
	db, _ := DB.OpenCartoon()
	var rows []ChargeCoinConfig
	db.Table("charge_coin_configs").Where("package_type = ?", packageType).Scan(&rows)
	return rows
}
