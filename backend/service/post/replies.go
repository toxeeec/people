package post

import (
	people "github.com/toxeeec/people/backend"
)

const (
	queryCreateReply      = "INSERT INTO post(user_id, content, replies_to) VALUES($1, $2, $3) RETURNING post_id, content, created_at, replies_to, replies"
	queryIncrementReplies = "UPDATE post SET replies = replies + 1 WHERE post_id = $1"
)

const (
	repliesBase = selectPostAndAuthor + " WHERE replies_to = $2"
)

var repliesQueries = people.PaginationQueries(repliesBase, end, before, after, beforeAfter)

func (s *service) CreateReply(postID, userID uint, post people.PostBody) (people.Post, error) {
	var p people.Post
	tx, err := s.db.Beginx()
	if err != nil {
		return people.Post{}, err
	}
	defer tx.Rollback()

	err = tx.Get(&p, queryCreateReply, userID, post.Content, postID)
	if err != nil {
		return people.Post{}, err
	}

	_, err = tx.Exec(queryIncrementReplies, postID)
	if err != nil {
		return people.Post{}, err
	}

	return p, tx.Commit()
}

func (s *service) Replies(postID uint, userID *uint, p people.IDPagination) (people.Posts, error) {
	if userID == nil {
		userID = new(uint)
	}
	return people.PaginationSelect[people.Post](s.db, &repliesQueries, p, userID, postID)
}
