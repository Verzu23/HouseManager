package gormiti

import "time"

type UserGorm struct {
	Id                    int            `gorm:"primaryKey;autoIncrement:true;not null"`
	Username              string         `gorm:"size:100;not null"`
	Password              string         `gorm:"size:100;not null"`
	Inquilino_IdInquilino int            `gorm:"not null"`
	Inquilino             *InquilinoGorm `gorm:"foreignKey:Inquilino_IdInquilino"`
	CreatedAt             time.Time      `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
}

func (UserGorm) TableName() string {
	return "user"
}
