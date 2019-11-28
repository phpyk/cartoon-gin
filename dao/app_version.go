package dao

import (
	"time"

	"cartoon-gin/DB"
	"cartoon-gin/utils"
)

type AppVersion struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time	`time_format:"2006-01-02 15:04:05"`
	UpdatedAt time.Time `time_format:"2006-01-02 15:04:05"`
	IsForcedUpdate int `json:"is_forced_update"`
	Content string `json:"content"`
	Version string `json:"version"`
	DeviceType string `json:"device_type"`
	Channel string `json:"channel"`
	DownloadUrl string `json:"download_url"`
	IsVerifying int `json:"is_verifying"`
	ShowRated int `json:"show_rated"`
	ShowUpdateTips int `json:"show_update_tips"`
}


func GetAppVersionRow(version,deviceType,channel string) (row AppVersion) {
	db,err := DB.OpenCartoon()
	utils.CheckError(err)
	query := db.Table("app_versions").Where("device_type = ?",deviceType).Where("version = ?",version)
	if channel != "" {
		query = query.Where("channel = ?",channel)
	}
	query.First(&row)
	return row
}

