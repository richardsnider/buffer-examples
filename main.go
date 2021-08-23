package main

import (
	"bytes"
	"fmt"
	"log"
)

func main() {
	sliceOfBytes := []byte{1, 2, 3, 4, 'a', 'b', 'c', 'd'}

	sliceOfBytes = createSliceOfBytesFromBuffer()
	sliceOfBytes = append(sliceOfBytes, 'B', 'Y', 'E')

	log.Println("The length of items in the byte slice is " + fmt.Sprint(len(sliceOfBytes)))

	for index, byteValue := range sliceOfBytes {
		indexString := fmt.Sprint(index)
		log.Println("Value at " + indexString + ": \"" + string(byteValue) + "\" == " + fmt.Sprint(byteValue))
	}

}

func createSliceOfBytesFromBuffer() []byte {
	var bytesBuffer bytes.Buffer
	bytesBuffer.WriteByte(42)
	bytesBuffer.WriteByte(13)
	bytesBuffer.WriteByte('a')  // ascii # 97
	bytesBuffer.WriteByte('\n') // ascii # 10
	bytesBuffer.WriteByte(255)  // Largest byte value (hex: 0xFF, octal: 077, binary: 0b11111111, "ÿ")
	// bytesBuffer.WriteByte(256) // Invalid
	// bytesBuffer.WriteByte(666) // Invalid
	bytesBuffer.WriteByte(0xC4)       // (13*16)+4 == 192+4 == 196 == "Ä"
	bytesBuffer.WriteByte(007)        // 0+7 == 7
	bytesBuffer.WriteByte(017)        // 8+7 == 15
	bytesBuffer.WriteByte(0b11000101) // 128+64+0+0+0+4+0+1 == 197 == "Å" == 0xC5

	var bytesLiteral = []byte{43, 'b', 14, 0xC5}
	bytesBuffer.WriteByte(bytesLiteral[0])
	bytesBuffer.Write(bytesLiteral)

	return bytesBuffer.Bytes()
}
