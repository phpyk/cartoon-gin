package utils

type RecommendCalculator struct {
	UserType     int
	CanShowRated bool
	IsVerifying  bool
}

const (
	VisitorUser = iota
	NormalUser
	WillingUser
	TargetUser
)

func (r *RecommendCalculator) SetCanShowRated(showRated bool) {
	r.CanShowRated = showRated
}

func (r *RecommendCalculator) SetAppVerifyStatus(isVerifying bool) {
	r.IsVerifying = isVerifying
}

func (r *RecommendCalculator) SetUserType(userType int) {
	r.UserType = userType
}

func (r *RecommendCalculator) GetRatedCount(totalCount int) (count int) {
	//总开关关闭、正在审核
	if !r.CanShowRated || r.IsVerifying {
		count = 0
	} else if totalCount == 1 { //数量为1 ==> 广告弹窗
		if InNight() {
			count = 1
		}
	} else {
		//数量大于1 ==> 为你推荐
		if r.UserType == TargetUser {
			if InNight() {
				count = 2
			} else {
				count = 1
			}
		} else if r.UserType == WillingUser {
			if InNight() {
				count = 1
			}
		}
	}
	return count
}
