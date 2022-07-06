package connection

import (
	config "studybuddy-backend-fast/config"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func Create() *gorm.DB {
	config.Init()

	var (
		user     string = viper.GetString(`database.user`)
		host     string = viper.GetString(`database.host`)
		name     string = viper.GetString(`database.name`)
		port     string = viper.GetString(`database.port`)
		password string = viper.GetString(`database.password`)
	)

	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// dsn := "root:9thT83srsS.r4/5KExOsJ@tcp(studybuddy_db_1:3306)/studybuddy?charset=utf8mb4&parseTime=True&loc=Local"
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// db, err := gorm.Open(mysql.Open("root:9thT83srsS.r4/5KExOsP@tcp(mysql:3306)/studybuddy?charset=utf8&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("Failed to connect to database")
	}
	return db
}
