package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	Id           *uuid.UUID `gorm:"primaryKey;index;type:uuid;default:uuid_generate_v4()" json:"id" db:"id"`
	Username     string     `gorm:"size:255" json:"username" db:"username"`
	TitleNameTH  string     `gorm:"size:255" json:"title_name_th" db:"title_name_th"`
	FirstNameTH  string     `gorm:"size:255" json:"first_name_th" db:"first_name_th"`
	LastNameTH   string     `gorm:"size:255" json:"last_name_th" db:"last_name_th"`
	TitleNameEN  string     `gorm:"size:255" json:"title_name_en" db:"title_name_en"`
	FirstNameEN  string     `gorm:"size:255" json:"first_name_en" db:"first_name_en"`
	LastNameEN   string     `gorm:"size:255" json:"last_name_en" db:"last_name_en"`
	Phone        string     `gorm:"size:10" json:"phone" db:"phone"`
	Email        string     `gorm:"size:50" json:"email" db:"email"`
	BirthdayDate time.Time `json:"birthday_date" db:"birthday_date"`
	Gender       string     `gorm:"size:10" json:"gender" db:"gender"`
	Nationality  string     `gorm:"size:20" json:"nationality" db:"nationality"`
	Status       string     `gorm:"type:user_status" json:"status" db:"status"`
	CreatedBy    string     `gorm:"size:50" json:"created_by" db:"created_by"`
	UpdatedBy    *string    `gorm:"size:50" json:"updated_by" db:"updated_by"`
	DeletedBy    *string    `gorm:"size:50" json:"deleted_by" db:"deleted_by"`
	CreatedAt    time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at"`
	DeletedAt    *time.Time `json:"deleted_at" db:"deleted_at"`
}

func (User) TableName() string {
	return "user"
}

func (u *User) GenUUID() {
	id := uuid.New()
	u.Id = &id
}
