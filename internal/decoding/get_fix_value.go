package decoding

import (
	"errors"
	"fmt"

	data "golay_code/data"
)

func getFixValue(syndromeValue uint16) (uint32, error) {
	fixValue, inFixValues := data.FixValues[syndromeValue]
	if !inFixValues {
		fixValueErrorMessage := fmt.Sprintf("[decoding/getFixValue]: Can't find %d in FixValues", syndromeValue)
		return 0, errors.New(fixValueErrorMessage)
	}
	return fixValue, nil
}
