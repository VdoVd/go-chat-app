package models

import (
	"github.com/jinzhu/gorm"
	"time"
)

type UserBasic struct {
	gorm.Model
	Name          string
	PassWord      string
	Phone         string
	Email         string
	Identity      string
	ClientIp      string
	ClientPort    string
	LoginTime     time.Time
	HeartbeatTime time.Time
	IsLogout      bool
	LoginOutTime  time.Time `gorm:"column:login_out_time json:"login_out_time`
	DeviceInfo    string
}

func (table *UserBasic) TableName() string {
	return "user_basic"
}
