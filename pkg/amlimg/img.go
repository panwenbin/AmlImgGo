package amlimg

import (
	"encoding/binary"
	"fmt"
	"io"
)

const (
	Magic = uint32(0x27B51956)
)

type ImgHeader struct {
	CRC       uint32
	Version   uint32
	Magic     uint32
	Size      uint64
	AlignSize uint32
	ItemCount uint32
	Reserved  [36]byte
}

func (header *ImgHeader) FillFrom(reader io.Reader) error {
	err := binary.Read(reader, binary.LittleEndian, header)
	if err != nil {
		return err
	}

	if header.Magic != Magic {
		return fmt.Errorf("invalid magic: %x", header.Magic)
	}

	return nil
}

func ReadHeader(reader io.Reader) (*ImgHeader, error) {
	header := ImgHeader{}
	return &header, header.FillFrom(reader)
}

type Img struct {
	*ImgHeader
	Items []Item
}

func NewImg(header *ImgHeader) *Img {
	return &Img{
		ImgHeader: header,
	}
}

func (img *Img) FillItems(reader io.Reader) error {
	img.Items = make([]Item, 0, img.ItemCount)
	for i := 0; i < int(img.ItemCount); i++ {
		if img.Version == 1 {
			item := &ItemV1{}
			if err := item.FillFrom(reader); err != nil {
				return err
			}
			img.Items = append(img.Items, item)
		}
		if img.Version == 2 {
			item := &ItemV2{}
			if err := item.FillFrom(reader); err != nil {
				return err
			}
			img.Items = append(img.Items, item)
		}
	}

	return nil
}
