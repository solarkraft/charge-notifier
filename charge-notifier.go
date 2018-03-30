//

package main

import (
	"github.com/0xAX/notificator"
	"github.com/distatus/battery"
	"fmt"
)

var notify *notificator.Notificator

func main() {

	//Thresholds
	tHigh := 0.8
	tLow := 0.2

	notify = notificator.New(notificator.Options { DefaultIcon: "battery", AppName: "SoftBattSaver" })

	batteries, err := battery.GetAll()
	if err != nil {
		fmt.Println("Could not get battery info!")
		return
	}

	for _, battery := range batteries {

		battCharge := battery.Current/battery.Full
		fmt.Println("Battery charge: ", battCharge)

		state := battery.State
		//State 4: Discharging
		//State 3: Charging

		if tLow < battCharge && battCharge < tHigh {
			fmt.Println("We're good. Battery is between", tLow, "and", tHigh)
		} else if 
		tLow > battCharge && state == 4 { //Wi Tu Lo
			fmt.Println("Too low and discharging!") 
			notify.Push("Battery low", "Start charging to preserve battery healh", "battery-low", notificator.UR_NORMAL)
		} else if 
		tHigh < battCharge  && state == 3{ //Bang Ding Ow
			fmt.Println("Too high and charging!")
			notify.Push("Battery high", "Stop charging to preserve battery healh", "battery-high", notificator.UR_NORMAL)
		}
	}
}