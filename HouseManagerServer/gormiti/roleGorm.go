package gormiti

type RoleGorm struct {
	Id       int    `gorm:"primaryKey;autoIncrement:true;not null"`
	RoleId   int    `gorm:"uniqueIndex"`
	Rolename string `gorm:"size:100"`
}

func (RoleGorm) TableName() string {
	return "usr_role"
}
