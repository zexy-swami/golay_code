package test

import (
	"testing"

	decode "golay_code/internal/decoding"
	encode "golay_code/internal/encoding"

	"github.com/stretchr/testify/require"
)

const (
	maxValueToEncode = uint16((1 << 12) - 1)
	maxValueToDecode = uint32((1 << 23) - 1)
)

func TestPossibleValuesToEncode(t *testing.T) {
	for argumentValue := uint16(0); argumentValue < maxValueToEncode; argumentValue++ {
		_, err := encode.Encoding(argumentValue)
		require.Nil(t, err)
	}
}

func TestPossibleValuesToDecode(t *testing.T) {
	for argumentValue := uint32(0); argumentValue < maxValueToDecode; argumentValue++ {
		_, err := decode.Decoding(argumentValue)
		require.Nil(t, err)
	}
}

func TestEncodeWrongValue(t *testing.T) {
	var wrongRangeArgument uint16 = maxValueToEncode + 1
	var errorMessage string = "[encoding/Encoding]: Value to encode 4096, expected values are [0, 4095]"

	if _, err := encode.Encoding(wrongRangeArgument); err != nil {
		require.Equal(t, err.Error(), errorMessage)
	} else {
		require.Equal(t, true, false)
	}
}

func TestDecodeWrongValue(t *testing.T) {
	var wrongRangeArgument uint32 = maxValueToDecode + 1
	var errorMessage string = "[decoding/Decoding]: Value to decode 8388608, expected values are [0, 8388607]"

	if _, err := decode.Decoding(wrongRangeArgument); err != nil {
		require.Equal(t, err.Error(), errorMessage)
	} else {
		require.Equal(t, true, false)
	}
}
