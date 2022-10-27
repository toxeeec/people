package post

import people "github.com/toxeeec/people/backend"

const (
	queryCreateReply      = "INSERT INTO post(user_id, content, replies_to) VALUES($1, $2, $3) RETURNING post_id, content, created_at, replies_to, replies"
	queryIncrementReplies = "UPDATE post SET replies = replies + 1 WHERE post_id = $1"
)

func (s *service) CreateReply(postID, userID uint, post people.PostBody) (people.Post, error) {
	var p people.Post
	tx, err := s.db.Beginx()
	if err != nil {
		return p, err
	}
	defer tx.Rollback()

	err = tx.Get(&p, queryCreateReply, userID, post.Content, postID)
	if err != nil {
		return p, err
	}

	_, err = tx.Exec(queryIncrementReplies, postID)
	if err != nil {
		return p, err
	}

	return p, tx.Commit()
}
