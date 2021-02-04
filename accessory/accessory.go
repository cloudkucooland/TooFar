package accessory

import (
	hcaccessory "github.com/brutella/hc/accessory"
	"github.com/brutella/hc/log"
	"github.com/cloudkucooland/toofar/action"
)

// TFAccessory is the accessory type, TooFar's stuff, plus hc's stuff
type TFAccessory struct {
	// These are set (if required by the platform) in the accessory's config file
	Platform string           // Kasa, Tradfri, Tradfri-Device, Shelly, etc
	Name     string           // the name used internally
	IP       string           // the IP address of the device
	Username string           // for Tradfri, Shelly -- the MAC for Konnected
	Password string           // for Tradfri, Shelly -- the Token for Konnected
	Info     hcaccessory.Info // defined at https://github.com/brutella/hc/blob/master/accessory/accessory.go
	Type     hcaccessory.AccessoryType

	// relevant only to Konnected boards
	KonnectedZones []Zone

	/* below this line are NOT set in config file */
	*hcaccessory.Accessory // set when the device is added to HomeControl

	Device interface{}

	Actions []action.Action
	Runner  func(*TFAccessory, *action.Action)
}

// exposed in accessory.KonnectedZones
type Zone struct {
	ID   uint8  `json:"ID"`
	Type string `json:"type"`
	// Sensor   sensor   `json:"sensor",omitempty`
	// Actuator actuator `json:"actuator",omitempty`
	// Command  command  `json:"command",omitempty`
}

// MatchActions returns a slice of actions that should be run
// jumping through hoops since including platform here would be circular
func (a TFAccessory) MatchActions(state string) []*action.Action {
	log.Info.Printf("MatchActions: %s", state)
	var actions []*action.Action
	for _, action := range a.Actions {
		if action.TriggerState == state {
			log.Info.Printf("%s: %+v", action.TriggerState, action)
			actions = append(actions, &action)
		}
	}
	log.Info.Printf("%+v", actions)
	return actions
}
