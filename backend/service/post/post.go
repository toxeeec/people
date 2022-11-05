package post

import (
	"github.com/jmoiron/sqlx"
	people "github.com/toxeeec/people/backend"
)

type service struct {
	db *sqlx.DB
}

func NewService(db *sqlx.DB) people.PostService {
	return &service{db}
}

const (
	selectPostAndUser = `SELECT post_id, content, created_at, replies_to, replies, likes, 
user_profile.handle AS "user.handle", user_profile.followers AS "user.followers", user_profile.following AS "user.following" 
FROM post JOIN user_profile ON post.user_id = user_profile.user_id`
)

const (
	queryCreate           = "INSERT INTO post(user_id, content) VALUES($1, $2) RETURNING post_id, content, created_at"
	queryGet              = selectPostAndUser + " WHERE post_id = $1"
	queryExists           = "SELECT EXISTS(SELECT 1 FROM post WHERE post_id = $1)"
	queryDelete           = `DELETE FROM post WHERE post_id = $1 AND user_id = $2 RETURNING replies_to`
	queryDecrementReplies = "UPDATE post SET replies = replies - 1 WHERE post_id = $1"
)

const (
	feedBase     = selectPostAndUser + " WHERE post.user_id IN (SELECT user_id FROM follower WHERE follower_id = $1) "
	fromUserBase = `SELECT post_id, content, created_at FROM post WHERE user_id = (SELECT user_id FROM user_profile WHERE handle = $1) AND replies_to IS NULL `
)

const (
	end         = " ORDER BY post_id DESC LIMIT $2"
	before      = " AND post_id < $3"
	after       = " AND post_id > $3"
	beforeAfter = " AND post_id < $3 AND post_id > $4"
)

var feedQueries = people.PaginationQueries(feedBase, end, before, after, beforeAfter)
var fromUserQueries = people.PaginationQueries(fromUserBase, end, before, after, beforeAfter)

func (s *service) Create(userID uint, post people.PostBody) (people.Post, error) {
	var p people.Post
	return p, s.db.Get(&p, queryCreate, userID, post.Content)
}

func (s *service) Get(id uint) (people.Post, error) {
	var p people.Post
	return p, s.db.Get(&p, queryGet, id)
}

func (s *service) Delete(postID, userID uint) error {
	var p people.Post
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	err = s.db.Get(&p, queryDelete, postID, userID)
	if err != nil {
		return err
	}

	if p.RepliesTo != nil && p.RepliesTo.Valid {
		_, err = s.db.Exec(queryDecrementReplies, postID)
		if err != nil {
			return err
		}
	}

	return tx.Commit()
}

func (s *service) FromUser(handle string, p people.IDPagination) (people.Posts, error) {
	return people.PaginationSelect[people.Post](s.db, &fromUserQueries, p, handle)
}

func (s *service) Feed(userID uint, p people.IDPagination) (people.Posts, error) {
	return people.PaginationSelect[people.Post](s.db, &feedQueries, p, userID)
}

func (s *service) Exists(postID uint) bool {
	var exists bool
	s.db.Get(&exists, queryExists, postID)
	return exists
}
