package post

import (
	people "github.com/toxeeec/people/backend"
)

const (
	queryCreateReply      = "INSERT INTO post(user_id, content, replies_to) VALUES($1, $2, $3) RETURNING post_id, content, created_at, replies_to, replies"
	queryIncrementReplies = "UPDATE post SET replies = replies + 1 WHERE post_id = $1"
)

const (
	repliesBase = `SELECT post_id, content, created_at, replies_to, replies, user_profile.handle AS "user.handle", 
user_profile.followers AS "user.followers", user_profile.following AS "user.following"
FROM post JOIN user_profile ON post.user_id = user_profile.user_id WHERE replies_to = $1`
)

var repliesQueries = people.SeekPaginationQueries(repliesBase, end, before, after, beforeAfter)

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

func (s *service) Replies(postID uint, p people.SeekPagination) (people.Posts, error) {
	return people.SeekPaginationSelect[people.Post](s.db, &repliesQueries, p, postID)
}
