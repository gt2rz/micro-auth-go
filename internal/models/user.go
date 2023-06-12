package models

import "time"

type User struct {
	Id                   string    `json:"id"`
	Email                string    `json:"email"`
	Password             string    `json:"password"`
	Firstname            string    `json:"firstname"`
	Lastname             string    `json:"lastname"`
	Phone                string    `json:"phone"`
	Verified             bool      `json:"verified"`
	PasswordResetTokenAt time.Time `json:"passwordResetTokenAt"`
	CreatedAt            string    `json:"createdAt"`
	UpdatedAt            string    `json:"updatedAt"`
}
