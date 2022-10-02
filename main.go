package main

import (
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"
)

func main() {
	WHHYconn()
}

// 网卡名称 如:wlan0
// 返回ip地址
func GetIP(NetName string) string {
	addrs, _ := net.InterfaceByName(NetName)
	addr, _ := addrs.Addrs()
	return strings.Split(addr[0].String(), "/")[0]
}

// 获取时间戳
func NowTime() string {
	now := time.Now() //获取当前时间
	return strconv.Itoa(int((now.UnixNano()/1e6)/1000) & 65535)
}
func WHHYconn() {
	TimePort := NowTime()
	ip := GetIP("wlan0")
	urlValue := url.Values{}
	urlValue.Add("action", "req_auth")
	urlValue.Add("user", "19308430661")
	urlValue.Add("pass", "017504")
	urlValue.Add("vlan", "22.0")
	urlValue.Add("ope", "单宽")
	urlValue.Add("ip", ip)
	urlValue.Add("serial", TimePort)
	urlValue.Add("bras", "Panabit")
	http.PostForm("http://172.16.10.9/ajax_cmcc.php", urlValue)
}
