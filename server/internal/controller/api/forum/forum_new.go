package forum

import "hotgo/api/api/forum"

type cV1 struct{}
type cPublicV1 struct{}

func NewV1() forum.IForumV1 {
	return &cV1{}
}

func NewPublicV1() forum.IForumPublicV1 {
	return &cPublicV1{}
}
