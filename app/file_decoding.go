package app

import (
	"encoding/binary"
	"errors"
	"fmt"
	"io"
	"os"

	decode "golay_code/internal/decoding"
)

func FileDecoding(sourceFile *os.File, destinationFile *os.File, useVerboseMode bool) error {
	bytesFromSourceFile, err := io.ReadAll(sourceFile)
	if err != nil {
		return err
	}

	slcLen := float64(len(bytesFromSourceFile))
	bytesForDestinationFile := make([]byte, 0)

	encodedValueAsByteSlc := make([]byte, 0)
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

		encodedValueAsByteSlc = append(encodedValueAsByteSlc, bytesFromSourceFile[i])

		if len(encodedValueAsByteSlc) == 4 {
			valueToDecode := binary.LittleEndian.Uint32(encodedValueAsByteSlc)

			decodedValue, err := decode.Decoding(valueToDecode)
			if err != nil {
				return err
			}

			bytesForDestinationFile = append(bytesForDestinationFile, byte(decodedValue))
			encodedValueAsByteSlc = []byte{}
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
