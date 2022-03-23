package main

import "encoding/binary"

func calc8636CC_BASE(eeprom []byte) byte {

	var acc uint32 = 0
	for i := 128; i <= 190; i++ {
		acc += uint32(eeprom[i])
	}
	bytebuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytebuf, acc)

	return bytebuf[0]
}

func calc8636CC_EXT(eeprom []byte) byte {

	var acc uint32 = 0
	for i := 192; i <= 222; i++ {
		acc += uint32(eeprom[i])
	}
	bytebuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(bytebuf, acc)

	return bytebuf[0]
}
