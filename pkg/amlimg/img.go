package amlimg

import (
	"io"
)

type Img struct {
	*Header
	Items []Item
}

func NewImg(header *Header) *Img {
	return &Img{
		Header: header,
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
