package drivers

import "cfsdk/kernel"

type IWM50001 struct {
	id uint8

	open bool
}

func (i *IWM50001) ID() uint8 {
	return i.id
}

func (i *IWM50001) Class() kernel.DeviceClass {
	return kernel.Actuator_DeviceClass
}

func (i *IWM50001) Type() kernel.DeviceType {
	return kernel.FuelInjectionNozzle_DeviceType
}

func (i *IWM50001) Read() []byte {
	if i.open {
		return []byte{0x1}
	}

	return []byte{0x0}
}

func (i *IWM50001) Write(buf []byte) {
	if buf[0] == 0x1 {
		i.open = true
		return
	}

	i.open = false
}
