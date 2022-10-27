package people

import "strings"

func (p *PostBody) TrimContent() {
	p.Content = strings.TrimSpace(p.Content)
}

type PostService interface {
	Create(userID uint, p PostBody) (Post, error)
	Get(postID uint) (Post, error)
	Delete(postID, userID uint) error
	FromUser(handle string, p Pagination) (Posts, error)
	Feed(userID uint, p SeekPagination) (FeedResponse, error)
	Exists(postID uint) bool
	CreateReply(postID, userID uint, p PostBody) (Post, error)
}
