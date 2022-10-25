package user

import (
	"errors"

	"github.com/lib/pq"
)

var (
	ErrSameUser        = errors.New("Current user and user being followed cannot be the same")
	ErrInvalidHandle   = errors.New("Invalid handle")
	ErrAlreadyFollowed = errors.New("This user is already followed by the current user")
	ErrNotFollowed     = errors.New("This user is not followed by the current user")
)

const (
	queryFollow             = "INSERT INTO follower(user_id, follower_id) SELECT user_id, $1 FROM user_profile WHERE handle = $2 RETURNING user_id"
	queryUnfollow           = "DELETE FROM follower WHERE follower_id = $1 AND user_id = (SELECT user_id FROM user_profile WHERE handle = $2) RETURNING user_id"
	queryIncrementFollowers = "UPDATE user_profile SET followers = followers + 1 WHERE user_id = $1"
	queryIncrementFollowing = "UPDATE user_profile SET following = following + 1 WHERE user_id = $1"
	queryDecrementFollowers = "UPDATE user_profile SET followers = followers -1 WHERE user_id = $1"
	queryDecrementFollowing = "UPDATE user_profile SET following = following -1 WHERE user_id = $1"
	queryIsFollowing        = "SELECT EXISTS(SELECT 1 FROM follower WHERE follower_id = $1 AND user_id = (SELECT user_id FROM user_profile WHERE handle = $2))"
	queryIsFollowed         = "SELECT EXISTS(SELECT 1 FROM follower WHERE user_id = $1 AND follower_id = (SELECT user_id FROM user_profile WHERE handle = $2))"
)

func (s *service) Follow(id uint, handle string) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var userID uint
	err = tx.Get(&userID, queryFollow, id, handle)
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Constraint == "different_user" {
				return ErrSameUser
			}
			if e.Constraint == "follower_pkey" {
				return ErrAlreadyFollowed
			}
			return err
		}
		return ErrInvalidHandle
	}

	_, err = tx.Exec(queryIncrementFollowers, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryIncrementFollowing, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *service) Unfollow(id uint, handle string) error {
	tx, err := s.db.Beginx()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	var userID uint
	err = tx.Get(&userID, queryUnfollow, id, handle)
	if err != nil {
		return ErrNotFollowed
	}

	_, err = tx.Exec(queryDecrementFollowers, userID)
	if err != nil {
		return err
	}

	_, err = tx.Exec(queryDecrementFollowing, id)
	if err != nil {
		return err
	}

	return tx.Commit()
}

func (s *service) IsFollowing(id uint, handle string) (bool, error) {
	var isFollowing bool
	return isFollowing, s.db.Get(&isFollowing, queryIsFollowing, id, handle)
}

func (s *service) IsFollowed(id uint, handle string) (bool, error) {
	var isFollowed bool
	return isFollowed, s.db.Get(&isFollowed, queryIsFollowed, id, handle)
}
