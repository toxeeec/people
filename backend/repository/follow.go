package repository

import (
	"errors"

	"github.com/toxeeec/people/backend/pagination"
)

type Follow interface {
	GetStatusFollowing(srcID, userID uint) bool
	GetStatusFollowed(srcID, userID uint) bool
	Create(targetID, userID uint) error
	Delete(targetID, userID uint) error
	ListStatusFollowing(userIDs []uint, srcID uint) (map[uint]struct{}, error)
	ListStatusFollowed(userIDs []uint, srcID uint) (map[uint]struct{}, error)
	ListFollowing(id uint, p *pagination.ID) ([]uint, error)
	ListFollowers(id uint, p *pagination.ID) ([]uint, error)
	DeleteFollower(userIDs []uint) error
	DeleteFollowing(userIDs []uint) error
}

var (
	ErrSameUser        = errors.New("Current user and user being followed cannot be the same")
	ErrAlreadyFollowed = errors.New("This user is already followed by the current user")
)
