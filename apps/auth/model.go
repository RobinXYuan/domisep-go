package auth

import (
	"domisep/db"

	"github.com/jinzhu/gorm"
)

func init() {
	migrateAll()
}

func migrateAll() {
	db.DB.AutoMigrate(&Role{})
}

//Role 角色表
type Role struct {
	gorm.Model
	Name        string `gorm:"type:VARCHAR;size:40;unique;not null"`
	Permissions string `gorm:"type:VARCHAR;size:255;not null"`
}
