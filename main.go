package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	var sliceOfBytes []byte
	sliceOfBytes = []byte{1, 2, 3, 4, 'a', 'b', 'c', 'd'}
	sliceOfBytes = createSliceOfBytesFromBuffer() // Overwrite
	sliceOfBytes = append(sliceOfBytes, 'B', 'Y', 'E')

	lengthOfSlicOfBytes := len(sliceOfBytes)
	sliceOfBytes[lengthOfSlicOfBytes-1] = 'F'
	// sliceOfBytes[lengthOfSlicOfBytes] = 'F' // Invalid
	// sliceOfBytes[lengthOfSlicOfBytes+1] = 'F' // Invalid
	log.Println("The length of items in the byte slice is " + fmt.Sprint(lengthOfSlicOfBytes))

	for index, byteValue := range sliceOfBytes {
		indexString := fmt.Sprint(index)
		log.Println("Value at " + indexString + ": '" + string(byteValue) + "' == " + fmt.Sprint(byteValue))
	}

	emptyArrayOfBytes := [5]byte{}
	emptySliceOfBytes := make([]byte, 5)

	for index, byteValue := range emptyArrayOfBytes {
		indexString := fmt.Sprint(index)
		log.Println("Empty array item #" + indexString + ": " + fmt.Sprint(byteValue) + " (slice: " + fmt.Sprint(emptySliceOfBytes[index]) + ")")
	}
}

func createSliceOfBytesFromBuffer() []byte {
	var bytesBuffer bytes.Buffer
	bytesBuffer.WriteByte(42)
	bytesBuffer.WriteByte(43)
	bytesBuffer.Write([]byte{1, 2, 48, 49, 50, 51})

	bytesBuffer.WriteByte('a')  // ascii # 97
	bytesBuffer.WriteByte('\n') // newline "line feed" (LF), ascii # 10
	bytesBuffer.WriteByte(13)   // newline "carriage return" (CR), ascii # 13
	bytesBuffer.WriteByte(255)  // Largest byte value (hex: 0xFF, octal: 0377, binary: 0b11111111, 'ÿ')
	// bytesBuffer.WriteByte(256) // Invalid
	// bytesBuffer.WriteByte(666) // Invalid

	bytesBuffer.WriteByte(0xC4)       // hex, (13*16)+4 == 192+4 == 196 == 'Ä'
	bytesBuffer.WriteByte(0b11000101) // binary, 128+64+0+0+0+4+0+1 == 197 == 'Å' == 0xC5
	bytesBuffer.WriteByte(0377)       // octal, (3*64)+(7*8)+7 == 255
	bytesBuffer.WriteByte(7)          // decimal 7
	bytesBuffer.WriteByte(07)         // octal 7
	bytesBuffer.WriteByte(007)        // octal, (0*8)+7 == 7
	bytesBuffer.WriteByte(0007)       // octal, (0*64)+(0*8)+7 == 7

	// bytesBuffer.WriteByte(func() {}) //Invalid
	// bytesBuffer.Write(func() {}) //Invalid
	// bytesBuffer.WriteByte(false) // Invalid
	// bytesBuffer.Write([]bool{true, false, true}) //Invalid

	var bytesLiteral = []byte{44, 'b', 0xC5, 'Æ'}
	bytesBuffer.WriteByte(bytesLiteral[0])
	bytesBuffer.WriteByte(bytesLiteral[1])
	bytesBuffer.Write(bytesLiteral[2:]) // Skip the first two items

	return bytesBuffer.Bytes()
}
