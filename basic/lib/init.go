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

func ServerNotify() {
	signal.Notify(ServerSignalChan, os.Interrupt)
	<-ServerSignalChan
}

func ShutDown(err error) {
	ServerSignalChan <- os.Interrupt
	log.Printf("database connect: %s", err.Error())
}
