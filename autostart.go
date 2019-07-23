//+build !windows

package main

import "github.com/ProtonMail/go-autostart"

var app = autostart.App{
	Name:        "wifi-login",
	Exec:        []string{"bash", "-c", getCurrentProgramPath()},
	DisplayName: "wifi auto login service",
}

func addToStartup() {
	if app.IsEnabled() {
		logger.Info("当前已在开机启动列表中")
	} else {
		err := app.Enable()
		if err != nil {
			logger.Error(err)
		}
		logger.Info("已加入开机启动")
	}
}

func removeStartup() {
	if app.IsEnabled() {
		err := app.Disable()
		if err != nil {
			logger.Error(err)
		}
		logger.Info("已取消开机启动")

	} else {
		logger.Info("当前不在开机启动列表中")
	}
}
