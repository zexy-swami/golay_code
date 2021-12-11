package decoding

import (
	"errors"
	"fmt"
)

const maxValueToDecode = uint32((1 << 23) - 1)

func Decoding(valueToDecode uint32) (uint16, error) {
	if valueToDecode > maxValueToDecode {
		valueRangeError := fmt.Sprintf("[decoding/Decoding]: Value to decode %d, expected values are [0, 8388607]", valueToDecode)
		return 0, errors.New(valueRangeError)
	}

	syndromeValue := getSyndromeValue(valueToDecode)

	if syndromeValue == 0 {
		return uint16(valueToDecode >> 11), nil
	}

	fixValue, err := getFixValue(syndromeValue)
	if err != nil {
		return 0, err
	}

	var decodedValue = uint16((valueToDecode ^ fixValue) >> 11)
	return decodedValue, nil
}
