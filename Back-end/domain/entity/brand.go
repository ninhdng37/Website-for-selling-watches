package entity

import "strings"

type Brand struct {
	BrandID   *uint32 `json:"brandID"`
	BrandName string  `json:"brandName"`
}

func (b *Brand) TrimSpace() {
	b.BrandName = strings.TrimSpace(b.BrandName)
}
