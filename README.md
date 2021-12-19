# Golay code

## Information
***Perfect Golay code*** - linear error-correcting code used in digital communications.  
This code's parameters are [23, 12, 7]:
```
23 - length of encoded word;
12 - length of original word;
7  - Hamming distance.
```


## Using example
###### Encoding
Encoding example presented in ***/example/file_encoding/main.go***:
```go
sourceFD, _ := os.Open("./source_file.txt")
destinationFD, _ := os.Create("./destination_file.txt")

err := app.FileEncoding(sourceFD, destinationFD, true)
```
Function ***FileEncoding*** placed in ***/app/file_encoding.go***  
With running main.go, file ***destination_file.txt*** with encoded data will be generated.  

###### Noise adding
For checking Golay code in action some noise can be added.  
***destination_file.txt*** from encoding example must be placed into ***/example/file_noise*** with name ***source_file.txt***  
Code for noise adding placed in ***/example/file_noise/main.go***:
```go
sourceFD, _ := os.Open("./source_file.txt")
destinationFD, _ := os.Create("./destination_file.txt")

err := app.CreateFileWithNoise(sourceFD, destinationFD)
```
Function ***CreateFileWithNoise*** placed in ***/app/file_noise_adding.go***

###### Decoding
Decoding example presented in ***/example/file_decoding/main.go***:
```go
sourceFD, _ := os.Open("./source_file.txt")
destinationFD, _ := os.Create("./destination_file.txt")

err := app.FileDecoding(sourceFD, destinationFD, true)
```
***destination_file.txt*** from encoding example or noise adding example can be used as ***source_file.txt*** in this example.  
Function ***FileDecoding*** placed in ***/app/file_decoding.go***.