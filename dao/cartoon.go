package dao

import (
	"cartoon-gin/DB"
	"cartoon-gin/utils"
	"strconv"

	"encoding/json"
	"strings"
)

const (
	//审核状态
	CartoonVerifyStatusUncheck = iota
	CartoonVerifyStatusPass
	CartoonVerifyStatusDeny
	//上架状态
	CartoonIsOnSale  = 1
	CartoonIsNotSale = 0
	//免费类型
	CartoonFreeTypeFree = iota //所有人免费
	CartoonFreeTypeVip         //vip免费
	CartoonFreeTypeNot         // 不免费，收费
	//完结状态
	CartoonIsEndYes = 1
	CartoonIsEndNo  = 0
	//cartoon类型
	CartoonTypeNormal = 1
	CartoonTypeHM = 2
	CartoonTypeExternal = 3
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

type QueryCartoons struct {
	Id            int          `json:"cartoon_id"`
	CartoonName   string       `json:"cartoon_name"`
	HoverImage    string       `json:"hover_image"`
	Author        string       `json:"author"`
	Tags          string       `json:"tags"`
	ExternalUrl   string       `json:"external_url"`
	IsEnd         int          `json:"is_end"`
	LatestChapter int          `json:"latest_chapter"`
	KeywordsIds   string       `json:"keywords_ids"`
	FreeType      int          `json:"free_type"`
	CreatedAt     utils.MyTime `json:"created_at" time_format:"2006-01-02 15:04:05"`
	UpdatedAt     utils.MyTime `json:"updated_at" time_format:"2006-01-02 15:04:05"`
	LastReadChapterId int `json:"last_read_chapter_id"`
	LastReadTime int `json:"last_read_time"`
}

type SearchRequest struct {
	IsFree   string `form:"is_free" json:"is_free"` //要判断值为0和空的情况，如果类型为int，绑定无法区分
	IsEnd    string `form:"is_end" json:"is_end"`
	CatId    int    `form:"cat_id" json:"cat_id"`
	Keywords string `form:"keywords" json:"keywords"`
	SortType int    `form:"sort_type" json:"sort_type"`
	PerPage  int    `form:"page_size" json:"page_size"`
	Page     int    `form:"page" json:"page"`
}

func GetCartoonById(id int) (cartoon Cartoon) {
	db, _ := DB.OpenCartoon()
	db.Where("id = ?", id).First(&cartoon)
	return cartoon
}

func CartoonExists(id int) bool {
	db, _ := DB.OpenCartoon()
	count := 0
	db.Where("id = ?", id).Where("is_on_sale = ?", CartoonIsOnSale).Where("verify_status = ?", CartoonVerifyStatusPass).Count(&count)
	return count > 0
}

func GetCartoonCount() (count int) {
	db, _ := DB.OpenCartoon()
	db.Table("cartoons").
		Where("verify_status = ?", CartoonVerifyStatusPass).Count(&count)
	return count
}

func GetCartoonRank(page, pageSize int, sortBy, sort string) []map[string]interface{} {
	db, _ := DB.OpenCartoon()
	var list []QueryCartoons
	//columns := "hover_image, cartoon_name, author, latest_chapter, is_end, keywords_ids, tags, id as cartoon_id, created_at, updated_at"
	db.Table("cartoons").
		Select("cartoons.*").
		Where("verify_status = ?", CartoonVerifyStatusPass).
		Order(sortBy + " " + sort).
		Limit(pageSize).Offset((page - 1) * pageSize).Scan(&list)
	return formatQueryCartoons(list)
}

func FindCartoonsHoverImageForUpload(limit, lastMaxId int) []Cartoon {
	db, _ := DB.OpenCartoon()
	var list []Cartoon
	db.Table("cartoons").Where("id > ?", lastMaxId).Where("hover_image like ?", "%qiniu.tblinker%").Order("id ASC").Limit(limit).Scan(&list)
	return list
}

func SearchCartoonByConditions(request SearchRequest) []map[string]interface{} {
	db, _ := DB.OpenCartoon()
	columns := "id, cartoon_name, hover_image, author, is_end, latest_chapter, tags, free_type, external_url, created_at, updated_at"
	query := db.Debug().Table("cartoons").
		Select(columns).
		Where("verify_status = ?", CartoonVerifyStatusPass).
		Where("is_on_sale = ?", CartoonIsOnSale)
	if request.IsFree != "" {
		if request.IsFree == "1" {
			query = query.Where("free_type = ?", CartoonFreeTypeFree)
		} else {
			query = query.Where("free_type in (?)", []int{CartoonFreeTypeVip, CartoonFreeTypeNot})
		}
	}
	if request.IsEnd != "" {
		if request.IsEnd == "1" {
			query = query.Where("is_end = ?", CartoonIsEndYes)
		} else {
			query = query.Where("is_end = ?", CartoonIsEndNo)
		}
	}
	if request.CatId > 1 {
		query = query.Where("FIND_IN_SET(" + strconv.Itoa(request.CatId) + ",cat_ids)")
	}
	keywords := strings.Trim(request.Keywords, " ")
	keywords = utils.FilterSpecialChar(keywords)
	if keywords != "" {
		query = query.Where("cartoon_name like '%" + keywords + "%' or FIND_IN_SET('" + keywords + "',tags)")
		//query = query.Where("cartoon_name like ?","%"+keywords+"%")
	}
	orderByColumn := "read_count"
	if request.SortType == 2 {
		orderByColumn = "updated_at"
	}
	if request.PerPage == 0 {
		request.PerPage = 18
	}
	var list []QueryCartoons
	query.Order(orderByColumn + " desc").
		Limit(request.PerPage).
		Offset((request.Page - 1) * request.PerPage).
		Scan(&list)
	return formatQueryCartoons(list)
}

func formatQueryCartoons(list []QueryCartoons) []map[string]interface{} {
	var result [](map[string]interface{})
	for _, row := range list {
		var item map[string]interface{}
		jsonb, err := json.Marshal(row)
		utils.CheckError(err)
		err = json.Unmarshal(jsonb, &item)

		item["tags"] = utils.GetTagsArray(row.Tags, 2)
		result = append(result, item)
	}
	return result
}
