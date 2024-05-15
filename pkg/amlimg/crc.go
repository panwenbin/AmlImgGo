package amlimg

import (
	"errors"
	"hash/crc32"
	"io"
)

func CRC32(crc uint32, p []byte) uint32 {
	tab := crc32.MakeTable(crc32.IEEE)
	return ^crc32.Update(^crc, tab, p)
}

func CRC32Img(reader io.Reader) (uint32, error) {
	crc := uint32(0xffffffff)
	var buf [4096]byte
	for {
		n, err := reader.Read(buf[:])
		crc = CRC32(crc, buf[:n])
		if errors.Is(err, io.EOF) {
			break
		} else if err != nil {
			return 0, err
		}
	}

	return crc, nil
}
