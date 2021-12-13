package app

import (
	"encoding/binary"
	"errors"
	"io"
	"os"

	errorgen "golay_code/internal/error_generator"
)

func CreateFileWithNoise(sourceFile *os.File, destinationFile *os.File) error {
	bytesFromSourceFile, err := io.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	bytesForDestinationFile := make([]byte, 0)
	encodedValueAsByteSlc := make([]byte, 0)

	for i := range bytesFromSourceFile {
		encodedValueAsByteSlc = append(encodedValueAsByteSlc, bytesFromSourceFile[i])

		if len(encodedValueAsByteSlc) == 4 {
			encodedValueAsUint32 := binary.LittleEndian.Uint32(encodedValueAsByteSlc)
			encodedValueWithNoise := errorgen.AddError(encodedValueAsUint32)

			encodedValueWithNoiseAsSlc := make([]byte, 4)
			binary.LittleEndian.PutUint32(encodedValueWithNoiseAsSlc, encodedValueWithNoise)

			for _, encodedByteToWrite := range encodedValueWithNoiseAsSlc {
				bytesForDestinationFile = append(bytesForDestinationFile, encodedByteToWrite)
			}
			encodedValueAsByteSlc = []byte{}
		}
	}

	countOfWrittenBytes, err := destinationFile.Write(bytesForDestinationFile)
	if err != nil {
		return err
	}

	if countOfWrittenBytes != len(bytesForDestinationFile) {
		return errors.New(wrongBytesCountMessage)
	}

	return nil
}
