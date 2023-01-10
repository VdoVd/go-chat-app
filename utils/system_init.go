package utils

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB
var Redis *redis.Client

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitRedis() {
	Redis = redis.NewClient(&redis.Options{

		Addr:         viper.GetString("redis.addr"),
		Password:     viper.GetString("redis.password"),
		DB:           viper.GetInt("redis.DB"),
		PoolSize:     viper.GetInt("redis.poolSize"),
		MinIdleConns: viper.GetInt("redis.minIdleConn"),
	})
	pong, err := Redis.Ping().Result()
	if err != nil {
		fmt.Println("init redis...", err)
	} else {
		fmt.Println("MySQL inited....", pong)
	}
}
func InitMySQL() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold: time.Second, //慢SQL阈值
			LogLevel:      logger.Info, //级别
			Colorful:      true,        //彩色
		},
	)
	DB, _ = gorm.Open(mysql.Open(viper.GetString(`mysql.dns`)), &gorm.Config{Logger: newLogger})
	fmt.Println(" MySQL inited 。。。。")
	//user := models.UserBasic{}
	//DB.Find(&user)
	//fmt.Println(user)
}

const (
	PublishKey = "websocket"
)

//Publish send message to Redis
func Publish(channel string, msg string) error {
	var err error
	fmt.Println("Publish ....", msg)
	err = Redis.Publish(channel, msg).Err()
	if err != nil {
		fmt.Println(err)
	}
	return err
}

//Subscriber subscribe massagen to redis
func Subscribe(channel string) (string, error) {
	sub := Redis.Subscribe(channel)
	fmt.Println("Subscribe ....")
	msg, err := sub.ReceiveMessage()
	if err != nil {
		fmt.Println(err)
		return "", err
	}
	fmt.Println("Subscribe ....", msg.Payload)
	return msg.Payload, err
}
