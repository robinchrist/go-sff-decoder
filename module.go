package main

import (
	"encoding/hex"
	"encoding/json"
)

type ModuleType string

const (
	TypeSff8079 = ModuleType("SFF-8079")
	TypeSff8636 = ModuleType("SFF-8636")
)

type Bytes []byte

type Module struct {
	Type                           ModuleType
	Identifier                     string
	ExtIdentifierDescs             []string
	Connector                      string
	TransceiverSpecComplianceDescs []string
	Encoding                       string
	BrNominal                      int
	RateIdentifier                 byte
	LengthSmfKm                    int
	LengthSmfM                     int
	LengthSmf                      int
	Length50umM                    int
	Length625umM                   int
	LengthCopper                   int
	LengthOm3                      int
	LengthOm2                      int
	LengthOm1                      int
	DeviceTechnologyDescs          []string
	TransmitterTechnology          string
	InfinibandSpeeds               []string
	Vendor                         string
	VendorOui                      string
	VendorNameFromOui              string
	VendorPn                       string
	VendorRev                      string
	Wavelength                     float32
	Attentuation25GHz              int
	Attentuation50GHz              int
	WavelengthTolerance            float32
	Attentuation70GHz              int
	Attentuation129GHz             int
	MaxCaseTemperature             int
	CC_BASE                        byte
	calced_CC_BASE                 byte
	LinkCodes                      string
	Options                        []string
	BrMax                          int
	BrMin                          int
	VendorSn                       string
	CC_EXT                         byte
	calced_CC_EXT                  byte
	DateCode                       string
	VendorSa                       int
	Eeprom                         Bytes
}

func (b *Bytes) MarshalJSON() ([]byte, error) {
	return json.Marshal(hex.EncodeToString(*b))
}

func (b *Bytes) UnmarshalJSON(in []byte) error {
	var s string
	if err := json.Unmarshal(in, &s); err != nil {
		return err
	}

	v, err := hex.DecodeString(s)
	if err != nil {
		return err
	}
	*b = v

	return nil
}
