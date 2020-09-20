package db

import "github.com/jinzhu/gorm"

var DB *gorm.DB

func init() {
	var err error
	DB, err = gorm.Open("mysql", "root:654321@tcp(localhost:3306)/demo?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
}
