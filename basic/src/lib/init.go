package lib

import (
	"juggle/basic/src/config"
	"log"
	"os"
	"os/signal"
)

var (
	// ServerSignalChan 信号
	ServerSignalChan chan os.Signal

	// 配置
	Config *config.Config
)

func init() {
	ServerSignalChan = make(chan os.Signal)
	Config = &config.Config{}
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
