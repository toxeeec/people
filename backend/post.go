package people

import "strings"

func (p *PostBody) TrimContent() {
	p.Content = strings.TrimSpace(p.Content)
}

func (p Post) Identify() uint {
	return p.ID
}

type PostService interface {
	Create(userID uint, p PostBody) (Post, error)
	Get(postID uint) (Post, error)
	Delete(postID, userID uint) error
	FromUser(handle string, p SeekPagination) (Posts, error)
	Feed(userID uint, p SeekPagination) (Posts, error)
	Exists(postID uint) bool
	CreateReply(postID, userID uint, p PostBody) (Post, error)
	Replies(postID uint, p SeekPagination) (Posts, error)
	Like(postID, userID uint) (Likes, error)
	Unlike(postID, userID uint) (Likes, error)
}
