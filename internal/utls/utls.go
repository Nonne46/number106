package utls

import (
	"fmt"
	"hash/crc32"
)

func CRC32Hash(input string) string {
	// Convert the input string to a byte slice
	data := []byte(input)

	// Create a new CRC32 hash
	hash := crc32.NewIEEE()

	// Write the data to the hash
	hash.Write(data)

	// Get the CRC32 checksum as an unsigned 32-bit integer
	crc32Checksum := hash.Sum32()

	// Convert the CRC32 checksum to a hexadecimal string
	crc32ChecksumHex := fmt.Sprintf("%X", crc32Checksum)

	return crc32ChecksumHex
}
