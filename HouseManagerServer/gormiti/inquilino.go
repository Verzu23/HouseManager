package gormiti

import "time"

type InquilinoGorm struct {
	IdInquilino int        `gorm:"primaryKey;autoIncrement:true;not null"`
	Nome        string     `gorm:"size:100;not null"`
	Cognome     string     `gorm:"size:100;not null"`
	Attivo      bool       `gorm:"not null"`
	Casa        []CasaGorm `gorm:"many2many:locazione;"`
	CreatedAt   time.Time  `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt   time.Time  `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
}

func (InquilinoGorm) TableName() string {
	return "inquilino"
}
