package dao

import "cartoon-gin/DB"

const (
	//审核状态
	CARTOON_VERIFY_STATUS_UNCHECK = 0
	CARTOON_VERIFY_STATUS_PASS = 1
	CARTOON_VERIFY_STATUS_DENY = 2
	//上架状态
	CARTOON_IS_ON_SALE = 1
	CARTOON_IS_NOT_SALE = 0
)
type Cartoon struct {
	MyGormModel
	CartoonName string `json:"cartoon_name"`
	HoverImage string `json:"hover_image"`
	Author string `json:"author"`
	Tags string `json:"tags"`
	Depiction string `json:"depiction"`
	LatestChapter int `json:"latest_chapter"`
	FreeType int `json:"free_type"`
	RecommLevel int `json:"recomm_level"`
	IsEnd int `json:"is_end"`
	VerifyStatus int `json:"verify_status"`
	IsOnSale int `json:"is_on_sale"`
	KeywordsIds string `json:"keywords_ids"`
	CatIds string `json:"cat_ids"`
	OnSaleTime int `json:"on_sale_time"`
	Source int `json:"source"`
	ReadCount int `json:"read_count"`
	OriginalBookId int `json:"original_book_id"`
	IsRated int `json:"is_rated"`
}

func FindCartoonById(id int) (cartoon Cartoon) {
	db,_ := DB.OpenCartoon()
	db.Where("id = ?",id).First(&cartoon)
	return cartoon
}