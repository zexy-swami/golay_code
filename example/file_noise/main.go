package main

import (
	"os"

	app "golay_code/app"
)

func main() {
	sourceFD, _ := os.Open("./source_file.txt")
	destinationFD, _ := os.Create("./destination_file.txt")

	err := app.CreateFileWithNoise(sourceFD, destinationFD)
	if err != nil {
		panic(err)
	}
}
