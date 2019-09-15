package dao

import (
	"cartoon-gin/DB"
	"github.com/jinzhu/gorm"
)

type HomeConfig struct {
	MyGormModel
	gorm.Model
	ModuleType int `json:"module_type"`
	ConfigType int `json:"config_type"`
	Sequence int `json:"sequence"`
	Url string `json:"url"`
	CartoonId int `json:"cartoon_id"`
	ImgUrl string `json:"img_url"`
}
const (
	MODULE_TYPE_SCROLL    = 1
	MODULE_TYPE_RECOMMEND = 2
	MODULE_TYPE_ELITE     = 3
	MODULE_TYPE_HOT       = 4
	MODULE_TYPE_ENDED     = 5
)
const (
	CONFIG_TYPE_CARTOON   = 1 //漫画
	CONFIG_TYPE_BROADCAST = 2 //系统公告
	CONFIG_TYPE_ADV       = 3 //外部广告
)

func GetConfigRows(moduleType, limit int) (list []HomeConfig) {
	db,_ := DB.OpenCartoon()

	db.Table("home_configs").Select("home_configs.*").Joins("inner join cartoons on home_configs.cartoon_id = cartoons.id").Where("home_configs.deleted_at is null").Where("home_configs.module_type = ?",moduleType).Limit(limit).Scan(&list)
	//common.CheckError(err)
	//for rows.Next() {
	//	var row HomeConfig
	//	err = rows.Scan(&row.ID,&row.ModuleType,&row.ConfigType,&row.Sequence,&row.Url,&row.CartoonId,&row.ImgUrl,&row.DeletedAt,&row.CreatedAt,&row.UpdatedAt)
	//	common.CheckError(err)
	//	list = append(list,row)
	//}
	return list
}


