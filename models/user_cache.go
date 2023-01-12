package models

import (
	"chat/utils"
	"time"
)

func SetUserOnlineInfo(key string, val []byte, timeTTL time.Duration) {
	utils.Redis.Set(key, val, timeTTL)
}
