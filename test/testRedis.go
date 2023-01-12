package mq

import (
	"chat/utils"
	"context"
	"fmt"
	"testing"
	"time"
)

var ctx context.Context

func init() {
	ctx = context.Background()
}

func TestPublish(t *testing.T) {
	msg := "当前时间:" + time.Now().Format("15.04.05")
	err := utils.Publish(utils.PublishKey, msg)
	if err != nil {
		fmt.Println(err)
	}
}
