package kernel

type Device interface {
	ID() uint8
	Read() []byte
	Write([]byte)
	Class() DeviceClass
	Type() DeviceType
}
