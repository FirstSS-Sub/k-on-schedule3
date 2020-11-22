package config

import (
	"github.com/jinzhu/gorm"
	"os"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open("postgres", os.Getenv("HOGE_HOGE_HOGE"))

	if err != nil {
		panic(err)
	}

	return db
}
