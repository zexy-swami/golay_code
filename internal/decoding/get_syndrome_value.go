package decoding

import (
	"math/bits"

	data "golay_code/data"
)

func getSyndromeValue(encodedValue uint32) uint16 {
	var syndromeValue uint16

	for i := 0; i < 11; i++ {
		if onesCount := bits.OnesCount32(encodedValue & data.HT[i]); onesCount%2 == 1 {
			syndromeValue += uint16(1) << (10 - i)
		}
	}
	return syndromeValue
}
