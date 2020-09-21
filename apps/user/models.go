package user

import (
	"github.com/jinzhu/gorm"
)

//User 用户表
type User struct {
	gorm.Model
	Name     string `gorm:"type=VARCHAR,size=80;"`
	Password string `gorm:"type=VARCHAR,NOT NULL;"`
	Gender   string `gorm:"gender"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Activate uint8  `json:"activate"`
}
