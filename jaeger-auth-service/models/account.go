package models

import (
	"time"

	"github.com/google/uuid"
)

type Account struct {
	Id              *uuid.UUID `gorm:"primaryKey;index;type:uuid;default:uuid_generate_v4()" json:"id" db:"id"`
	UserId          *uuid.UUID `gorm:"index;type:uuid" json:"user_id" db:"user_id"`
	Username        string     `gorm:"size:255" json:"username" db:"username"`
	PassPlainText   string     `gorm:"column:password" json:"password" db:"password"`
	PassBcrypt      string     `gorm:"-" json:"-"`
	Status          string     `gorm:"type:account_status" json:"status" db:"status"`
	WebAccess       string     `gorm:"type:web_access" json:"web_access" db:"web_access"`
	RevokeTokenCode string     `json:"revoke_token_code" db:"revoke_token_code"`
	CreatedBy       string     `gorm:"size:50" json:"created_by" db:"created_by"`
	UpdatedBy       *string    `gorm:"size:50" json:"updated_by" db:"updated_by"`
	CreatedAt       time.Time  `json:"created_at" db:"created_at"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at"`
}

func (Account) TableName() string {
	return "account"
}
