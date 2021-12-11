package data

// GT - transposed generator matrix
var GT = [23]uint16{
	0b100000000000,
	0b010000000000,
	0b001000000000,
	0b000100000000,
	0b000010000000,
	0b000001000000,
	0b000000100000,
	0b000000010000,
	0b000000001000,
	0b000000000100,
	0b000000000010,
	0b000000000001,
	0b111110010010,
	0b011111001001,
	0b110001110110,
	0b011000111011,
	0b110010001111,
	0b100111010101,
	0b101101111000,
	0b010110111100,
	0b001011011110,
	0b000101101111,
	0b111100100101,
}

// HT - transposed parity check matrix
var HT = [11]uint32{
	0b11111001001010000000000,
	0b01111100100101000000000,
	0b11000111011000100000000,
	0b01100011101100010000000,
	0b11001000111100001000000,
	0b10011101010100000100000,
	0b10110111100000000010000,
	0b01011011110000000001000,
	0b00101101111000000000100,
	0b00010110111100000000010,
	0b11110010010100000000001,
}
