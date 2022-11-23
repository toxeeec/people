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
	selectIDByHandle = "SELECT user_id FROM user_profile WHERE handle = "
	isFollowing      = " EXISTS(SELECT 1 FROM follower WHERE follower_id = user_profile.user_id AND user_id = $1) as is_following"
)

const (
	queryFollow             = "INSERT INTO follower(user_id, follower_id) SELECT user_id, $1 FROM user_profile WHERE handle = $2"
	queryUnfollow           = "DELETE FROM follower WHERE follower_id = $1 AND user_id = (" + selectIDByHandle + "$2)"
	queryIncrementFollowers = "UPDATE user_profile SET followers = followers + 1 WHERE handle = $2 RETURNING following, followers," + isFollowing
	queryDecrementFollowers = "UPDATE user_profile SET followers = followers -1 WHERE handle = $2 RETURNING following, followers," + isFollowing
	queryIncrementFollowing = "UPDATE user_profile SET following = following + 1 WHERE user_id = $1"
	queryDecrementFollowing = "UPDATE user_profile SET following = following -1 WHERE user_id = $1"
	queryIsFollowing        = "SELECT EXISTS(SELECT 1 FROM follower WHERE follower_id = $1 AND user_id = (" + selectIDByHandle + "$2))"
	queryIsFollowed         = "SELECT EXISTS(SELECT 1 FROM follower WHERE user_id = $1 AND follower_id = (" + selectIDByHandle + "$2))"
)

const (
	followingBase = selectUser + " JOIN follower on user_profile.user_id = follower.user_id WHERE follower_id = $2"
	followersBase = selectUser + " JOIN follower on user_profile.user_id = follower_id WHERE follower.user_id = $2"
)

const (
	end                  = " ORDER BY followed_at DESC LIMIT $3"
	followingBefore      = " AND followed_at < (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + "$4) AND follower_id = $2)"
	followersBefore      = " AND followed_at < (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + "$4) AND user_id = $2)"
	followingAfter       = " AND followed_at > (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + "$4) AND follower_id = $2)"
	followersAfter       = " AND followed_at > (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + "$4) AND user_id = $2)"
	followingBeforeAfter = " AND followed_at < (SELECT followed_at FROM follower WHERE user_id = (" + selectIDByHandle + `$4) AND follower_id = $2) 
AND followed_at > (SELECT followed_at FROM follower WHERE user_id = (` + selectIDByHandle + "$5) AND follower_id = $2)"
	followersBeforeAfter = " AND followed_at < (SELECT followed_at FROM follower WHERE follower_id = (" + selectIDByHandle + `$4) AND user_id = $2) 
AND followed_at > (SELECT followed_at FROM follower WHERE follower_id = (` + selectIDByHandle + "$5) AND user_id = $2)"
)

var followingQueries = people.PaginationQueries(followingBase, end, followingBefore, followingAfter, followingBeforeAfter)
var followersQueries = people.PaginationQueries(followersBase, end, followersBefore, followersAfter, followersBeforeAfter)

func (s *service) Follow(id uint, handle string) (people.Follows, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return people.Follows{}, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(queryFollow, id, handle)
	if err != nil {
		var e *pq.Error
		if errors.As(err, &e) {
			if e.Constraint == "different_user" {
				return people.Follows{}, ErrSameUser
			}
			if e.Constraint == "follower_pkey" {
				return people.Follows{}, ErrAlreadyFollowed
			}
		}
		return people.Follows{}, err
	}

	var follows people.Follows
	err = tx.Get(&follows, queryIncrementFollowers, id, handle)
	if err != nil {
		return people.Follows{}, ErrInvalidHandle
	}
	follows.IsFollowed = true

	_, err = tx.Exec(queryIncrementFollowing, id)
	if err != nil {
		return people.Follows{}, err
	}

	return follows, tx.Commit()
}

func (s *service) Unfollow(id uint, handle string) (people.Follows, error) {
	tx, err := s.db.Beginx()
	if err != nil {
		return people.Follows{}, err
	}
	defer tx.Rollback()

	_, err = tx.Exec(queryUnfollow, id, handle)
	if err != nil {
		return people.Follows{}, err
	}

	var follows people.Follows
	err = tx.Get(&follows, queryDecrementFollowers, id, handle)
	if err != nil {
		return people.Follows{}, err
	}
	follows.IsFollowed = false

	_, err = tx.Exec(queryDecrementFollowing, id)
	if err != nil {
		return people.Follows{}, err
	}

	return follows, tx.Commit()
}

func (s *service) IsFollowing(id uint, handle string) (bool, error) {
	var isFollowing bool
	return isFollowing, s.db.Get(&isFollowing, queryIsFollowing, id, handle)
}

func (s *service) IsFollowed(id uint, handle string) (bool, error) {
	var isFollowed bool
	return isFollowed, s.db.Get(&isFollowed, queryIsFollowed, id, handle)
}

func (s *service) Following(id uint, userID *uint, p people.HandlePagination) (people.Users, error) {
	if userID == nil {
		userID = new(uint)
	}
	return people.PaginationSelect[people.User](s.db, &followingQueries, p, userID, id)
}

func (s *service) Followers(id uint, userID *uint, p people.HandlePagination) (people.Users, error) {
	if userID == nil {
		userID = new(uint)
	}
	return people.PaginationSelect[people.User](s.db, &followersQueries, p, userID, id)
}
