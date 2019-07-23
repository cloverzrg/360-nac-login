# 360-nac-login
360 终端准入控制系统（NAC）用户认证自动登录
开发者使用macos系统进行开发，不确保其他系统是否正常

```
# ./wifi-login -h
Usage of ./wifi-login:
  -c string
    	配置文件 (default "config.json")
  -d	daemon 后台运行
  -i	Install 打开开机启动
  -log string
    	日志保存位置
  -u	Uninstall 关闭开机启动
```


```
{
  "username": "",  // 用户名
  "password": "",  // 密码
  "cron_spec": [    // 自动登录定时任务描述
    "*/15 * * * * ?", 
    "0 0 10 * * MON-FRI"
  ],
  "wifi_regexp": [   // wifi名称正则，只有通过的才会自动认证，可以只写wifi名称的部分字符串
    "360\\d",
    "360Internal\\d"
  ]
}

```
