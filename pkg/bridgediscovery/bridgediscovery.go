package bridgediscovery

import (
	"fmt"
	"log"
	"net"
	"strings"
)

const (
	maxDatagramSize        = 8192
	multiCastAddressString = "239.255.255.250:1900"
	discoveryMessage       = "M-SEARCH * HTTP/1.1\r\n" +
		"HOST:239.255.255.250:1900\r\n" +
		"ST:ssdp:all\r\n" +
		"Man:\"ssdp:discover\"\r\n" +
		"MX:3\r\n\r\n"
)

// Bridge holds information about found bridges in the network
type Bridge struct {
	ID       string
	Location string
	USN      string
}

var bridges map[string]Bridge

// StartDiscovery initiates discovery of Hue bridges
func StartDiscovery() {
	bridges = make(map[string]Bridge)
	// go listenForBridgeAdvertisements(multiCastAddressString)
	go discoverBridges()
}

func discoverBridges() {
	con := openConnection()
	defer con.Close()

	sendMultiCastMessage(*con)

	listenForReplies(con)
}

func openConnection() *net.UDPConn {
	// Make sure we send out the broadcast from the right interface
	localIP := getLocalIP() + ":0"
	listenAddr, err := net.ResolveUDPAddr("udp", localIP)

	if err != nil {
		log.Panicf("Could not resolve listen address address %v, %v", localIP, err.Error())
	}

	con, err := net.ListenUDP("udp", listenAddr)

	if err != nil {
		log.Panicf("Unable to set up UDP connection to broadcast address, %v", err.Error())
	}

	con.SetReadBuffer(maxDatagramSize)

	return con
}

func sendMultiCastMessage(con net.UDPConn) {
	log.Println("Sending on : " + con.LocalAddr().(*net.UDPAddr).String())

	mcastAddr, err := net.ResolveUDPAddr("udp", multiCastAddressString)

	if err != nil {
		log.Panicf("Could not resolve multicast address %v, %v", multiCastAddressString, err.Error())
	}

	_, err2 := con.WriteToUDP([]byte(discoveryMessage), mcastAddr)

	if err2 != nil {
		log.Panicf("Unable to send discovery message, %v", err2.Error())
	}
}

func listenForReplies(con *net.UDPConn) {

	for {
		b := make([]byte, maxDatagramSize)

		bytesRead, _, err := con.ReadFromUDP(b)
		if err != nil {
			log.Fatal("ReadFromUDP failed:", err)
		}

		// Convert to string and ignore part of the slice that contains no data
		s := string(b[0 : bytesRead-1])

		split := strings.Split(s, "\r\n")

		bridge := new(Bridge)

		for _, field := range split {
			parseField(bridge, field)
		}

		if bridge.ID != "" {
			log.Printf("Bridge found, %v", bridge.Location)
			bridges[bridge.ID] = *bridge
		}
	}
}

func parseField(bridge *Bridge, field string) {
	if strings.HasPrefix(field, "hue-bridgeid:") {
		bridge.ID = readAttribute(field, "hue-bridgeid")
	} else if strings.HasPrefix(field, "LOCATION:") {
		bridge.Location = readAttribute(field, "LOCATION")
	} else if strings.HasPrefix(field, "USN:") {
		bridge.USN = readAttribute(field, "USN")
	}
}

func listenForBridgeAdvertisements(address string) {
	listenIf := getListenInterface()
	fmt.Println(listenIf.Name)

	addr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Panicf("Error resolving address %v, %v", address, err.Error())
	}

	l, err := net.ListenMulticastUDP("udp", &listenIf, addr)
	if err != nil {
		log.Panicf("Error listening for multicast, %v", err.Error())
	}

	log.Printf("Listening on local address %v for broadcasts", l.LocalAddr())

	listenForReplies(l)
}

func readAttribute(fromString, attribute string) string {
	result := strings.Replace(fromString, attribute+":", "", -1)
	result = strings.Trim(result, " ")
	return result
}

func getLocalIP() string {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		log.Fatal(err)
	}

	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}

func getListenInterface() net.Interface {
	ifs, err := net.Interfaces()

	if err != nil {
		log.Fatal(err)
	}

	for _, i := range ifs {
		log.Println(i)
		if (i.Flags&net.FlagLoopback == 0) && (i.Flags&net.FlagUp == 1) {
			return i
		}
	}

	return ifs[0]
}
