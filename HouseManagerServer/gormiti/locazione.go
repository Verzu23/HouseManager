package gormiti

import "time"

type LocazioneGorm struct {
	IdLocazione                 int       `gorm:"primaryKey"`
	Casa_gorm_id_casa           int       `gorm:"primaryKey"`
	Inquilino_gorm_id_inquilino int       `gorm:"primaryKey"`
	ValidityFrom                time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	ValidityTo                  time.Time `gorm:"not null;type:TIMESTAMP"`
	CreatedAt                   time.Time `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt                   time.Time `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`
	//QUESTO CREA FK DUPLICATE
	/*Inquilino *InquilinoGorm `gorm:"foreignKey:Inquilino_gorm_id_inquilino;references:IdInquilino"`
	Casa      *CasaGorm      `gorm:"foreignKey:Casa_gorm_id_casa;references:IdCasa"` */

	//User      *InquilinoGorm `gorm:"foreignKey:Inquilino_user_iduser;references:User_iduser"`
}

func (LocazioneGorm) TableName() string {
	return "locazione"
}
