package main

import (
	"encoding/binary"
	"log"
	"main/controllers"
	"main/mappings"
	"math"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	logFolder := controllers.GetLogFolder()
	err := os.MkdirAll(logFolder, os.ModePerm)
	if err != nil {
		log.Fatal("Unable to create log folder")
		os.Exit(255)
	}

	os.Setenv("ACCESS_SECRET", "bLX,&;kRMPU_CF6N3wr\bkn{HT')7%ukDp-k5(AwBP=9Ha)=m(~c@3=Hw;be") //this should be in an env file

	mappings.CreateUrlMappings()

	controllers.InitLogger()

	mappings.Router.Run(":3300")
	controllers.InfoPrintln("Listening HTTP on 3300")
	if gin.Mode() == gin.ReleaseMode {
		go mappings.Router.Run(":80")
		controllers.InfoPrintln("Listening HTTP on 80")
	}
	if err == nil {

	} else {
		println("Could not Start HTTPS Server, missing Certificate")
	}

	//client := modbus.TCPClient("localhost:502")

}

func Float64frombytes(bytes []byte) float64 {
	bits := binary.LittleEndian.Uint64(bytes)
	float := math.Float64frombits(bits)
	return float
}
