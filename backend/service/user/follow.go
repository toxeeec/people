package user

import (
	"errors"

	"github.com/lib/pq"
	people "github.com/toxeeec/people/backend"
)

var (
	ErrSameUser        = errors.New("Current user and user being followed cannot be the same")
	ErrInvalidHandle   = errors.New("Invalid handle")
	ErrAlreadyFollowed = errors.New("This user is already followed by the current user")
)

const (
	selectUser       = "SELECT handle, followers, following FROM user_profile"
	selectIDByHandle = "SELECT user_id FROM user_profile WHERE handle = "
)

const (
	queryFollow             = "INSERT INTO follower(user_id, follower_id) SELECT user_id, $1 FROM user_profile WHERE handle = $2 RETURNING user_id"
	queryUnfollow           = "DELETE FROM follower WHERE follower_id = $1 AND user_id = (" + selectIDByHandle + "$2) RETURNING user_id"
	queryIncrementFollowers = "UPDATE user_profile SET followers = followers + 1 WHERE user_id = $1"
	queryDecrementFollowers = "UPDATE user_profile SET followers = followers -1 WHERE user_id = $1"
	queryIncrementFollowing = "UPDATE user_profile SET following = following + 1 WHERE user_id = $1"
	queryDecrementFollowing = "UPDATE user_profile SET following = following -1 WHERE user_id = $1"
	queryIsFollowing        = "SELECT EXISTS(SELECT 1 FROM follower WHERE follower_id = $1 AND user_id = (" + selectIDByHandle + "$2))"
	queryIsFollowed         = "SELECT EXISTS(SELECT 1 FROM follower WHERE user_id = $1 AND follower_id = (" + selectIDByHandle + "$2))"
	queryFollowing          = "SELECT handle, followers, following FROM user_profile JOIN follower on user_profile.user_id = follower.user_id WHERE follower_id = $1 ORDER BY followed_at DESC OFFSET $2 LIMIT $3"
	queryFollowers          = "SELECT handle, followers, following FROM user_profile JOIN follower on user_profile.user_id = follower.follower_id WHERE follower.user_id = $1 ORDER BY followed_at DESC OFFSET $2 LIMIT $3"
)

const (
	followingBase = selectUser + " JOIN follower on user_profile.user_id = follower.user_id WHERE follower_id = $1"
	followersBase = selectUser + " JOIN follower on user_profile.user_id = follower_id WHERE follower.user_id = $1"
)

const (
	end                  = " ORDER BY followed_at DESC LIMIT $2"
	followingBefore      = " AND followed_at < (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + "$3))"
	followersBefore      = " AND followed_at < (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + "$3))"
	followingAfter       = " AND followed_at > (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + "$3))"
	followersAfter       = " AND followed_at > (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + "$3))"
	followingBeforeAfter = " AND followed_at < (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + `$3)) 
AND followed_at > (SELECT followed_at FROM follower WHERE user_id = (` + selectIDByHandle + "$4))"
	followersBeforeAfter = " AND followed_at < (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + `$3)) 
AND followed_at > (SELECT followed_at FROM follower WHERE follower_id = (` + selectIDByHandle + "$4))"
)

var followingQueries = people.PaginationQueries(followingBase, end, followingBefore, followingAfter, followingBeforeAfter)
var followersQueries = people.PaginationQueries(followersBase, end, followersBefore, followersAfter, followersBeforeAfter)

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
		return err
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

func (s *service) Following(id uint, p people.HandlePagination) (people.Users, error) {
	return people.PaginationSelect[people.User](s.db, &followingQueries, p, id)
}

func (s *service) Followers(id uint, p people.HandlePagination) (people.Users, error) {
	return people.PaginationSelect[people.User](s.db, &followersQueries, p, id)
}
