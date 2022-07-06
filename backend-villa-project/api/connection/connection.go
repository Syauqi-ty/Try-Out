package connection

import (
	"villa-akmali/config"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func Create() *gorm.DB {
	config.Init()
	// var (
	// 	user     string = viper.GetString(`database.user`)
	// 	host     string = viper.GetString(`database.host`)
	// 	name     string = viper.GetString(`database.name`)
	// 	port     string = viper.GetString(`database.port`)
	// 	password string = viper.GetString(`database.password`)
	// )

	// dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn),&gorm.Config{})
	dsn := "root:villaakmalidev@tcp(akmalidb:3306)/villa-akmali?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil{
		panic("Failed to connect")
	}
	return db
}