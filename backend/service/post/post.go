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
	feedBase = `SELECT post_id, content, created_at, 
user_profile.handle AS "user.handle", user_profile.followers AS "user.followers", user_profile.following AS "user.following" 
FROM post JOIN user_profile ON post.user_id = user_profile.user_id WHERE post.user_id IN (SELECT user_id FROM follower WHERE follower_id = $1)`
	feedEnd = " ORDER BY post_id DESC LIMIT $2"
)

const (
	queryCreate = "INSERT INTO post(user_id, content) VALUES($1, $2) RETURNING post_id, content, created_at"
	queryGet    = `SELECT post_id, content, created_at, replies_to, replies, user_profile.handle AS "user.handle", 
user_profile.followers AS "user.followers", user_profile.following AS "user.following"
FROM post JOIN user_profile ON post.user_id = user_profile.user_id WHERE post_id = $1`
	queryDelete           = `DELETE FROM post WHERE post_id = $1 AND user_id = $2 RETURNING replies_to`
	queryDecrementReplies = "UPDATE post SET replies = replies - 1 WHERE post_id = $1"
	queryFromUser         = `SELECT post_id, content, created_at FROM post 
WHERE user_id = (SELECT user_id FROM user_profile WHERE handle = $1) AND replies_to IS NULL ORDER BY post_id DESC OFFSET $2 LIMIT $3`
	queryFeedNone        = feedBase + feedEnd
	queryFeedBefore      = feedBase + " AND post_id < $3" + feedEnd
	queryFeedAfter       = feedBase + " AND post_id > $3" + feedEnd
	queryFeedBeforeAfter = feedBase + " AND post_id < $3 AND post_id > $4" + feedEnd
	queryExists          = "SELECT EXISTS(SELECT 1 FROM post WHERE post_id = $1)"
)

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

func (s *service) FromUser(handle string, p people.Pagination) (people.Posts, error) {
	posts := make([]people.Post, 0, p.Limit)
	return posts, s.db.Select(&posts, queryFromUser, handle, p.Offset, p.Limit)
}

func (s *service) Feed(userID uint, p people.SeekPagination) (people.FeedResponse, error) {
	res := people.FeedResponse{}
	res.Data = make([]people.Post, 0, p.Limit)
	var err error
	switch p.Mode {
	case people.PaginationModeNone:
		err = s.db.Select(&res.Data, queryFeedNone, userID, p.Limit)
	case people.PaginationModeBefore:
		err = s.db.Select(&res.Data, queryFeedBefore, userID, p.Limit, p.Before)
	case people.PaginationModeAfter:
		err = s.db.Select(&res.Data, queryFeedAfter, userID, p.Limit, p.After)
	case people.PaginationModeBeforeAfter:
		err = s.db.Select(&res.Data, queryFeedBeforeAfter, userID, p.Limit, p.Before, p.After)
	}
	if err != nil {
		return people.FeedResponse{}, err
	}

	if len(res.Data) > 0 {
		res.Meta = &people.SeekPaginationMeta{
			NewestID: res.Data[0].ID,
			OldestID: res.Data[len(res.Data)-1].ID,
		}
	}
	return res, nil
}

func (s *service) Exists(postID uint) bool {
	var exists bool
	s.db.Get(&exists, queryExists, postID)
	return exists
}
