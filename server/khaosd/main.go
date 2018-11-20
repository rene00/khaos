package main

import (
	"fmt"
	"log"
	"syscall"

	"github.com/fvbock/endless"

	"github.com/rene00/khaos/models"
	"github.com/rene00/khaos/pkg/setting"
	"github.com/rene00/khaos/routers"
)

func main() {
	setting.Setup()
	models.Setup()

	routersInit := routers.InitRouter()
	readTimeout := setting.ServerSetting.ReadTimeout
	writeTimeout := setting.ServerSetting.WriteTimeout
	endPoint := fmt.Sprintf(":%d", setting.ServerSetting.HttpPort)
	maxHeaderBytes := 1 << 20

	endless.DefaultReadTimeOut = readTimeout
	endless.DefaultWriteTimeOut = writeTimeout
	endless.DefaultMaxHeaderBytes = maxHeaderBytes
	server := endless.NewServer(endPoint, routersInit)
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}

}
