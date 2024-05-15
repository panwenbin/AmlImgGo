package amlimg

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Item interface {
	FillFrom(reader io.Reader) error
	TypeString() string
	NameString() string
}

type ItemHeader struct {
	Id            uint32
	ImgType       uint32
	OffsetOfItem  uint64
	OffsetOfImage uint64
	Size          uint64
}

type ItemV1 struct {
	ItemHeader
	Type     [32]byte
	Name     [32]byte
	Reserved [32]byte
}

func (itemv1 *ItemV1) FillFrom(reader io.Reader) error {
	return binary.Read(reader, binary.LittleEndian, itemv1)
}

func (itemv1 ItemV1) TypeString() string {
	return string(bytes.TrimRight(itemv1.Type[:], "\x00"))
}

func (itemv1 ItemV1) NameString() string {
	return string(bytes.TrimRight(itemv1.Name[:], "\x00"))
}

type ItemV2 struct {
	ItemHeader
	Type     [256]byte
	Name     [256]byte
	Reserved [32]byte
}

func (itemv2 *ItemV2) FillFrom(reader io.Reader) error {
	return binary.Read(reader, binary.LittleEndian, itemv2)
}

func (itemv2 ItemV2) TypeString() string {
	return string(bytes.TrimRight(itemv2.Type[:], "\x00"))
}

func (itemv2 ItemV2) NameString() string {
	return string(bytes.TrimRight(itemv2.Name[:], "\x00"))
}
