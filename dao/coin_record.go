package dao

type CoinRecord struct {
	MyGormModel
	UserId int `json:"user_id"`
	Amount int `json:"amount"`
	Balance int `json:"balance"`
	ActType int `json:"act_type"`
	CartoonId int `json:"cartoon_id"`
	ChapterId int `json:"chapter_id"`
	BusinessType int `json:"business_type"`
	Remark string `json:"remark"`
	ReferBizId int `json:"refer_biz_id"`
}

const (
	CoinActTypeDecrease = 1
	CoinActTypeIncrease = 2
)