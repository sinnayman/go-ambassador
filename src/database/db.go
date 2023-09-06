package database

import (
	"ambassador/src/models"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Database struct {
	*gorm.DB
}

func NewDatabase() (*Database, error) {
	var err error

	for retries := 1; retries <= 10; retries++ {
		gormDB, err := gorm.Open(mysql.Open("ambassador:ambassador_secret@tcp(db:3306)/ambassador?parseTime=true"), &gorm.Config{})

		if err == nil {
			db := &Database{gormDB}
			return db, nil
		}

		fmt.Printf("Attempt %d: Could not connect to db, retrying in 5 seconds, error was %v\n", retries, err)
		time.Sleep(5 * time.Second)
	}

	return nil, err
}

func (d *Database) Close() {
	sqlDB, _ := d.DB.DB()
	sqlDB.Close()
}

func (d *Database) AutoMigrate() {
	d.DB.AutoMigrate(models.UserWrite{},  models.ProductWrite{})
}
