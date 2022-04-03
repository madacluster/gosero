package relay

import (
	"fmt"

	"github.com/madacluster/gosero/pkg/helpers"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

var RELAYPORT = helpers.GetEnv("RELAY_PORT", "11")

func Start() error {
	r := raspi.NewAdaptor()
	if err := r.Connect(); err != nil {
		return err
	}
	relay := gpio.NewRelayDriver(r, RELAYPORT)
	if err := relay.On(); err != nil {
		fmt.Printf("error: %v", err)
		return err
	}

	if err := relay.Start(); err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func Stop() {
	r := raspi.NewAdaptor()
	relay := gpio.NewRelayDriver(r, RELAYPORT)
	r.Connect()
	if err := relay.Off(); err != nil {
		fmt.Printf("error: %v", err)
	}

	if err := relay.Start(); err != nil {
		fmt.Println(err)
	}

}
