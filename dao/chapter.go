package dao

import (
	"cartoon-gin/DB"
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
}

func UpdateChapterCoverImage(row Chapter,url string) bool {
	db, _ := DB.OpenCartoon()
	db.Model(&row).Updates(Chapter{HoverImage:url})
	return true
}

func GetChapterRow(cid string) Chapter {
	db, _ := DB.OpenCartoon()
	var row Chapter
	db.Table("cartoon_chapters").Where("id = ?",cid).First(&row)
	fmt.Printf("%+v",row)
	return row
}