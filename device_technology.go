package main

const (
	DeviceTechMask = 0b00000111

	TransmitterNotTunable = 0b00000000
	TransmitterTunable    = 0b00000001

	PinDetector = 0b00000000
	APDDetector = 0b00000010

	UncooledTransmitterDevice = 0b00000000
	CooledTransmitter         = 0b00000100

	NoWavelengthControl     = 0b00000000
	ActiveWavelengthControl = 0b00001000

	TransmitterTechMask = 0b11110000
)

var transmitterTechnology = map[byte]string{
	(0b0000 << 4): "850 nm VCSEL",
	(0b0001 << 4): "1310 nm VCSEL",
	(0b0010 << 4): "1550 nm VCSEL",
	(0b0011 << 4): "1310 nm FP",
	(0b0100 << 4): "1310 nm DFB",
	(0b0101 << 4): "1550 nm DFB",
	(0b0110 << 4): "1310 nm EML",
	(0b0111 << 4): "1550 nm EML",
	(0b1000 << 4): "Other / Undefined",
	(0b1001 << 4): "1490 nm DFB",
	(0b1010 << 4): "Copper cable unequalized",
	(0b1011 << 4): "Copper cable passive equalized",
	(0b1100 << 4): "Copper cable, near and far end limiting active equalizers",
	(0b1101 << 4): "Copper cable, far end limiting active equalizers",
	(0b1110 << 4): "Copper cable, near end limiting active equalizers",
	(0b1111 << 4): "Copper cable, linear active equalizers",
}

func getDeviceTechnologyDescs8636(v byte) []string {
	r := []string{}

	if v&TransmitterTunable != 0 {
		r = append(r, "Transmitter tunable")
	} else {
		r = append(r, "Transmitter not tunable")
	}

	if v&APDDetector != 0 {
		r = append(r, "APD Detector")
	} else {
		r = append(r, "Pin Detector")
	}

	if v&CooledTransmitter != 0 {
		r = append(r, "Cooled transmitter")
	} else {
		r = append(r, "Uncooled transmitter device")
	}

	if v&ActiveWavelengthControl != 0 {
		r = append(r, "Active Wavelength control")
	} else {
		r = append(r, "No wavelength control")
	}

	return r
}

func getTransmitterTechnology8636(v byte) string {
	name, ok := transmitterTechnology[v&TransmitterTechMask]
	if !ok {
		return "Unknown!"
	}

	return name
}
