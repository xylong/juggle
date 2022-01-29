package lib

import (
	"log"
	"os"
	"os/signal"
)

// ServerSignalChan 信号
var ServerSignalChan chan os.Signal

func init() {
	ServerSignalChan = make(chan os.Signal)
}

// ServerNotify 监听信号 
func ServerNotify() {
	signal.Notify(ServerSignalChan, os.Interrupt)
	<-ServerSignalChan
}

// ShutDown 关闭服务
func ShutDown(err error) {
	log.Printf("database connect: %s\n", err.Error())
	ServerSignalChan <- os.Interrupt
}
