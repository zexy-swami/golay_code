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

	slcLen := float64(len(bytesFromSourceFile))
	bytesForDestinationFile := make([]byte, 0)

	for i := range bytesFromSourceFile {
		if useVerboseMode {
			if i >= int((slcLen/100)*25) && flag25 {
				fmt.Println("25% of sending is done.")
				flag25 = false
			} else if i >= int((slcLen/100)*50) && flag50 {
				fmt.Println("50% of sending is done.")
				flag50 = false
			} else if i >= int((slcLen/100)*75) && flag75 {
				fmt.Println("75% of sending is done.")
				flag75 = false
			}
		}

		encodedValue, err := encode.Encoding(uint16(bytesFromSourceFile[i]))
		if err != nil {
			return err
		}

		encodedValueAsByteSlc := make([]byte, 4)
		binary.LittleEndian.PutUint32(encodedValueAsByteSlc, encodedValue)

		for _, encodedByte := range encodedValueAsByteSlc {
			bytesForDestinationFile = append(bytesForDestinationFile, encodedByte)
		}
	}

	countOfWrittenBytes, err := destinationFile.Write(bytesForDestinationFile)
	if err != nil {
		return err
	}

	if countOfWrittenBytes != len(bytesForDestinationFile) {
		return errors.New("wrong count of bytes have been written")
	}

	fmt.Println("100% of sending is done.")
	return nil
}
