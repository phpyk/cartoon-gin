package dao

import (
	"cartoon-gin/DB"
	"cartoon-gin/common"
	"encoding/json"
	"strings"
)

const (
	//审核状态
	CARTOON_VERIFY_STATUS_UNCHECK = 0
	CARTOON_VERIFY_STATUS_PASS    = 1
	CARTOON_VERIFY_STATUS_DENY    = 2
	//上架状态
	CARTOON_IS_ON_SALE  = 1
	CARTOON_IS_NOT_SALE = 0
)

type Cartoon struct {
	MyGormModel
	CartoonName    string `json:"cartoon_name"`
	HoverImage     string `json:"hover_image"`
	Author         string `json:"author"`
	Tags           string `json:"tags"`
	CartoonType    int    `json:"cartoon_type"`
	ExternalUrl    string `json:"external_url"`
	Depiction      string `json:"depiction"`
	LatestChapter  int    `json:"latest_chapter"`
	FreeType       int    `json:"free_type"`
	RecommLevel    int    `json:"recomm_level"`
	IsEnd          int    `json:"is_end"`
	VerifyStatus   int    `json:"verify_status"`
	IsOnSale       int    `json:"is_on_sale"`
	KeywordsIds    string `json:"keywords_ids"`
	CatIds         string `json:"cat_ids"`
	OnSaleTime     int    `json:"on_sale_time"`
	Source         int    `json:"source"`
	ReadCount      int    `json:"read_count"`
	OriginalBookId int    `json:"original_book_id"`
	IsRated        int    `json:"is_rated"`
}

type QueryObj struct {
	Id            int           `json:"cartoon_id"`
	CartoonName   string        `json:"cartoon_name"`
	HoverImage    string        `json:"hover_image"`
	Author        string        `json:"author"`
	Tags          string        `json:"tags"`
	ExternalUrl   string        `json:"external_url"`
	IsEnd         int           `json:"is_end"`
	LatestChapter int           `json:"latest_chapter"`
	KeywordsIds   string        `json:"keywords_ids"`
	CreatedAt     common.MyTime `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt     common.MyTime `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

func FindCartoonById(id int) (cartoon Cartoon) {
	db, _ := DB.OpenCartoon()
	db.Where("id = ?", id).First(&cartoon)
	return cartoon
}

func GetCartoonCount() (count int) {
	db, _ := DB.OpenCartoon()
	db.Table("cartoons").
		Where("verify_status = ?", CARTOON_VERIFY_STATUS_PASS).Count(&count)
	return count
}

func GetCartoonRank(page, pageSize int, sortBy, sort string) []map[string]interface{} {
	db, _ := DB.OpenCartoon()
	var list []QueryObj
	//columns := "hover_image, cartoon_name, author, latest_chapter, is_end, keywords_ids, tags, id as cartoon_id, created_at, updated_at"
	db.Table("cartoons").
		Select("cartoons.*").
		Where("verify_status = ?", CARTOON_VERIFY_STATUS_PASS).
		Order(sortBy + " " + sort).
		Limit(pageSize).Offset((page - 1) * pageSize).Scan(&list)

	var result [](map[string]interface{})
	for _, row := range list {
		var item map[string]interface{}
		jsonb, err := json.Marshal(row)
		common.CheckError(err)
		err = json.Unmarshal(jsonb, &item)

		tagArr := strings.Split(row.Tags, ",")
		item["tags"] = tagArr
		result = append(result, item)
	}
	return result
}
