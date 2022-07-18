package kernel

type DeviceClass uint8

const (
	Sensor_DeviceClass DeviceClass = iota
	Controller_DeviceClass
	Actuator_DeviceClass
)
