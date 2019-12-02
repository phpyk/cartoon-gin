package dao

import "cartoon-gin/DB"

type ChargeVipConfig struct {
	MyGormModel
	Label string `json:"label"`
	Price float64 `json:"price"`
	MonthCou int `json:"month_cou"`
	Extra int `json:"extra"`
	IsRecommend bool `json:"is_recommend"`
	IsForever bool `json:"is_forever"`
	ApplepayProductId string `json:"-"`
	PackageType string `json:"-"`
	ClientType int `json:"-"`
}

func GetChargeVipConfigs(packageType string) []ChargeVipConfig {
	db, _ := DB.OpenCartoon()
	var rows []ChargeVipConfig
	db.Table("charge_vip_configs").Where("package_type = ?",packageType).Scan(&rows)
	return rows
}