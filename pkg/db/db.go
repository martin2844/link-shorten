package db

import (
	"database/sql/driver"
	"fmt"
	"time"

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

type TimeWrapper struct {
	time.Time
}

func (t *TimeWrapper) Scan(value interface{}) error {
	switch v := value.(type) {
	case time.Time:
		t.Time = v
	case []byte:
		strVal := string(v)
		parsedTime, err := time.Parse("2006-01-02 15:04:05.000", strVal)
		if err != nil {
			return fmt.Errorf("failed to parse time from: %v", strVal)
		}
		t.Time = parsedTime
	default:
		return fmt.Errorf("unsupported type: %T", value)
	}
	return nil
}

func (t TimeWrapper) Value() (driver.Value, error) {
	return t.Time, nil
}

type Link struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt TimeWrapper
	UpdatedAt TimeWrapper
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Original  string         `gorm:"type:varchar(191);unique"`
	Short     string         `gorm:"type:varchar(191);unique"`
	ViewCount uint           `gorm:"default:0"`
}
