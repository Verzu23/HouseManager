package mappings

import (
	"io"
	"log"
	"main/controllers"
	"main/services"
	"main/websockets"
	"net/http"
	"os"
	"time"

	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	//"golang.org/x/crypto/acme/autocert"
)

var Router *gin.Engine

func CreateUrlMappings() {

	//gin.SetMode(gin.ReleaseMode)

	gin.ForceConsoleColor()

	websockets.Pools = websockets.NewPool()
	go websockets.Pools.Start()

	//f, _ := os.Create("gin.log")
	rl, err := rotatelogs.New(
		controllers.GetLogFolder()+"\\gin.log.%Y-%m-%d",
		rotatelogs.WithLinkName(controllers.GetLogFolder()+"\\gin.log"),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)
	gormrl, err := rotatelogs.New(
		controllers.GetLogFolder()+"\\gorm.log.%Y-%m-%d",
		rotatelogs.WithLinkName(controllers.GetLogFolder()+"\\gorm.log"),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		log.Printf("failed to create rotatelogs: %s", err)
		return
	}
	gin.DefaultWriter = io.MultiWriter(rl, os.Stdout)

	err = services.OpenGormConnection(gormrl)
	if err != nil {
		//TODO print log
		os.Exit(1)
	}

	Router = gin.Default()

	Router.Use(controllers.Cors())
	Router.Use(gzip.Gzip(gzip.DefaultCompression))

	v1 := Router.Group(("/v1"))

	v1.GET("/ws", func(c *gin.Context) {
		controllers.ServeWs(websockets.Pools, c.Writer, c.Request)
		//controllers.Wshandler(c.Writer, c.Request)
	})

	v1.GET("/client", func(c *gin.Context) {
		var s []map[string]string
		for i, _ := range websockets.Pools.Clients {

			s = append(s, map[string]string{"address": i.Conn.RemoteAddr().String(), "ID": i.ID})
		}
		if len(s) == 0 {
			c.JSON(http.StatusOK, "No client connected")
		} else {
			c.JSON(http.StatusAccepted, s)
		}
		//controllers.Wshandler(c.Writer, c.Request)
	})
	/*usrManagement := Router.Group(("/usrManagement"))

	v1.POST("/login", controllers.Login)
	v1.POST("/logout", controllers.TokenAuthMiddleware(), controllers.Logout)
	v1.GET("/GetWtReference", controllers.TokenAuthMiddleware(), controllers.GetWtReference)
	v1.POST("/ChangeFlapConfig", controllers.TokenAuthMiddleware(), controllers.ChangeFlapConfig)
	v1.POST("/ChangeField", controllers.TokenAuthMiddleware(), controllers.ChangeField)
	v1.POST("/ChangeValue", controllers.TokenAuthMiddleware(), controllers.ChangeValue)
	v1.POST("/CalcBalance", controllers.TokenAuthMiddleware(), controllers.GetBalance)
	usrManagement.GET("/usrList", controllers.RoleAuthMiddleware(), controllers.GetManagementData)
	usrManagement.POST("/postUser", controllers.RoleAuthMiddleware(), controllers.PostUser)
	usrManagement.DELETE("/deleteUser", controllers.RoleAuthMiddleware(), controllers.DeleteUser)
	usrManagement.PATCH("/updateUser", controllers.RoleAuthMiddleware(), controllers.UpdateUser)
	usrManagement.PATCH("/updateAccountInfo", controllers.TokenAuthMiddleware(), controllers.UpdateAccountInfo)
	Router.Use(static.Serve("/", static.LocalFile("./public", false))) */

}
