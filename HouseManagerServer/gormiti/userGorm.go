package gormiti

import (
	"time"
)

type UserGorms struct {
	Id           int    `gorm:"primaryKey;autoIncrement:true;not null"`
	Username     string `gorm:"size:100"`
	Password     string `gorm:"size:100" json:"-"`
	IdRole       int
	UserId       int
	Name         string
	Surname      string
	Email        string
	IdTeam       int
	ValidityFrom time.Time `gorm:"not null;default:CURRENT_TIMESTAMP;type:TIMESTAMP"`
	ValidityTo   time.Time `gorm:"not null;type:TIMESTAMP"`
	Role         *RoleGorm `gorm:"foreignKey:IdRole;references:RoleId"`
}

func (UserGorms) TableName() string {
	return "usesdr"
}
