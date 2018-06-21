package bridgediscovery

import (
	"strings"
)

// Bridge holds information about found bridges in the network
type BridgeInfo struct {
	ID       string
	Location string
	USN      string
}

func (bridge *BridgeInfo) parseField(field string) {
	if strings.HasPrefix(field, "hue-bridgeid:") {
		bridge.ID = readAttribute(field, "hue-bridgeid")
	} else if strings.HasPrefix(field, "LOCATION:") {
		bridge.Location = readAttribute(field, "LOCATION")
	} else if strings.HasPrefix(field, "USN:") {
		bridge.USN = readAttribute(field, "USN")
	}
}

func readAttribute(fromString, attribute string) string {
	result := strings.Replace(fromString, attribute+":", "", -1)
	result = strings.Trim(result, " ")
	return result
}
