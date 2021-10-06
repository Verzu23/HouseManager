package gormiti

import "time"

type BollettaGorm struct {
	IdBolletta            int            `gorm:"primaryKey;autoIncrement:true;not null"`
	Servizio              string         `gorm:"size:100;not null"`
	Fornitore             string         `gorm:"size:100;not null"`
	Descrizione           string         `gorm:"size:255;not null"`
	PeriodoDal            time.Time      `gorm:"not null"`
	PeriodoAl             time.Time      `gorm:"not null"`
	Locazione_IdLocazione int            `gorm:"not null"`
	Locazione             *LocazioneGorm `gorm:"foreignKey:Locazione_IdLocazione;references:IdLocazione"`
	CreatedAt             time.Time      `gorm:"autoCreateTime;not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	UpdatedAt             time.Time      `gorm:"autoUpdateTime;not null;default:CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP;type:TIMESTAMP"`

	/* Locazione_casa_gorm_id_casa           int            `gorm:"not null"`
	Casa                                  []LocazioneGorm `gorm:"foreignKey:Locazione_casa_gorm_id_casa;references:Casa_gorm_id_casa"`
	Locazione_inquilino_gorm_id_inquilino int            `gorm:"not null"`
	Inquilino                             []LocazioneGorm `gorm:"foreignKey:Locazione_inquilino_gorm_id_inquilino;references:Inquilino_gorm_id_inquilino"`
	*/
	/* Locazione_casa_idcasa           int `gorm:"not null"`
	Locazione_inquilino_idinquilino int `gorm:"not null"`
	Locazione_inquilino_user_iduser int `gorm:"not null"`

	Casa      *LocazioneGorm `gorm:"foreignKey:Locazione_casa_idcasa;references:Casa_idcasa"`
	Inquilino *LocazioneGorm `gorm:"foreignKey:Locazione_inquilino_idinquilino;references:Inquilino_idinquilino"`
	User      *LocazioneGorm `gorm:"foreignKey:Locazione_inquilino_user_iduser;references:Inquilino_user_iduser"` */
}

func (BollettaGorm) TableName() string {
	return "bolletta"
}
