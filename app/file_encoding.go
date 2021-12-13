package app

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	encode "golay_code/internal/encoding"
)

func FileEncoding(sourceFile *os.File, destinationFile *os.File, useVerboseMode bool) error {
	bytesFromSourceFile, err := io.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	countOfBytesFromSource := float64(len(bytesFromSourceFile))
	bytesForDestinationFile := make([]byte, 0)

	for i := range bytesFromSourceFile {
		if useVerboseMode {
			switch {
			case i >= int((countOfBytesFromSource/100)*25) && flag25:
				fmt.Println(statusMessage25)
				flag25 = false
			case i >= int((countOfBytesFromSource/100)*50) && flag50:
				fmt.Println(statusMessage50)
				flag50 = false
			case i >= int((countOfBytesFromSource/100)*75) && flag75:
				fmt.Println(statusMessage75)
				flag75 = false
			}
		}

		encodedValue, err := encode.Encoding(uint16(bytesFromSourceFile[i]))
		if err != nil {
			return err
		}

		encodedValueAsByteSlc := make([]byte, 4)
		binary.LittleEndian.PutUint32(encodedValueAsByteSlc, encodedValue)

		for _, encodedByteToWrite := range encodedValueAsByteSlc {
			bytesForDestinationFile = append(bytesForDestinationFile, encodedByteToWrite)
		}
	}

	countOfWrittenBytes, err := destinationFile.Write(bytesForDestinationFile)
	if err != nil {
		return err
	}

	if countOfWrittenBytes != len(bytesForDestinationFile) {
		return errors.New(wrongBytesCountMessage)
	}

	fmt.Println(statusMessage100)
	return nil
}
