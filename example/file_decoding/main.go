package main

import (
	"os"

	app "golay_code/app"
)

func main() {
	sourceFD, _ := os.Open("./source_file.txt")
	destinationFD, _ := os.Create("./destination_file.txt")

	err := app.FileDecoding(sourceFD, destinationFD, true)
	if err != nil {
		panic(err)
	}
}
