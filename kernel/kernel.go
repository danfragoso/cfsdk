package kernel

import (
	"machine"
	"time"
)

var kConfig KernelConfig

func K_main(cfg KernelConfig) {
	kConfig = cfg

	led := machine.LED
	led.Configure(machine.PinConfig{Mode: machine.PinOutput})

	for {
		switch kConfig.Class {
		case Controller_DeviceClass:
			led.Low()
			time.Sleep(time.Millisecond * 1000)

			led.High()
			time.Sleep(time.Millisecond * 2000)

		case Actuator_DeviceClass:
			led.Low()
			time.Sleep(time.Millisecond * 100)

			led.High()
			time.Sleep(time.Millisecond * 200)
		}

	}
}
