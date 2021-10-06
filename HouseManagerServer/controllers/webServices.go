package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/gormiti"
	"main/models"
	"main/services"
	"main/websockets"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/unrolled/secure"
)

func Cors() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Add("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization,GUItype, Content-Length, Accept-Encoding, X-CSRF-Token, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, DELETE, PATCH")
		//c.Writer.Header().Add("Cache-Control", "max-age=3600")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}

		c.Next()
	}
}

func SecureFunc() gin.HandlerFunc {
	return func(c *gin.Context) {
		var endPoint string
		if len(os.Args) > 1 {
			if os.Args[1] == "prod" {
				endPoint = ":443"
			} else {
				endPoint = ":443"
			}
		}
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect:          true,
			SSLTemporaryRedirect: false,
			SSLHost:              endPoint,
			IsDevelopment:        (gin.Mode() != gin.ReleaseMode),
		})
		fmt.Println()
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		if gin.Mode() == gin.ReleaseMode {
			err := services.TokenValid(c.Request)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				Println(c.Request.RemoteAddr + " " + err.Error())
				c.Abort()
				return
			}
			ad, err := services.ExtractTokenMetadata(c.Request)
			_, err = services.FetchAuth(ad)
			if err != nil {
				c.JSON(http.StatusUnauthorized, err.Error())
				Println(c.Request.RemoteAddr + " " + err.Error())
				c.Abort()
				return
			}

			c.Next()
		} else {
			c.Next()
		}
	}
}

func RoleAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		//if gin.Mode() == gin.ReleaseMode {
		err := services.TokenValid(c.Request)
		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			Println(c.Request.RemoteAddr + " " + err.Error())
			c.Abort()
			return
		}
		ad, err := services.ExtractTokenMetadata(c.Request)
		_, err = services.FetchAuth(ad)

		if err != nil {
			c.JSON(http.StatusUnauthorized, err.Error())
			Println(c.Request.RemoteAddr + " " + err.Error())
			c.Abort()
			return
		}
		auth, _ := services.IsAdministratorUser(ad.UserName)
		if !auth {
			c.JSON(http.StatusUnauthorized, "ACCESS DENIED - Not sufficient privileges for "+ad.UserName)
			Println(c.Request.RemoteAddr + " " + "ACCESS DENIED - Not sufficient privileges for " + ad.UserName)
			c.Abort()
			return
		}

		c.Next()
		/*} else {
			c.Next()
		}*/
	}
}

func Login(c *gin.Context) {
	var user models.User
	errBind := c.Bind(&user)

	if errBind != nil || user.Username == "" || user.Password == "" {
		Println("Login request failed. ErrBind: " + errBind.Error() + ", username: " + user.Username)
		c.JSON(http.StatusBadRequest, "Malformed request")
		return
	}

	ts, err := services.CreateToken(user.Username)
	if err != nil {
		Println("Unable to create Token for user " + user.Username + ", Error: " + err.Error())
		c.JSON(http.StatusUnprocessableEntity, err.Error())
		return
	}
	saveErr := services.CreateAuth(user.Username, ts)
	if saveErr != nil {
		Println("Unable to insert Token for user " + user.Username + ", Error: " + err.Error())
		c.JSON(http.StatusUnprocessableEntity, saveErr.Error())
	}

	user.Token = ts.AccessToken
	userData, _ := services.GetUserPrivileges(user.Username, user.Password)

	if (gormiti.UserGorm{}) == userData {
		Println("Connection failed - User Not in DB. Username: " + user.Username)
		c.JSON(http.StatusUnauthorized, "Connection failed - User Not in DB. Username: "+user.Username)
		return
	}

	user.Password = ""
	/* user.Name = userData.Name
	user.Surname = userData.Surname
	user.Email = userData.Email
	user.Rolename = userData.Role.Rolename */

	if gin.Mode() == gin.ReleaseMode {
		InfoPrintln("Correct login for " + user.Username)
	}

	c.JSON(http.StatusOK, user)
}

func Logout(c *gin.Context) {
	au, err := services.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	delErr := services.DeleteAuth(au.AccessUuid)
	if delErr != nil { //if any goes wrong
		c.JSON(http.StatusUnauthorized, "unauthorized")
		return
	}
	Println(au.UserName + " - Successfully logged out")
	c.JSON(http.StatusOK, "Successfully logged out")
}

func GetManagementData(c *gin.Context) {

	var list models.UserManagement
	userList, err := services.GetUserList()
	if err != nil {

	}

	roleList, err := services.GetRoleList()
	if err != nil {

	}

	list.Users = userList
	list.Roles = roleList

	c.JSON(http.StatusOK, list)
}

func PostUser(c *gin.Context) {

	var userRaw map[string]interface{}
	var user gormiti.UserGorm

	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &userRaw)

	data, _ := json.Marshal(userRaw)
	json.Unmarshal(data, &user)

	/* layout1 := "2006-01-02T15:04:05.000"
	t1, err1 := time.Parse(layout1, userRaw["ValidityFrom"].(string))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1.Error())
		Println("Post user Error (Parse Time ValidityFrom) " + err1.Error())
		return
	} */

	hashedPass, errPass := services.HashPassword(userRaw["Password"].(string))
	if errPass != nil {
		c.JSON(http.StatusBadRequest, errPass.Error())
		Println("Post user Error (Hash Password) " + errPass.Error())
		return
	}
	user.Password = hashedPass

	_, err := services.PostUser(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		Println("Post user Error " + err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func DeleteUser(c *gin.Context) {

	idUser := c.Query("idUser")
	err := services.DeleteUser(idUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		Println("Delete user Error " + err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateUser(c *gin.Context) {

	var userRaw map[string]interface{}
	var user gormiti.UserGorm

	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &userRaw)

	data, _ := json.Marshal(userRaw)
	json.Unmarshal(data, &user)

	/* layout1 := "2006-01-02T15:04:05.000"
	t1, err1 := time.Parse(layout1, userRaw["ValidityFrom"].(string))
	if err1 != nil {
		c.JSON(http.StatusBadRequest, err1.Error())
		Println("Update user Error (Parse Time ValidityFrom) " + err1.Error())
		return
	} */

	/* t2, err2 := time.Parse(layout1, userRaw["ValidityTo"].(string))
	if err2 != nil {
		c.JSON(http.StatusBadRequest, err1.Error())
		Println("Update user Error (Parse Time ValidityTo) " + err1.Error())
		return
	}
	user.ValidityFrom = t1
	user.ValidityTo = t2 */

	if userRaw["Password"] != "" {
		hashedPass, errPass := services.HashPassword(userRaw["Password"].(string))
		if errPass != nil {
			c.JSON(http.StatusBadRequest, errPass.Error())
			Println("Post user Error (Hash Password) " + errPass.Error())
			return
		}
		user.Password = hashedPass
	}

	err := services.UpdateUser(user)

	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		Println("Update user Error " + err.Error())
		return
	}

	c.JSON(http.StatusOK, "")
}

func UpdateAccountInfo(c *gin.Context) {

	var userRaw map[string]interface{}
	var user gormiti.UserGorm

	body, _ := ioutil.ReadAll(c.Request.Body)
	json.Unmarshal(body, &userRaw)

	data, _ := json.Marshal(userRaw)
	json.Unmarshal(data, &user)

	ad, tokenErr := services.ExtractTokenMetadata(c.Request)
	if tokenErr != nil {
		c.JSON(http.StatusUnauthorized, tokenErr.Error())
		Println(c.Request.RemoteAddr + " " + tokenErr.Error())
		return
	}
	if ad.UserName != user.Username {
		c.JSON(http.StatusUnavailableForLegalReasons, "Can't update "+user.Username+" because you're "+ad.UserName)
		Println("Can't update " + user.Username + " because you're " + ad.UserName)
		return
	}

	if userRaw["Password"] != "" {
		hashedPass, errPass := services.HashPassword(userRaw["Password"].(string))
		if errPass != nil {
			c.JSON(http.StatusBadRequest, errPass.Error())
			Println("Post user Error (Hash Password) " + errPass.Error())
			return
		}
		user.Password = hashedPass
	}

	err := services.UpdateAccountInfo(user)
	if err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		Println("Update user Error " + err.Error())
		return
	}

	c.JSON(http.StatusOK, "User Information Updated")
}

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func ServiceCheckStart() {
	var timeControl string
	for {
		var dataFrame []models.AlleantiaProbe

		/*if services.AlleantiaDB == nil {
			Println("Alleantia database's connection not initialized")
			break
		}*/

		var temp []map[string]interface{}
		start := time.Now()

		//services.AlleantiaDB.Raw("SELECT LAST_VALUE(QUALITY) OVER (PARTITION BY VARIABLE_ID ORDER BY LOG_TIMESTAMP_UTC) FROM dbo.LOG_DATA_A where VARIABLE_ID IN (1,2,13,14,15,16)").Scan(&temp)
		//services.AlleantiaDB.Raw("SELECT TOP (100) * FROM dbo.LOG_DATA_A where VARIABLE_ID IN (1,2,13,14,15,16)").Scan(&temp)
		duration := time.Since(start)
		fmt.Println(duration)
		temp1, err1 := json.Marshal(temp)
		if err1 != nil {
			fmt.Println("Failed to Marshal in websocket: ", err1)
			return
		}
		json.Unmarshal(temp1, &dataFrame)

		if timeControl != "MOCK UNCOMMENT" { //string(dataFrame[99].LogTimestampUtc.Format("15:04")) {
			timeControl = string(dataFrame[99].LogTimestampUtc.Format("15:04"))

			for client, _ := range websockets.Pools.Clients {
				fmt.Println(timeControl)
				if err := client.Conn.WriteJSON(dataFrame); err != nil {
					fmt.Println(err)
					return
				}
			}
			/* _, _, err1 := conn.ReadMessage()
			if err1 != nil {
				break
			} */
		}
		/* var result []models.AlleantiaProbe
		var temp []map[string]interface{}

		AlleantiaDB.Raw("SELECT TOP (1000) * FROM dbo.LOG_DATA_A").Scan(&temp)
		b, err := json.Marshal(temp)
		json.Unmarshal(b, &result)
		bs, err := json.Marshal(result)
		fmt.Println(string(bs)) */

		time.Sleep(10 * time.Second)

	}
}

func ServeWs(pool *websockets.Pool, w http.ResponseWriter, r *http.Request) {
	fmt.Println("WebSocket Endpoint Hit")
	conn, err := websockets.Upgrade(w, r)
	if err != nil {
		fmt.Fprintf(w, "%+v\n", err)
	}

	client := &websockets.Client{
		Conn: conn,
		Pool: pool,
	}

	pool.Register <- client
	client.Read()
}
