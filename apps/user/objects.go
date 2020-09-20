package user

import "time"

type UserCreateObject struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
	Gender   string `json:"gender"`
}

type UserObject struct {
	ID       string    `json:"id"`
	Password string    `json:"password"`
	Name     string    `json:"name"`
	Gender   string    `json:"gender"`
	Email    string    `json:"email"`
	Role     string    `json:"role"`
	Activate bool      `json:"activate"`
	JoinTime time.Time `json:"joinTime"`
}
