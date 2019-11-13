package dao

import (
	"cartoon-gin/DB"
	"cartoon-gin/utils"
	"fmt"
	"strconv"

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
	//免费类型
	CARTOON_FREE_TYPE_FREE = 0; //所有人免费
	CARTOON_FREE_TYPE_VIP = 1; //vip免费
	CARTOON_FREE_TYPE_NOT = 2; // 不免费，收费
	//完结状态
	CARTOON_IS_END_YES = 1;
	CARTOON_IS_END_NO = 0;
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
	Id            int          `json:"cartoon_id"`
	CartoonName   string       `json:"cartoon_name"`
	HoverImage    string       `json:"hover_image"`
	Author        string       `json:"author"`
	Tags          string       `json:"tags"`
	ExternalUrl   string       `json:"external_url"`
	IsEnd         int          `json:"is_end"`
	LatestChapter int          `json:"latest_chapter"`
	KeywordsIds   string       `json:"keywords_ids"`
	FreeType	  int 			`json:"free_type"`
	CreatedAt     utils.MyTime `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt     utils.MyTime `json:"updated_at" time_format:"2006-01-02 15:04:05"`
}

type SearchRequest struct {
	IsFree string `form:"is_free" json:"is_free"` //要判断值为0和空的情况，如果类型为int，绑定无法区分
	IsEnd string `form:"is_end" json:"is_end"`
	CatId int `form:"cat_id" json:"cat_id"`
	Keywords string `form:"keywords" json:"keywords"`
	SortType int `form:"sort_type" json:"sort_type"`
	PageSize int `form:"page_size" json:"page_size"`
	Page int `form:"page" json:"page"`
}

func GetCartoonById(id int) (cartoon Cartoon) {
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
	return formatQueryObj(list)
}

func FindCartoonsHoverImageForUpload(limit,lastMaxId int) []Cartoon {
	db, _ := DB.OpenCartoon()
	var list []Cartoon
	db.Table("cartoons").Where("id > ?",lastMaxId).Where("hover_image like ?","%qiniu.tblinker%").Order("id ASC").Limit(limit).Scan(&list)
	return list
}

func SearchCartoonByConditions(request SearchRequest) []map[string]interface{} {
	db,_ := DB.OpenCartoon()
	columns := "id as cartoon_id, cartoon_name, hover_image, author, is_end, latest_chapter, free_type, external_url"
	query := db.Debug().Table("cartoons").
		Select(columns).
		Where("verify_status = ?",CARTOON_VERIFY_STATUS_PASS).
		Where("is_on_sale = ?",CARTOON_IS_ON_SALE)
	if request.IsFree != "" {
		if request.IsFree == "1" {
			query = query.Where("free_type = ?",CARTOON_FREE_TYPE_FREE)
		}else {
			query = query.Where("free_type in (?)",[]int{CARTOON_FREE_TYPE_VIP,CARTOON_FREE_TYPE_NOT})
		}
	}
	if request.IsEnd != "" {
		if request.IsEnd == "1" {
			query = query.Where("is_end = ?",CARTOON_IS_END_YES)
		}else {
			query = query.Where("is_end = ?",CARTOON_IS_END_NO)
		}
	}
	if request.CatId > 1 {
		query = query.Where("FIND_IN_SET("+strconv.Itoa(request.CatId)+",cat_ids)")
	}
	keywords := strings.Trim(request.Keywords," ")
	fmt.Println("keywords: ",keywords)
	if keywords != "" {
		query = query.Where("cartoon_name like '%"+keywords+"%' or FIND_IN_SET('"+keywords+"',tags)")
	}
	orderByColumn := "read_count"
	if request.SortType == 2 {
		orderByColumn = "updated_at"
	}
	if request.PageSize == 0 {
		request.PageSize = 20
	}
	var list []QueryObj
	query.Order(orderByColumn + " desc").
		Limit(request.PageSize).
		Offset(request.Page * request.PageSize).
		Scan(&list)
	return formatQueryObj(list)
}

func formatQueryObj(list []QueryObj) []map[string]interface{} {
	var result [](map[string]interface{})
	for _, row := range list {
		var item map[string]interface{}
		jsonb, err := json.Marshal(row)
		utils.CheckError(err)
		err = json.Unmarshal(jsonb, &item)

		tagArr := strings.Split(row.Tags, ",")
		item["tags"] = tagArr
		result = append(result, item)
	}
	return result
}