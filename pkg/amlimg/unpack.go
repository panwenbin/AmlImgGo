package amlimg

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
)

func Unpack(file *os.File, checkCRC bool) (*Img, error) {
	header, err := ReadHeader(file)
	if err != nil {
		return nil, err
	}

	if checkCRC {
		_, err = file.Seek(int64(binary.Size(header.CRC)), io.SeekStart)
		if err != nil {
			return nil, err
		}
		crc, err := CRC32Img(file)
		if err != nil {
			return nil, err
		}
		if header.CRC != crc {
			return nil, fmt.Errorf("invalid crc: expect %08X, got %08X", crc, header.CRC)
		}
	}

	img := NewImg(header)
	_, err = file.Seek(int64(binary.Size(header)), io.SeekStart)
	if err != nil {
		return nil, err
	}
	if err := img.FillItems(file); err != nil {
		return nil, err
	}

	return img, nil
}
