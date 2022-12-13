package repository

import (
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

type Follow interface {
	GetStatusFollowing(srcID, userID uint) bool
	GetStatusFollowed(srcID, userID uint) bool
	Create(targetID, userID uint) error
	Delete(targetID, userID uint) error
	ListStatusFollowing(srcIDs []uint, userID uint) (map[uint]struct{}, error)
	ListStatusFollowed(srcIDs []uint, userID uint) (map[uint]struct{}, error)
	ListFollowing(id uint, p pagination.ID) ([]people.User, error)
	ListFollowers(id uint, p pagination.ID) ([]people.User, error)
}

var (
	ErrSameUser        = errors.New("Current user and user being followed cannot be the same")
	ErrAlreadyFollowed = errors.New("This user is already followed by the current user")
)
