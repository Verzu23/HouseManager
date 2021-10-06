package models

import "main/gormiti"

type UserManagement struct {
	Users []gormiti.UserGorm `json:"userList"`
	Roles []gormiti.RoleGorm `json:"roleList"`
}
