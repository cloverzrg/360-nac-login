package main

import (
	"github.com/robfig/cron/v3"
	"github.com/sirupsen/logrus"
	wifiname "github.com/yelinaung/wifi-name"
	"os"
	"os/exec"
)

var logger = logrus.New()

func main() {
	c := cron.New(cron.WithSeconds())
	for _, v := range Config.CronSpec {
		logger.Info("cron:", v)
		_, err := c.AddFunc(v, login)
		if err != nil {
			logger.Error(err)
		}
	}
	c.Start()

	waitStopSignal()
}

func init() {
	logger.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
		ForceColors:   true,
	})

	//logger.SetReportCaller(true)
	params := parseCmdParams()
	if params.Install {
		addToStartup()
		os.Exit(0)
	}
	if params.Uninstall {
		removeStartup()
		os.Exit(0)
	}
	if params.LogPath != "" {
		logger.Info("log file:", params.LogPath)
		f, err := os.OpenFile(params.LogPath, os.O_WRONLY|os.O_CREATE, 0755)
		if err != nil {
			logger.Error(err)
		}
		logger.SetOutput(f)
	}
	err := LoadConfig(params.ConfigPath)
	if err != nil {
		logger.Error(err)
		os.Exit(1)
	}

	if len(Config.CronSpec) == 0 {
		logger.Panic("cron_spec 未填写")
		os.Exit(1)
	}

	logger.Info("process path:", getCurrentProgramPath())
	logger.Info("current wifi name:", wifiname.WifiName())
	logger.Info("wifi name check passed:", checkWifiName(wifiname.WifiName()))

	if params.Daemon {
		args := os.Args[1:]

		for i := 0; i < len(args); i++ {
			if args[i] == "-d=true" {
				args[i] = "-d=false"
				break
			}
			if args[i] == "-d" {
				args[i] = "-d=false"
				break
			}
		}
		cmd := exec.Command(os.Args[0], args...)
		err := cmd.Start()
		if err != nil {
			logger.Error(err)
		}
		logger.Info("[PID]", cmd.Process.Pid)
		os.Exit(0)
	}

}
