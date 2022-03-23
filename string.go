package main

import (
	"fmt"

	"github.com/mickep76/color"
)

var (
	colorField = color.Cyan.String()
	colorValue = color.Green.String()
	colorList  = color.Yellow.String()
)

func (m *Module) String() string {
	switch m.Type {
	case "SFF-8079":
		return m.string8079()
	case "SFF-8636":
		return m.string8636()
	}
	return ""
}

func (m *Module) string8079() string {
	f1 := fmt.Sprintf("\t\t%s%%-52s%s : %s%%v%s", colorField, color.Reset, colorValue, color.Reset)
	f2 := fmt.Sprintf("\t\t%s%%-52s%s : %s%%v%s", colorField, color.Reset, colorList, color.Reset)

	s := fmt.Sprintf("\n"+f1, "Type", m.Type) +
		fmt.Sprintf("\n"+f1, "Identifier [0]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[0]), m.Identifier)) +
		fmt.Sprintf("\n"+f1, "Extended Identifier [1]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[1]), m.ExtIdentifierDescs)) +
		fmt.Sprintf("\n"+f1, "Connector [2]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[2]), m.Connector)) +
		fmt.Sprintf("\n"+f1, "Transceiver Codes [3-10]", fmt.Sprintf("0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x", m.Eeprom[3], m.Eeprom[4], m.Eeprom[5], m.Eeprom[6], m.Eeprom[7], m.Eeprom[8], m.Eeprom[9], m.Eeprom[10]))

	if len(m.TransceiverSpecComplianceDescs) > 0 {
		s += fmt.Sprintf("\n"+f2, "Transceiver Type", m.TransceiverSpecComplianceDescs[0])
		for _, t := range m.TransceiverSpecComplianceDescs[1:] {
			s += fmt.Sprintf("\n"+f2, "", t)
		}
	} else {
		s += fmt.Sprintf("\n"+f2, "Transceiver Type", "None!")
	}

	s += fmt.Sprintf("\n"+f1, "Encoding [11]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[11]), m.Encoding)) +
		fmt.Sprintf("\n"+f1, "BR, Nominal [12]", m.BrNominal) +
		fmt.Sprintf("\n"+f1, "Rate Identifier [13]", fmt.Sprintf("0x%02x", m.RateIdentifier)) +
		fmt.Sprintf("\n"+f1, "Length (SMF) [14]", m.LengthSmfKm) +
		fmt.Sprintf("\n"+f1, "Length (SMF) [15]", m.LengthSmfM) +
		fmt.Sprintf("\n"+f1, "Length (50um) [16]", m.Length50umM) +
		fmt.Sprintf("\n"+f1, "Length (62.5um) [17]", m.Length625umM) +
		fmt.Sprintf("\n"+f1, "Length (Copper) [18]", m.LengthCopper) +
		fmt.Sprintf("\n"+f1, "Length (OM3) [19]", m.LengthOm3) +
		fmt.Sprintf("\n"+f1, "Vendor [20-35]", m.Vendor) +
		fmt.Sprintf("\n"+f1, "Vendor OUI [37-39]", m.VendorOui) +
		fmt.Sprintf("\n"+f1, "Vendor PN [40-55]", m.VendorPn) +
		fmt.Sprintf("\n"+f1, "Vendor Rev [56-59]", m.VendorRev) +
		// fmt.Sprintf("\n"+f1, "Option Values [64-65]", fmt.Sprintf("0x%02x 0x%02x", m.Options[0], m.Options[1])) +
		fmt.Sprintf("\n"+f1, "BR Margin, Max [66]", m.BrMax) +
		fmt.Sprintf("\n"+f1, "BR Margin, Min [67]", m.BrMin) +
		fmt.Sprintf("\n"+f1, "Vendor SN [68-83]", m.VendorSn) +
		fmt.Sprintf("\n"+f1, "Date Code [84-91]", m.DateCode)

	/*
		if s.Vendor.String() == "Arista Networks" && strings.HasPrefix(s.VendorPn.String(), "CAB-Q-S-") {
			str += fmt.Sprintf("%-50s : %x\n", "Vendor SA [120]", s.VendorAristaSa)
		}
	*/

	return s
}

func (m *Module) string8636() string {
	f1 := fmt.Sprintf("\t\t%s%%-60s%s : %s%%v%s", colorField, color.Reset, colorValue, color.Reset)
	f2 := fmt.Sprintf("\t\t%s%%-60s%s : %s%%v%s", colorField, color.Reset, colorList, color.Reset)

	s := fmt.Sprintf("\n"+f1, "Type", m.Type) +
		fmt.Sprintf("\n"+f1, "Identifier [128]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[0]), m.Identifier)) +
		fmt.Sprintf("\n"+f1, "Extended Identifier [129]", fmt.Sprintf("0x%02x", byte(m.Eeprom[129]))) +
		fmt.Sprintf("\n"+f2, "Extended Identifier Descriptions [129]", m.ExtIdentifierDescs[0])

	for _, i := range m.ExtIdentifierDescs[1:] {
		s += fmt.Sprintf("\n"+f2, "", i)
	}

	s += fmt.Sprintf("\n"+f1, "Connector [130]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[130]), m.Connector)) +
		fmt.Sprintf("\n"+f1, "Transceiver Specification Compliance [131-138]", fmt.Sprintf("0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x 0x%02x", m.Eeprom[131], m.Eeprom[132], m.Eeprom[133], m.Eeprom[134], m.Eeprom[135], m.Eeprom[136], m.Eeprom[137], m.Eeprom[138]))

	if len(m.TransceiverSpecComplianceDescs) > 0 {
		s += fmt.Sprintf("\n"+f2, "Transceiver Specification Compliance Descriptions [131-138]", m.TransceiverSpecComplianceDescs[0])

		for _, t := range m.TransceiverSpecComplianceDescs[1:] {
			s += fmt.Sprintf("\n"+f2, "", t)
		}
	} else {
		s += fmt.Sprintf("\n"+f2, "Transceiver Specification Compliance Descriptions", "None")
	}

	s += fmt.Sprintf("\n"+f1, "Encoding [139]", fmt.Sprintf("0x%02x (%s)", byte(m.Eeprom[139]), m.Encoding)) +
		fmt.Sprintf("\n"+f1, "Signaling Rate, Nominal (Baud) [140]", m.BrNominal) +
		fmt.Sprintf("\n"+f1, "Extended Rate Select Compliance [141]", fmt.Sprintf("0x%02x", m.RateIdentifier)) +
		fmt.Sprintf("\n"+f1, "Length (SMF) [142]", m.LengthSmf) +
		fmt.Sprintf("\n"+f1, "Length (OM3 50um) [143]", m.LengthOm3) +
		fmt.Sprintf("\n"+f1, "Length (OM2 50um) [144]", m.LengthOm2) +
		fmt.Sprintf("\n"+f1, "Length (OM1 62.5um) [145]", m.LengthOm1) +
		fmt.Sprintf("\n"+f1, "Length (Copper or Active cable, OM4 50um [2m]) [146]", m.LengthCopper) +
		fmt.Sprintf("\n"+f2, "Device Technology Descriptions [147]", m.DeviceTechnologyDescs[0])

	for _, t := range m.DeviceTechnologyDescs[1:] {
		s += fmt.Sprintf("\n"+f2, "", t)
	}

	s +=
		fmt.Sprintf("\n"+f1, "Transmitter Technology [147]", m.TransmitterTechnology) +
			fmt.Sprintf("\n"+f1, "Vendor [148-163]", m.Vendor)

	if len(m.InfinibandSpeeds) > 0 {
		s += fmt.Sprintf("\n"+f2, "Extended Module (Infiniband Speeds) [164]", m.InfinibandSpeeds[0])
		for _, t := range m.InfinibandSpeeds[1:] {
			s += fmt.Sprintf("\n"+f2, "", t)
		}
	} else {
		s += fmt.Sprintf("\n"+f1, "Extended Module (Infiniband Speeds) [164]", "None!")
	}

	s +=
		fmt.Sprintf("\n"+f1, "Vendor OUI [165-167]", fmt.Sprintf("%s (Vendor is %s)", m.VendorOui, m.VendorNameFromOui)) +
			fmt.Sprintf("\n"+f1, "Vendor PN [168-183]", m.VendorPn) +
			fmt.Sprintf("\n"+f1, "Vendor Rev [184-185]", m.VendorRev) +
			fmt.Sprintf("\n"+f1, "Wavelength nm [186-187]", m.Wavelength) +
			fmt.Sprintf("\n"+f1, "Attentuation 2.5GHz dB [186]", m.Attentuation25GHz) +
			fmt.Sprintf("\n"+f1, "Attentuation 5.0GHz dB [187]", m.Attentuation50GHz) +
			fmt.Sprintf("\n"+f1, "Wavelength tolerance nm [188:189]", m.WavelengthTolerance) +
			fmt.Sprintf("\n"+f1, "Attentuation 7.0GHz dB [188]", m.Attentuation70GHz) +
			fmt.Sprintf("\n"+f1, "Attentuation 12.9GHz dB [189]", m.Attentuation129GHz) +
			fmt.Sprintf("\n"+f1, "Maximum Case Temperature °C [190]", m.MaxCaseTemperature)

	if m.CC_BASE == m.calced_CC_BASE {
		s += fmt.Sprintf("\n"+f1, "CC_BASE Checksum [191]", fmt.Sprintf("correct (Is 0x%02x, should be 0x%02x)", m.CC_BASE, m.calced_CC_BASE))
	} else {
		s += fmt.Sprintf("\n"+f1, "CC_BASE Checksum [191]", fmt.Sprintf("INCORRECT (Is 0x%02x, should be 0x%02x)", m.CC_BASE, m.calced_CC_BASE))
	}

	s +=
		fmt.Sprintf("\n"+f1, "Link Codes [192]", m.LinkCodes)

	if len(m.Options) > 0 {
		s += fmt.Sprintf("\n"+f2, "Options [193-195]", m.Options[0])
		for _, t := range m.Options[1:] {
			s += fmt.Sprintf("\n"+f2, "", t)
		}
	} else {
		s += fmt.Sprintf("\n"+f1, "Options [193-195]", "None!")
	}

	s += fmt.Sprintf("\n"+f1, "Vendor SN [196-211]", m.VendorSn) +
		fmt.Sprintf("\n"+f1, "Date Code [212-219]", m.DateCode)

	if m.CC_EXT == m.calced_CC_EXT {
		s += fmt.Sprintf("\n"+f1, "CC_EXT Checksum [223]", fmt.Sprintf("correct (Is 0x%02x, should be 0x%02x)", m.CC_EXT, m.calced_CC_EXT))
	} else {
		s += fmt.Sprintf("\n"+f1, "CC_EXT Checksum [223]", fmt.Sprintf("INCORRECT (Is 0x%02x, should be 0x%02x)", m.CC_EXT, m.calced_CC_EXT))
	}

	return s
}
