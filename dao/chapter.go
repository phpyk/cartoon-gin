package dao

import (
	"encoding/json"

	"cartoon-gin/DB"
	"cartoon-gin/utils"

	"fmt"
)

type Chapter struct {
	MyGormModel
	CartoonId	int `json:"cartoon_id"`
	ChapterName      string `json:"chapter_name"`
	Sequence  int `json:"sequence"`
	HoverImage string `json:"hover_image"`
	FreeType int `json:"free_type"`
	OriginalPrice int `json:"original_price"`
	SalePrice int `json:"sale_price"`
	IsBuy int `json:"is_buy" gorm:"-"`
}

type QueryChapters struct {
	CartoonId	int `json:"cartoon_id"`
	ChapterName      string `json:"chapter_name"`
	Sequence  int `json:"sequence"`
	HoverImage string `json:"hover_image"`
	FreeType int `json:"free_type"`
	OriginalPrice int `json:"original_price"`
	SalePrice int `json:"sale_price"`
}

func UpdateChapterCoverImage(row Chapter,url string) bool {
	db, _ := DB.OpenCartoon()
	db.Model(&row).Updates(Chapter{HoverImage:url})
	return true
}

func GetChapterRow(chapterId int) Chapter {
	db, _ := DB.OpenCartoon()
	var row Chapter
	db.Table("cartoon_chapters").Where("id = ?", chapterId).First(&row)
	fmt.Printf("%+v",row)
	return row
}

func GetChapterNeighbors(currentChapter Chapter) (last Chapter, next Chapter) {
	db, _ := DB.OpenCartoon()
	db.Table("cartoon_chapters").Where("cartoon_id = ?",currentChapter.CartoonId).Where("sequence < ?",currentChapter.Sequence).Order("sequence DESC").First(&last)
	db.Table("cartoon_chapters").Where("cartoon_id = ?",currentChapter.CartoonId).Where("sequence > ?",currentChapter.Sequence).Order("sequence ASC").First(&next)
	return
}

func GetChaptersCount(cartoonId int) int {
	db, _ := DB.OpenCartoon()
	var count int
	db.Table("cartoon_chapters").Where("cartoon_id = ?",cartoonId).Count(&count)
	return count
}

func FindChaptersForUpload(limit int,lastMaxId int) []Chapter {
	db, _ := DB.OpenCartoon()
	var list []Chapter
	db.Table("cartoon_chapters").Where("id > ?",lastMaxId).Where("hover_image != ?","").Limit(limit).Order("id ASC").Scan(&list)
	return list
}

func GetChapterList(cartoonId, sortType int,paginate bool,page,pageSize int) []map[string]interface{} {
	var sortStr string
	if sortType == 1 {
		sortStr = "ASC"
	}else {
		sortStr = "DESC"
	}
	columns := "id, id as chapter_id, cartoon_id,chapter_name, sequence, hover_image, free_type, original_price, sale_price"
	db, _ := DB.OpenCartoon()
	query := db.Debug().Table("cartoon_chapters").Select(columns).Where("cartoon_id = ?",cartoonId).Order("sequence "+sortStr)
	if paginate {
		query = query.Limit(pageSize).Offset((page-1) * pageSize)
	}
	var list []QueryChapters
	query.Scan(&list)

	return formatQueryChatpers(list)
}

func formatQueryChatpers(list []QueryChapters) []map[string]interface{} {
	var result [](map[string]interface{})
	for _, row := range list {
		var item map[string]interface{}
		jsonb, err := json.Marshal(row)
		utils.CheckError(err)
		err = json.Unmarshal(jsonb, &item)

		result = append(result, item)
	}
	return result
}