package utils

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
)

var DB *gorm.DB

func InitConfig() {
	viper.SetConfigName("app")
	viper.AddConfigPath("chat\\config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("config app:", viper.Get("app"))
	fmt.Println("config mysql:", viper.Get("mysql"))
}

func InitMySQL() {
	DB, _ = gorm.Open(mysql.Open(viper.GetString("mysql.dns")), &gorm.Config{})

}
