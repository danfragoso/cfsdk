package kernel

type MCUType uint8

const (
	RP2040_MCUType MCUType = iota
	STM32F103_MCUType
)
