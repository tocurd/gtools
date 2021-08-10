/*
 * @Author: your name
 * @Date: 2021-07-30 21:32:46
 * @LastEditTime: 2021-07-30 21:34:10
 * @LastEditors: Please set LastEditors
 * @Description: In User Settings Edit
 * @FilePath: \gtools\gtools\ip.go
 */
package gtools

import (
	"bytes"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
)

var Ip ipInterface

type ipInterface interface {
	GetNetworkIP() (exip string)
	GetLocalIP() (ip string)
	GetClientIP(r *http.Request) (ip string)
}

type ip struct{}

func init() {
	Ip = &ip{}
}

// GetWwwIP 获取公网IP地址
func (_ip ip) GetNetworkIP() (exip string) {
	resp, err := http.Get("http://myexternalip.com/raw")
	if err != nil {
		return ""
	}
	defer resp.Body.Close()
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return ""
	}
	return string(bytes.TrimSpace(b))
}

// GetLocalIP 获取内网ip
func (_ip ip) GetLocalIP() (ip string) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return
	}
	for _, a := range addrs {
		if ipnet, ok := a.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip = ipnet.IP.String()
				break
			}
		}
	}
	return
}

// GetClientIP 获取用户ip
func (_ip ip) GetClientIP(r *http.Request) (ip string) {
	ip = r.Header.Get("X-Real-Ip")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return
}
