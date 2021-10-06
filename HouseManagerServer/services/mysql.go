package services

import (
	"database/sql"
	"errors"
	"io"
	"log"
	"main/gormiti"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var GlobalDB *gorm.DB = nil

func OpenGormConnection(logGorm *rotatelogs.RotateLogs) error {

	sqlConnection := "root:admin@tcp(127.0.0.1:3306)/HouseManager?charset=utf8mb4&parseTime=True"

	conn, err := sql.Open("mysql", sqlConnection)
	if err != nil {
		return err
	}

	internalLogger := log.New(os.Stdout, "\r\n", log.LstdFlags)
	internalLogger.SetOutput(io.MultiWriter(os.Stdout, logGorm))
	newLogger := logger.New(
		internalLogger, // io writer
		logger.Config{
			SlowThreshold: time.Second,   // Slow SQL threshold
			LogLevel:      logger.Silent, // Log level
			Colorful:      false,         // Disable color
		},
	)
	db, err := gorm.Open(mysql.New(mysql.Config{Conn: conn}), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}

	/*var count int64

	var roles []gormiti.RoleGorm
	roles = append(roles, gormiti.RoleGorm{RoleId: 1, Rolename: "Administrator"})
	for _, s := range roles {
		_ = db.Where(&s).Find(&s).Count(&count)
		if count == 0 {
			db.Create(&s)
		}
	}*/
	err2 := db.SetupJoinTable(&gormiti.InquilinoGorm{}, "Casa", &gormiti.LocazioneGorm{})
	if err2 != nil {
		return err
	}
	err = db.AutoMigrate(&gormiti.CasaGorm{}, &gormiti.InquilinoGorm{}, &gormiti.UserGorm{}, &gormiti.BollettaGorm{})
	if err != nil {
		return err
	}

	GlobalDB = db
	return nil
}

func CreateAuth(username string, td *TokenDetails) error {
	at := time.Unix(td.AtExpires, 0) //converting Unix to UTC(to Time object)
	if GlobalDB == nil {
		return errors.New("database connection not initialized")
	}

	token := gormiti.AccessTokenGorm{Uuid: td.AccessUuid, Username: username, ExpiresAt: at}

	result := GlobalDB.Create(&token)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func FetchAuth(authD *AccessDetails) (string, error) {

	if GlobalDB == nil {
		return "", errors.New("database connection not initialized")
	}

	var token gormiti.AccessTokenGorm
	token.Uuid = authD.AccessUuid
	result := GlobalDB.Take(&token)
	if result.Error != nil {
		return "", result.Error
	}

	return token.Username, nil
}

func DeleteAuth(givenUuid string) error {

	if GlobalDB == nil {
		return errors.New("database connection not initialized")
	}

	result := GlobalDB.Delete(gormiti.AccessTokenGorm{}, "uuid = ?", givenUuid)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func IsAdministratorUser(username string) (bool, error) {

	if GlobalDB == nil {
		return false, errors.New("database connection not initialized")
	}

	var users []gormiti.UserGorm

	var userG gormiti.UserGorm
	userG.Username = username

	GlobalDB.Model(userG).Where(userG).Find(&users)

	/* for _, u := range users {
		if u.IdRole == 1 {
			return true, nil
		}
	} */

	return false, nil
}

func ExternalUserLogin(username string, password string) (bool, error) {

	if GlobalDB == nil {
		return false, errors.New("database connection not initialized")
	}
	var user gormiti.UserGorm
	var userG gormiti.UserGorm
	userG.Username = username

	GlobalDB.Model(userG).Where(userG).Find(&user)

	//asd, _ := HashPassword(password)
	//fmt.Println(asd)
	//fmt.Println(user.Password)

	auth := CheckPasswordHash(password, user.Password)

	return auth, nil
}

func GetUserPrivileges(username string, password string) (gormiti.UserGorm, error) {

	if GlobalDB == nil {
		return gormiti.UserGorm{}, errors.New("database connection not initialized")
	}

	var user gormiti.UserGorm
	var userG gormiti.UserGorm
	userG.Username = username

	GlobalDB.Model(userG).Where(userG).Joins("Role").Joins("Team").Find(&user)

	return user, nil
}

func GetUserList() ([]gormiti.UserGorm, error) {

	if GlobalDB == nil {
		return []gormiti.UserGorm{}, errors.New("database connection not initialized")
	}

	var user []gormiti.UserGorm
	var userG gormiti.UserGorm

	GlobalDB.Model(userG).Joins("Role").Joins("Team").Order("id_role").Order("username").Find(&user)

	return user, nil
}

func GetRoleList() ([]gormiti.RoleGorm, error) {

	if GlobalDB == nil {
		return []gormiti.RoleGorm{}, errors.New("database connection not initialized")
	}

	var role []gormiti.RoleGorm
	var roleG gormiti.RoleGorm

	GlobalDB.Model(roleG).Find(&role)

	return role, nil
}

func PostUser(user gormiti.UserGorm) (gormiti.UserGorm, error) {

	if GlobalDB == nil {
		return gormiti.UserGorm{}, errors.New("database connection not initialized")
	}
	res := GlobalDB.Create(&user)

	var userG gormiti.UserGorm
	GlobalDB.Model(userG).Find(&user)
	return user, res.Error
}

func DeleteUser(userId string) error {

	if GlobalDB == nil {
		return errors.New("database connection not initialized")
	}
	res := GlobalDB.Delete(&gormiti.UserGorm{}, userId)

	return res.Error
}

func UpdateUser(user gormiti.UserGorm) error {

	if GlobalDB == nil {
		return errors.New("database connection not initialized")
	}
	var res *gorm.DB
	/* if user.IdTeam == 0 {
		res = GlobalDB.Model(user).Select("id_team").Updates(map[string]interface{}{"id_team": nil})
		if res.Error != nil {
			return res.Error
		}
	} */

	res = GlobalDB.Omit("Role").Omit("Team").Updates(&user)

	return res.Error
}

func UpdateAccountInfo(user gormiti.UserGorm) error {

	if GlobalDB == nil {
		return errors.New("database connection not initialized")
	}

	res := GlobalDB.Model(user).Omit("username").Where("username = ?", user.Username).Updates(&user)

	return res.Error
}
