package gormiti

import (
	"time"
)

type AccessTokenGorm struct {
	Id        int    `gorm:"primaryKey;autoIncrement:true;not null"`
	Uuid      string `gorm:"size:500"`
	Username  string `gorm:"size:100"`
	ExpiresAt time.Time
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
}

func (AccessTokenGorm) TableName() string {
	return "access_tokens"
}
