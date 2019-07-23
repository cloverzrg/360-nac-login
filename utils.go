package main

import (
	"flag"
	"github.com/jinzhu/configor"
	"os"
	"os/signal"
	"regexp"
)

type cmdParams struct {
	ConfigPath string
	Install    bool
	Uninstall  bool
	Daemon     bool
	LogPath    string
}

func parseCmdParams() (params *cmdParams) {
	params = &cmdParams{}
	flag.StringVar(&params.ConfigPath, "c", "config.json", "配置文件")
	flag.StringVar(&params.LogPath, "log", "", "日志保存位置")
	flag.BoolVar(&params.Install, "i", false, "Install 打开开机启动")
	flag.BoolVar(&params.Uninstall, "u", false, "Uninstall 关闭开机启动")
	flag.BoolVar(&params.Daemon, "d", false, "daemon 后台运行")
	flag.Parse()
	return params
}

var Config = struct {
	Username   string   `json:"username"`
	Password   string   `json:"password"`
	CronSpec   []string `json:"cron_spec"`
	WifiRegexp []string `json:"wifi_regexp"`
}{}

func LoadConfig(path string) error {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsNotExist(err) {
			return err
		}
	}
	err = configor.Load(&Config, path)
	if err != nil {
		logger.Panic(err)
		return err
	}
	logger.Infof("config: %+v", Config)
	return err
}

func waitStopSignal() {
	signals := make(chan os.Signal, 1)
	signal.Notify(signals, os.Interrupt)
Loop:
	select {
	case <-signals:
		break Loop
	}
}

func checkWifiName(wifiName string) bool {
	//wifiName := wifiname.WifiName()
	for _, v := range Config.WifiRegexp {
		reg := regexp.MustCompile(v)
		if reg.MatchString(wifiName) {
			return true
		}
	}

	return false
}

func getCurrentProgramPath() string {
	ex, err := os.Executable()
	if err != nil {
		logger.Panic(err)
	}
	return ex
}
