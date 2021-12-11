package encoding

import (
	"errors"
	"fmt"
	"math/bits"

	data "golay_code/data"
)

const maxValueToEncode = uint16((1 << 12) - 1)

func Encoding(valueToEncode uint16) (uint32, error) {
	if valueToEncode > maxValueToEncode {
		valueRangeError := fmt.Sprintf("[encoding/Encoding]: Value to encode %d, expected values are [0, 4095]", valueToEncode)
		return 0, errors.New(valueRangeError)
	}

	var encodedValue = uint32(valueToEncode) << 11

	for i := 12; i < 23; i++ {
		if onesCount := bits.OnesCount16(valueToEncode & data.GT[i]); onesCount%2 == 1 {
			encodedValue += uint32(1) << (22 - i)
		}
	}
	return encodedValue, nil
}
