package networking

import (
	"log"
	"net"
	"strings"
)

// GetLocalIPs retrieves a list of ip addresses that can bound to
func GetLocalIPs() []string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	addresses := make([]string, 0)

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				ip := ipnet.IP.String()

				// not interested in ip's in this range
				if !strings.HasPrefix(ip, "169.254") {
					addresses = append(addresses, ip)
				}
			}
		}
	}

	return addresses
}

// GetListenInterface gets a list of interfaces that can be bound to
func GetListenInterface() []net.Interface {
	res := make([]net.Interface, 0)
	ifs, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range ifs {
		if (i.Flags&net.FlagLoopback == 0) && (i.Flags&net.FlagUp == 1) {
			log.Println(i)
			res = append(res, i)
		}
	}

	return res
}
