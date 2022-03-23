package main

import "sort"

var infiniband8636 = map[byte]string{
	(1 << 7): "Infiniband: Reserved (Bit 7)",
	(1 << 6): "Infiniband: Reserved (Bit 6)",
	(1 << 5): "Infiniband: HDR",
	(1 << 4): "Infiniband: EDR",
	(1 << 3): "Infiniband: FDR",
	(1 << 2): "Infiniband: QDR",
	(1 << 1): "Infiniband: DDR",
	(1 << 0): "Infiniband: SDR",
}

func getInfinibandSpeeds(v byte) []string {
	r := []string{}

	keys := []byte{}
	for k := range infiniband8636 {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool {
		return keys[i] < keys[j]
	})

	for _, k := range keys {
		if k&v != 0 {
			r = append(r, infiniband8636[k])
		}
	}

	return r
}
