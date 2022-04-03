package relay

import (
	"fmt"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/raspi"
)

func Start() {
	r := raspi.NewAdaptor()
	relay := gpio.NewRelayDriver(r, "11")

	work := func() {
		if err := relay.On(); err != nil {
			fmt.Printf("error: %v", err)
		}
	}

	robot := gobot.NewRobot("startRelay",
		[]gobot.Connection{r},
		[]gobot.Device{relay},
		work,
	)

	if err := robot.Start(); err != nil {
		fmt.Println(err)
	}
	if err := robot.Stop(); err != nil {
		fmt.Println(err)
	}
}

func Stop() {
	r := raspi.NewAdaptor()
	relay := gpio.NewRelayDriver(r, "11")

	work := func() {
		if err := relay.Off(); err != nil {
			fmt.Printf("error: %v", err)
		}
	}

	robot := gobot.NewRobot("stopRelay",
		[]gobot.Connection{r},
		[]gobot.Device{relay},
		work,
	)

	if err := robot.Start(); err != nil {
		fmt.Println(err)
	}
	if err := robot.Stop(); err != nil {
		fmt.Println(err)
	}
}
