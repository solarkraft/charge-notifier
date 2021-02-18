package main

import (
	"fmt"
	"math"
	"os/exec"
	"time"

	"github.com/0xAX/notificator"
	"github.com/distatus/battery"
)

var notify *notificator.Notificator

func main() {

	//Thresholds
	tHigh := 0.8
	tLow := 0.2
	tOff := 0.05

	sleepTime := 180.0

	notify = notificator.New(notificator.Options{DefaultIcon: "battery", AppName: "SoftBattSaver"})

	for {
		batteries, _ := battery.GetAll()
		st := sleepTime

		for i, battery := range batteries {
			fmt.Println("--- Battery", i)
			battCharge := battery.Current / battery.Full

			if math.IsNaN(battCharge) {
				fmt.Println("error")
				break
			}

			fmt.Println("Charge: ", battCharge)

			state := battery.State
			//State 4: Discharging
			//State 3: Charging

			if tLow < battCharge && battCharge < tHigh {
				fmt.Println("We're good. Battery is between", tLow, "and", tHigh)
			} else if tLow > battCharge && state == 4 { // Wi Tu Lo
				fmt.Println("Too low and discharging!")
				if tOff > battCharge && state == 4 { // Bang Ding Ow
					notify.Push("Charge level critical", "Suspending", "battery-critical", notificator.UR_NORMAL)
					fmt.Println("Suspending")
					time.Sleep(3 * time.Second)
					exec.Command("systemctl", "suspend").Run() // Linux/Systemd specific. echo "mem" > /sys/power/state would require root privileges
					time.Sleep(5 * time.Second)
					fmt.Println("Woke back up")
					notify.Push("Charge level was too low", "Your system was suspended", "battery-low", notificator.UR_NORMAL)
					st = 7
				} else {
					notify.Push("Battery low", "Start charging to preserve battery healh", "battery-low", notificator.UR_NORMAL)
					st *= battCharge / tLow
				}
			} else if tHigh < battCharge && state == 3 { // Sum Ting Wong
				fmt.Println("Too high and charging!")
				notify.Push("Battery high", "Stop charging to preserve battery healh", "battery-high", notificator.UR_NORMAL)
				st *= battCharge / tHigh
			}
		}

		fmt.Println("Sleeping for", st)
		time.Sleep(time.Duration(st) * time.Second)
	}
}
