package gormiti

import "time"

type CasaGorm struct {
	IdCasa    int    `gorm:"primaryKey;autoIncrement:true;not null"`
	Indirizzo string `gorm:"size:100;not null"`
	Civico    string `gorm:"size:100;not null"`
	Nome      int    `gorm:"not null"`
	Locatore  int
	CreatedAt time.Time `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt time.Time `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	//Inquilini []InquilinoGorm `gorm:"many2many:locazione;"`
}

func (CasaGorm) TableName() string {
	return "casa"
}
