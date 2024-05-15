package amlimg

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	Magic = uint32(0x27B51956)
)

type Header struct {
	CRC       uint32
	Version   uint32
	Magic     uint32
	Size      uint64
	AlignSize uint32
	ItemCount uint32
	Reserved  [36]byte
}

func (header *Header) FillFrom(reader io.Reader) error {
	err := binary.Read(reader, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	if header.Magic != Magic {
		return fmt.Errorf("invalid magic: %x", header.Magic)
	}

	return nil
}

func ReadHeader(reader io.Reader) (*Header, error) {
	header := Header{}
	return &header, header.FillFrom(reader)
}
