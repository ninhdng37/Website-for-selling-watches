package entity

import "strings"

type TypeOfWatch struct {
	TypeOfWatchID   *uint32 `json:"typeOfWatchID"`
	TypeOfWatchName string  `json:"typeOfWatchName"`
}

func (t *TypeOfWatch) TrimSpace() {
	t.TypeOfWatchName = strings.TrimSpace(t.TypeOfWatchName)
}
