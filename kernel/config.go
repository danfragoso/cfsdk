package kernel

type KernelConfig struct {
	Class DeviceClass
	MCU   MCUType

	BufferSize uint16
	MaxAddress uint32

	GCEnabled     bool
	TimeSensitive bool
}
