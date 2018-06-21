package bridgediscovery

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"github.com/joostaarts/GolangHue/models"
	"github.com/joostaarts/GolangHue/pkg/networking"
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

var connections *networking.ConnectionContainer
var bridges map[string]models.Bridge
var mutex sync.Mutex

func dispose() {
	connections.Dispose()
}

// StartDiscovery initiates discovery of Hue bridges
func StartDiscovery() {
	connections = new(networking.ConnectionContainer)
	bridges = make(map[string]models.Bridge)

	go listenForBridgeAdvertisements(multiCastAddressString)
	go discoverBridges()
}

func resolveBridge(bridgeInfo *BridgeInfo) {
	resp, err := http.Get(bridgeInfo.Location)

	if err != nil {
		log.Println(err)
		return
	}

	body, err := ioutil.ReadAll(resp.Body)
	var b models.Root

	err2 := xml.Unmarshal(body, &b)
	// log.Print(string(body))

	if err2 != nil {
		log.Fatal(err2)
	}

	mutex.Lock()
	defer mutex.Unlock()
	bridges[bridgeInfo.ID] = b.Device
}

func discoverBridges() {
	// Make sure we send out the broadcast from the right interface
	localIPs := networking.GetLocalIPs()

	for _, ip := range localIPs {
		con := openConnection(ip)

		connections.AddConnection(*con)

		sendMultiCastMessage(*con)

		go listenForReplies(con)
	}
}

func openConnection(localIP string) *net.UDPConn {
	listenAddr, err := net.ResolveUDPAddr("udp", localIP+":0")

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

func listenForBridgeAdvertisements(address string) {
	listenIfs := networking.GetListenInterface()

	for _, listenIf := range listenIfs {
		addr, err := net.ResolveUDPAddr("udp", address)
		if err != nil {
			log.Panicf("Error resolving address %v, %v", address, err.Error())
		}

		l, err := net.ListenMulticastUDP("udp", &listenIf, addr)
		if err != nil {
			log.Panicf("Error listening for multicast, %v", err.Error())
		}

		connections.AddConnection(*l)

		log.Printf("Listening on local address %v for broadcasts", l.LocalAddr())

		go listenForReplies(l)
	}
}

func listenForReplies(con *net.UDPConn) {
	con.SetDeadline(time.Now().Add(10 * time.Second))

	for {
		b := make([]byte, maxDatagramSize)

		n, err := con.Read(b)

		if err != nil {
			if operror, ok := err.(*net.OpError); ok {
				if e, ok := operror.Err.(net.Error); !ok || e.Timeout() {
					connections.CloseConnection(con)
					return
				}
			}

			log.Fatal("ReadFromUDP failed:", err)
		}

		// Convert to string and ignore part of the slice that contains no data
		bytesRead := b[0 : n-1]

		bridge := processResponse(&bytesRead)
		if bridge != nil {
			go resolveBridge(bridge)
		}
	}
}

func processResponse(b *[]byte) *BridgeInfo {
	s := string(*b)
	split := strings.Split(s, "\r\n")

	bridge := new(BridgeInfo)

	for _, field := range split {
		bridge.parseField(field)
	}

	if bridge.ID != "" {
		log.Printf("Bridge found, %v", bridge.Location)
		return bridge
	}

	return nil
}

func GetBridges() []models.Bridge {
	values := []models.Bridge{}

	for _, value := range bridges {
		values = append(values, value)
	}

	return values
}
