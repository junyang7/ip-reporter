package main

import (
	"fmt"
	"ip-reporter/_datetime"
	"ip-reporter/_email"
	"net"
)

var (
	HOST     = ""
	PORT     = 0
	USERNAME = ""
	PASSWORD = ""
	FROM     = ""
	TO       = []string{}
)

func main() {
	ip := getLocalIp()
	datetime := _datetime.Get()
	subject := "ip-reporter"
	content := datetime + "\t" + ip
	fmt.Println(content)
	m := _email.New(HOST, PORT, USERNAME, PASSWORD, FROM, TO, subject, content)
	m.Send()
	fmt.Println("success")
}
func getLocalIp() string {
	addrList, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, addr := range addrList {
		ipNet, _ := addr.(*net.IPNet)
		if ipNet.IP.To4() != nil && !ipNet.IP.IsLoopback() {
			return ipNet.IP.String()
		}
	}
	return "127.0.0.1"
}
