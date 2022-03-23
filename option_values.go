package main

const (
	RxOutputAmplitudeFixedProgrammableSettingsImplemented = 0b00000001
	RxOutputEmphasisFixedProgrammableSettingsImplemented  = 0b00000010
	TxInputEqualizersFixedProgrammableSettingsImplemented = 0b00000100
	TxInputEqualizersAutoAdaptiveCapableImplemented       = 0b00001000
	TxInputAdaptiveEqualizersFreezeCapable                = 0b00010000
	IntLRxLOSLOutputSignalIsConfigurable                  = 0b00100000
	LPModeTxDisInputSignalIsConfigurable                  = 0b01000000
	B193B7ReservedSet                                     = 0b10000000

	TxSquelchImplemented           = 0b00000001
	TxSquelchDisableImplemented    = 0b00000010
	RxOutputDisableImplemented     = 0b00000100
	RxSquelchDisableImplemented    = 0b00001000
	RxCDRLossOfLockFlagImplemented = 0b00010000
	TxCDRLossOfLockFlagImplemented = 0b00100000
	RxCDROnOffControlImplemented   = 0b01000000
	TxCDROnOffControlImplemented   = 0b10000000

	Pages2021hImplemented                             = 0b00000001
	TxLossOfSignalImplemented                         = 0b00000010
	TxSquelchImplementedToReducePave                  = 0b00000100
	Tx_FaultSignalImplemented                         = 0b00001000
	Tx_DisableIsImplementedAndDisablesTheSerialOutput = 0b00010000
	RateSelectIsImplemented                           = 0b00100000
	MemoryPage01hProvided                             = 0b01000000
	MemoryPage02Provided                              = 0b10000000
)

func getDeviceOptions8636(v []byte) []string {
	r := []string{}

	if v[0]&RxOutputAmplitudeFixedProgrammableSettingsImplemented != 0 {
		r = append(r, "Rx output amplitude fixed-programmable settings implemented")
	}

	if v[0]&RxOutputEmphasisFixedProgrammableSettingsImplemented != 0 {
		r = append(r, "Rx output emphasis fixed-programmable settings implemented")
	}

	if v[0]&TxInputEqualizersFixedProgrammableSettingsImplemented != 0 {
		r = append(r, "Tx input equalizers fixed-programmable settings implemented")
	}

	if v[0]&TxInputEqualizersAutoAdaptiveCapableImplemented != 0 {
		r = append(r, "Tx input equalizers are auto-adaptive capable")
	}

	if v[0]&TxInputAdaptiveEqualizersFreezeCapable != 0 {
		r = append(r, "Tx input adaptive equalizers are freeze capable")
	}

	if v[0]&IntLRxLOSLOutputSignalIsConfigurable != 0 {
		r = append(r, "IntL/RxLOSL output signal is configurable using byte 99")
	}

	if v[0]&LPModeTxDisInputSignalIsConfigurable != 0 {
		r = append(r, "LPMode/TxDis input signal is configurable using byte 99")
	}

	if v[0]&B193B7ReservedSet != 0 {
		r = append(r, "Byte 193 Bit 7 (Reserved)")
	}
	//
	//
	//
	if v[1]&TxSquelchImplemented != 0 {
		r = append(r, "Tx Squelch implemented")
	}

	if v[1]&TxSquelchDisableImplemented != 0 {
		r = append(r, "Tx Squelch Disable implemented")
	}

	if v[1]&RxOutputDisableImplemented != 0 {
		r = append(r, "Rx Output Disable implemented")
	}

	if v[1]&RxSquelchDisableImplemented != 0 {
		r = append(r, "Rx Squelch Disable implemented")
	}

	if v[1]&RxCDRLossOfLockFlagImplemented != 0 {
		r = append(r, "Rx CDR Loss of Lock (LOL) flag implemented")
	}

	if v[1]&TxCDRLossOfLockFlagImplemented != 0 {
		r = append(r, "Tx CDR Loss of Lock (LOL) flag implemented")
	}

	if v[1]&RxCDROnOffControlImplemented != 0 {
		r = append(r, "Rx CDR On/Off Control implemented")
	}

	if v[1]&TxCDROnOffControlImplemented != 0 {
		r = append(r, "Tx CDR On/Off Control implemented")
	}
	//
	//
	//
	if v[2]&Pages2021hImplemented != 0 {
		r = append(r, "Pages 20-21h implemented")
	}

	if v[2]&TxLossOfSignalImplemented != 0 {
		r = append(r, "Tx Loss of Signal implemented")
	}

	if v[2]&TxSquelchImplementedToReducePave != 0 {
		r = append(r, "Tx Squelch implemented to reduce Pave")
	} else {
		r = append(r, "Tx Squelch implemented to reduce OMA")
	}

	if v[2]&Tx_FaultSignalImplemented != 0 {
		r = append(r, "Tx_Fault signal implemented")
	}

	if v[2]&Tx_DisableIsImplementedAndDisablesTheSerialOutput != 0 {
		r = append(r, "Tx_Disable is implemented and disables the serial output")
	}

	if v[2]&RateSelectIsImplemented != 0 {
		r = append(r, "Rate select is implemented")
	}

	if v[2]&MemoryPage01hProvided != 0 {
		r = append(r, "Memory Page 01h provided")
	}

	if v[2]&MemoryPage02Provided != 0 {
		r = append(r, "Memory Page 02 provided")
	}

	return r
}
