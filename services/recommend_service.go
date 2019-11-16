package services

import "cartoon-gin/dao"

type RecommendService struct {
	User dao.User
	RateScale float32
}

func (r RecommendService) SetUser(user dao.User)  {
	r.User = user
}

func (r RecommendService) initScale() {
	//userType := r.User.UserType

}