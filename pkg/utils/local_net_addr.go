package utils

import (
	"net"
)

func GetLocalAddr() string {
	var localArr string

	addrList, err := net.InterfaceAddrs()
	if err != nil {
		panic(err)
	}
	for _, address := range addrList {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				localArr = ipNet.IP.String()
			}
		}
	}
	return localArr
}
