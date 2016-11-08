package main

import "fmt"

func getAllIp(ipStr string) {
	length := len(source)
	if length > 12 || length < 4 {
		return
	}
	var ips []string
	if length == 4 {
		ip := string(ipStr[0]) + "." + string(ipStr[1]) + "." + string(ipStr[2]) + "." + string(ipStr[3])
		ips = append(ips, ip)
	} else {

	}
}

func main() {
	var ipStr string
	fmt.Scan(&ipStr)
	getAllIp(ipStr)
}
