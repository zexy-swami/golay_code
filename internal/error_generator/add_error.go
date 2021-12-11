package error_generator

import (
	"math/rand"
	"time"
)

func AddError(encodedValue uint32) uint32 {
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < 3; i++ {
		shiftValue := rand.Intn(23)
		encodedValue ^= uint32(1) << shiftValue
	}
	return encodedValue
}
