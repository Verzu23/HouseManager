package models

type User struct {
	Username string `json:"Username" binding:"required"`
	Password string `json:"Password,omitempty" binding:"required"`
	Name     string `json:"Name"`
	Surname  string `json:"Surname"`
	Email    string `json:"Email"`
	Rolename string `json:"Rolename"`
	Token    string `json:"Access_Token"`
}
