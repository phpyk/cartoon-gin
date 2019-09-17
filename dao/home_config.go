package dao

import (
	"cartoon-gin/DB"
	"cartoon-gin/common"
	"encoding/json"
	"strings"
)

type HomeConfig struct {
	MyGormModel
	ModuleType int    `json:"module_type"`
	ConfigType int    `json:"config_type"`
	Sequence   int    `json:"sequence"`
	Url        string `json:"url"`
	CartoonId  int    `json:"cartoon_id"`
	ImgUrl     string `json:"img_url"`
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

func GetHomeConfigRows(moduleType, limit int) []map[string]interface{} {
	db, _ := DB.OpenCartoon()
	var list []HomeConfig
	db.Table("home_configs").Select("home_configs.*").Joins("left join cartoons on home_configs.cartoon_id = cartoons.id").Where("home_configs.deleted_at is null").Where("home_configs.module_type = ?", moduleType).Limit(limit).Scan(&list)

	var result [](map[string]interface{})
	for _, row := range list {
		//此处直接使用structs.Map(row),转换后key都是大写：eg. CartoonId,CreatedAt
		//item := structs.Map(row)
		//所以先将row转为json，再反序列化，生成的key为小写
		jsonb, err := json.Marshal(row)
		common.CheckError(err)

		var item map[string]interface{}
		err = json.Unmarshal(jsonb, &item)
		common.CheckError(err)

		if row.ConfigType == CONFIG_TYPE_CARTOON && row.CartoonId > 0 {
			cartoon := FindCartoonById(row.CartoonId)
			item["hover_image"] = cartoon.HoverImage
			item["cartoon_name"] = cartoon.CartoonName
			item["latest_chapter"] = cartoon.LatestChapter
			item["is_end"] = cartoon.IsEnd
			item["tags"] = strings.Split(cartoon.Tags, ",")
		}
		result = append(result, item)
	}
	return result
}
func GetMoreHomeConfigRows(moduleName string, page, pageSize int) []map[string]interface{} {
	db, _ := DB.OpenCartoon()
	var list []HomeConfig
	moduleType := mappingModuleType(moduleName)

	columns := "cartoons.hover_image," +
		"cartoons.cartoon_name," +
		"cartoons.author," +
		"cartoons.latest_chapter," +
		"home_configs.cartoon_id," +
		"cartoons.tags," +
		"cartoons.is_end," +
		"cartoons.keywords_ids"
	db.Table("home_configs").Select(columns).Joins("left join cartoons on home_configs.cartoon_id = cartoons.id").Where("home_configs.deleted_at is null").Where("home_configs.module_type = ?", moduleType).Where("cartoons.verify_status = ?",CARTOON_VERIFY_STATUS_PASS).Limit(pageSize).Offset((page-1)*pageSize).Scan(&list)

	//TODO
	return nil
}

func mappingModuleType(moduleName string) int {
	var t int
	switch moduleName {
	case "recommend":
		t = 2
	case "elite":
		t = 3
	case "hot":
		t = 4
	case "ended":
		t = 5
	default:
		t = 2
	}
	return t
}
