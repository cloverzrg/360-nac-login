package main

import (
	"encoding/base64"
	wifiname "github.com/yelinaung/wifi-name"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strings"
	"time"
)

func detect() (needLogin bool, err error) {
	needLogin = false
	client := http.Client{}
	client.Timeout = time.Second * 2
	resp, err := client.Get("http://nossl.jeongen.com/")
	if err != nil {
		urlErr := err.(*url.Error)
		if urlErr.Temporary() {
			return true, nil
		}
		return needLogin, err
	}

	if resp.StatusCode == 302 {
		needLogin = true
	}
	return needLogin, err
}

func login() {
	if !checkWifiName(wifiname.WifiName()) {
		logger.Info("wifi 名称检查不通过，不需要自动登录")
		return
	}
	needLogin, err := detect()
	if err != nil {
		logger.Error(err)
		return
	}
	if !needLogin {
		logger.Info("当前为登陆状态")
		return
	}
	logger.Info("start login")
	username := Config.Username
	password := Config.Password
	passwordBase64 := base64.StdEncoding.EncodeToString([]byte(password))
	resp, err := http.Get("http://172.16.0.92/portal/login")
	if err != nil {
		logger.Error(err)
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	reg := regexp.MustCompile(`{"csrf_token": "(.+)"}`)
	result := reg.FindSubmatch(body)
	csrfToken := string(result[1])
	logger.Info("csrf_token:", csrfToken)
	client := http.Client{Timeout: 1 * time.Second}
	data := url.Values{}
	data.Set("user_name", username)
	data.Set("pwd", passwordBase64)
	data.Set("request_url", "http://163.com/")
	data.Set("os_name", "Mac OS")
	data.Set("browser_name", "chrome 71.0.3578.98")
	data.Set("force_change", "0")
	req, err := http.NewRequest("POST", "http://172.16.0.92/api/portal/login", strings.NewReader(data.Encode()))
	if err != nil {
		logger.Error(err)
		return
	}
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("X-CSRFToken", csrfToken)
	cookie := http.Cookie{Name: resp.Cookies()[0].Name, Value: resp.Cookies()[0].Value}
	req.AddCookie(&cookie)
	resp, err = client.Do(req)
	if err != nil {
		logger.Error(err)
		return
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error(err)
		return
	}
	logger.Info(string(body))
	logger.Info("login finish")
	pushNotify("已自动登录wifi网络")
}
