package main

import "cfsdk/kernel"

func main() {
	kernel.K_main(kernel.KernelConfig{
		Class: kernel.Controller_DeviceClass,
		MCU:   kernel.RP2040_MCUType,
	})
}
