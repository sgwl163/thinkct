package user

import (
	"hotgo/api/api/user"
)

type ControllerV1 struct{}

func NewV1() user.IUserV1 {
	return &ControllerV1{}
}
