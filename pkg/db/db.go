package db

import (
	"github.com/lynxsecurity/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB

func Init() {
	dsn := viper.GetString("DB")
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}
	//Creates database entry and hoists it to exported var for other packages to interact with DB
	Instance = db
}

func AutoMigrate() {
	// Generates migrations which creates the tables in the database
	Instance.AutoMigrate(&Link{})
}

type Link struct {
	gorm.Model
	Original string `gorm:"unique"` // Original URL
	Short    string `gorm:"unique"`
}
