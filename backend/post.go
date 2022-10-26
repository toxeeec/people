package people

import "strings"

func (p *PostBody) TrimContent() {
	p.Content = strings.TrimSpace(p.Content)
}

type PostService interface {
	Create(userID uint, p PostBody) (Post, error)
}
