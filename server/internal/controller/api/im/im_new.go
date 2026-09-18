package im

import (
	"hotgo/api/api/im"
)

type ControllerV1 struct{}

func NewV1() im.IImV1 {
	return &ControllerV1{}
}