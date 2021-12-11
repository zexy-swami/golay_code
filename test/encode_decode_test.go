package test

import (
	"testing"

	decode "golay_code/internal/decoding"
	encode "golay_code/internal/encoding"
	error "golay_code/internal/error_generator"

	"github.com/stretchr/testify/require"
)

func TestEncodingDecoding(t *testing.T) {
	for argumentValue := uint16(0); argumentValue <= maxValueToEncode; argumentValue++ {
		encodedValue, _ := encode.Encoding(argumentValue)
		decodedValue, _ := decode.Decoding(encodedValue)
		require.Equal(t, decodedValue, argumentValue)
	}
}

func TestEncodingDecodingWithErrors(t *testing.T) {
	for argumentValue := uint16(0); argumentValue <= maxValueToEncode; argumentValue++ {
		encodedValue, _ := encode.Encoding(argumentValue)
		encodedValue = error.AddError(encodedValue)
		decodedValue, _ := decode.Decoding(encodedValue)
		require.Equal(t, decodedValue, argumentValue)
	}
}
