package user

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"golang.org/x/sync/errgroup"
)

type HandlePaginationParams struct {
	Limit  *uint
	Before *string
	After  *string
}

type Service interface {
	GetUserWithStatus(ctx context.Context, srcID, userID uint, auth bool) (people.User, error)
	GetUser(ctx context.Context, handle string, userID uint, auth bool) (people.User, error)
	Follow(ctx context.Context, handle string, userID uint) (people.User, error)
	Unfollow(ctx context.Context, handle string, userID uint) (people.User, error)
	ListFollowing(ctx context.Context, handle string, userID uint, auth bool, params HandlePaginationParams) (people.Users, error)
	ListFollowers(ctx context.Context, handle string, userID uint, auth bool, params HandlePaginationParams) (people.Users, error)
	ListCurrUserFollowing(ctx context.Context, userID uint, params HandlePaginationParams) (people.Users, error)
	ListCurrUserFollowers(ctx context.Context, userID uint, params HandlePaginationParams) (people.Users, error)
	ListPostLikes(ctx context.Context, postID, userID uint, auth bool, params HandlePaginationParams) (people.Users, error)
	ListStatus(ctx context.Context, srcIDs []uint, userID uint) (map[uint]people.FollowStatus, error)
	ListUsersWithStatus(ctx context.Context, srcIDs []uint, userID uint, auth bool) ([]people.User, error)
}

type userService struct {
	ur repository.User
	fr repository.Follow
	lr repository.Like
}

func NewService(ur repository.User, fr repository.Follow, lr repository.Like) Service {
	s := userService{}
	s.ur = ur
	s.fr = fr
	s.lr = lr
	return &s
}

func (s *userService) GetUser(ctx context.Context, handle string, userID uint, auth bool) (people.User, error) {
	id, err := s.ur.GetID(handle)
	if err != nil {
		return people.User{}, service.NewError(people.NotFoundError, "User not found")
	}
	return s.GetUserWithStatus(ctx, id, userID, auth)
}

func (s *userService) Follow(ctx context.Context, handle string, userID uint) (people.User, error) {
	id, err := s.ur.GetID(handle)
	if err != nil {
		return people.User{}, service.NewError(people.NotFoundError, "User not found")
	}
	err = s.fr.Create(id, userID)
	if err != nil {
		if errors.Is(err, repository.ErrAlreadyFollowed) {
			return people.User{}, service.NewError(people.ConflictError, errors.Unwrap(err).Error())
		}
		if errors.Is(err, repository.ErrSameUser) {
			return people.User{}, service.NewError(people.NotFoundError, errors.Unwrap(err).Error())
		}
		return people.User{}, err
	}
	u, err := s.GetUserWithStatus(context.Background(), id, userID, true)
	if err != nil {
		return people.User{}, err
	}
	return u, nil
}

func (s *userService) Unfollow(ctx context.Context, handle string, userID uint) (people.User, error) {
	id, err := s.ur.GetID(handle)
	if err != nil {
		return people.User{}, service.NewError(people.NotFoundError, "User not found")
	}
	err = s.fr.Delete(id, userID)
	if err != nil {
		return people.User{}, service.NewError(people.NotFoundError, "User not found")
	}
	u, err := s.GetUserWithStatus(context.Background(), id, userID, true)
	if err != nil {
		return people.User{}, err
	}
	return u, nil
}

func (s *userService) ListFollowing(ctx context.Context, handle string, userID uint, auth bool, params HandlePaginationParams) (people.Users, error) {
	hp := pagination.New(params.Before, params.After, params.Limit)
	var p pagination.ID
	var id uint
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		p, err = pagination.Handle(hp).IDPagination(ctx, s.ur.GetID)
		if err != nil {
			return service.NewError(people.NotFoundError, "User not found")
		}
		return nil
	})
	g.Go(func() error {
		var err error
		id, err = s.ur.GetID(handle)
		if err != nil {
			return service.NewError(people.NotFoundError, "User not found")
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return people.Users{}, err
	}

	us, err := s.fr.ListFollowing(id, p)
	if err != nil {
		return people.Users{}, err
	}
	if auth {
		fss, err := s.ListStatus(context.Background(), Slice(us).IDs(), userID)
		if err != nil {
			return people.Users{}, err
		}
		Slice(us).AddStatus(fss)
	}
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) ListFollowers(ctx context.Context, handle string, userID uint, auth bool, params HandlePaginationParams) (people.Users, error) {
	hp := pagination.New(params.Before, params.After, params.Limit)
	var p pagination.ID
	var id uint
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		p, err = pagination.Handle(hp).IDPagination(ctx, s.ur.GetID)
		if err != nil {
			return service.NewError(people.NotFoundError, "User not found")
		}
		return nil
	})
	g.Go(func() error {
		var err error
		id, err = s.ur.GetID(handle)
		if err != nil {
			return service.NewError(people.NotFoundError, "User not found")
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return people.Users{}, err
	}

	us, err := s.fr.ListFollowers(id, p)
	if err != nil {
		return people.Users{}, err
	}
	if auth {
		fss, err := s.ListStatus(context.Background(), Slice(us).IDs(), userID)
		if err != nil {
			return people.Users{}, err
		}
		Slice(us).AddStatus(fss)
	}
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) ListCurrUserFollowing(ctx context.Context, userID uint, params HandlePaginationParams) (people.Users, error) {
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.Users{}, err
	}
	return s.ListFollowing(ctx, u.Handle, userID, true, params)
}

func (s *userService) ListCurrUserFollowers(ctx context.Context, userID uint, params HandlePaginationParams) (people.Users, error) {
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.Users{}, err
	}
	return s.ListFollowers(ctx, u.Handle, userID, true, params)
}

func (s *userService) ListPostLikes(ctx context.Context, postID, userID uint, auth bool, params HandlePaginationParams) (people.Users, error) {
	hp := pagination.New(params.Before, params.After, params.Limit)
	p, err := pagination.Handle(hp).IDPagination(ctx, s.ur.GetID)
	if err != nil {
		return people.Users{}, service.NewError(people.NotFoundError, "User not found")
	}

	us, err := s.lr.ListUsers(postID, p)
	if err != nil {
		return people.Users{}, err
	}
	if auth {
		fss, err := s.ListStatus(context.Background(), Slice(us.Data).IDs(), userID)
		if err != nil {
			return people.Users{}, err
		}
		Slice(us.Data).AddStatus(fss)
	}
	return us, nil
}
