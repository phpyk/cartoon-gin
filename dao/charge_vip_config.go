package dao

import "cartoon-gin/DB"

type ChargeVipConfig struct {
	MyGormModel
	Label             string  `json:"label"`
	Price             float64 `json:"price"`
	MonthCou          int     `json:"month_cou"`
	Extra             int     `json:"extra"`
	IsRecommend       int     `json:"is_recommend"`
	IsForever         int     `json:"is_forever"`
	ApplepayProductId string  `json:"-"`
	PackageType       string  `json:"-"`
	ClientType        int     `json:"-"`
}

func GetChargeVipConfigs(packageType string) []ChargeVipConfig {
	db, _ := DB.OpenCartoon()
	var rows []ChargeVipConfig
	db.Table("charge_vip_configs").Where("package_type = ?", packageType).Scan(&rows)
	return rows
}

func GetChargeVipConfigRow(id int) (row ChargeVipConfig) {
	db, _ := DB.OpenCartoon()
	db.Table("charge_vip_configs").Where("id = ?", id).First(&row)
	return row
}