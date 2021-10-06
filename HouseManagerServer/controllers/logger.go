package controllers

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
)

var Logger *log.Logger
var infoLogger *log.Logger

func GetLogFolder() string {

	logFolder := os.Getenv("LOCALAPPDATA") + "\\IDEAServer\\logs\\"
	_ = ioutil.WriteFile(".\\aaaaa\\dat1.txt", []byte(logFolder), 0700)
	return logFolder
}

func InitLogger() {

	Logf, err := rotatelogs.New(
		GetLogFolder()+"\\Idea.log.%Y-%m-%d",
		rotatelogs.WithLinkName(GetLogFolder()+"\\Idea.log"),
		rotatelogs.WithMaxAge(30*24*time.Hour),
		rotatelogs.WithRotationTime(24*time.Hour),
	)

	if err != nil {
		log.Fatal(err)
	}

	infoLogger = log.New(Logf, "INFO: ", log.Ldate|log.Ltime)
	Logger = log.New(Logf, "ERROR: ", log.Ldate|log.Ltime)

	infoLogger.Println("File initialization")

}

func InfoPrintln(message string) {
	fmt.Println(message)
	infoLogger.Println(message)
}

func Println(message string) {
	fmt.Println(message)
	Logger.Println(message)
}
