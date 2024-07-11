package model

import (
	"time"

	"github.com/google/uuid"
)

func (r *User) TableName() string {
	return "users"
}

type User struct {
	ID        uuid.UUID `gorm:"primaryKey;type:uuid;default:uuid_generate_v4()" json:"id"`
	Email     string    `gorm:"index;size:100;unique;not null"`
	FullName  string    `gorm:"column:full_name;not null"`
	CreatedAt time.Time `gorm:"column:created_at;default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
}
