package user

import (
	"context"
	"errors"

	"github.com/go-playground/validator/v10"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	// TODO: GetUserFromContext
	GetUserWithStatus(ctx context.Context, srcID, userID uint, auth bool) (people.User, error)
	GetUser(ctx context.Context, handle string, userID uint, auth bool) (people.User, error)
	Follow(ctx context.Context, handle string, userID uint) (people.User, error)
	Unfollow(ctx context.Context, handle string, userID uint) (people.User, error)
	ListFollowing(ctx context.Context, handle string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error)
	ListFollowers(ctx context.Context, handle string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error)
	ListCurrUserFollowing(ctx context.Context, userID uint, params pagination.HandleParams) (people.Users, error)
	ListCurrUserFollowers(ctx context.Context, userID uint, params pagination.HandleParams) (people.Users, error)
	ListPostLikes(ctx context.Context, postID, userID uint, auth bool, params pagination.HandleParams) (people.Users, error)
	ListStatus(ctx context.Context, userIDs []uint, srcID uint) (map[uint]people.FollowStatus, error)
	ListUsersWithStatus(ctx context.Context, userIDs []uint, srcID uint, auth bool) ([]people.User, error)
	ListMatches(ctx context.Context, query string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error)
	Delete(userID uint) error
	Update(userID uint, handle string) (people.User, error)
	Validate(u people.AuthUser) error
}

type userService struct {
	v  *validator.Validate
	ur repository.User
	fr repository.Follow
	lr repository.Like
}

func NewService(v *validator.Validate, ur repository.User, fr repository.Follow, lr repository.Like) Service {
	return &userService{
		v, ur, fr, lr,
	}
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

func (s *userService) ListFollowing(ctx context.Context, handle string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error) {
	hp := pagination.New(params)
	var p pagination.ID
	var id uint
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		p, err = pagination.IntoID(ctx, hp, s.ur.GetID)
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

	ids, err := s.fr.ListFollowing(id, &p)
	if err != nil {
		return people.Users{}, err
	}
	us, err := s.ListUsersWithStatus(context.Background(), ids, userID, auth)
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) ListFollowers(ctx context.Context, handle string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error) {
	hp := pagination.New(params)
	var p pagination.ID
	var id uint
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		p, err = pagination.IntoID(ctx, hp, s.ur.GetID)
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

	ids, err := s.fr.ListFollowers(id, &p)
	if err != nil {
		return people.Users{}, err
	}
	us, err := s.ListUsersWithStatus(context.Background(), ids, userID, auth)
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) ListCurrUserFollowing(ctx context.Context, userID uint, params pagination.HandleParams) (people.Users, error) {
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.Users{}, err
	}
	return s.ListFollowing(ctx, u.Handle, userID, true, params)
}

func (s *userService) ListCurrUserFollowers(ctx context.Context, userID uint, params pagination.HandleParams) (people.Users, error) {
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.Users{}, err
	}
	return s.ListFollowers(ctx, u.Handle, userID, true, params)
}

func (s *userService) ListPostLikes(ctx context.Context, postID, userID uint, auth bool, params pagination.HandleParams) (people.Users, error) {
	hp := pagination.New(params)
	p, err := pagination.IntoID(ctx, hp, s.ur.GetID)
	if err != nil {
		return people.Users{}, service.NewError(people.NotFoundError, "User not found")
	}
	ids, err := s.lr.ListPostLikes(postID, &p)
	if err != nil {
		return people.Users{}, err
	}
	us, err := s.ur.List(ids)
	if err != nil {
		return people.Users{}, err
	}

	if auth {
		fss, err := s.ListStatus(context.Background(), IDs(us), userID)
		if err != nil {
			return people.Users{}, err
		}
		AddStatuses(us, fss)
	}
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) ListMatches(ctx context.Context, query string, userID uint, auth bool, params pagination.HandleParams) (people.Users, error) {
	hp := pagination.New(params)
	p, err := pagination.IntoID(ctx, hp, s.ur.GetID)
	if err != nil {
		return people.Users{}, service.NewError(people.NotFoundError, "User not found")
	}
	us, err := s.ur.ListMatches(query, p)
	if err != nil {
		return people.Users{}, err
	}
	if auth {
		fss, err := s.ListStatus(context.Background(), IDs(us), userID)
		if err != nil {
			return people.Users{}, err
		}
		AddStatuses(us, fss)
	}
	return pagination.NewResults[people.User, string](us), nil
}

func (s *userService) Delete(userID uint) error {
	// TODO: mark images as unused
	ids, err := s.lr.ListUserLikes(userID, nil)
	if err != nil {
		return err
	}
	err = s.lr.DeleteLike(ids)
	if err != nil {
		return err
	}
	ids, err = s.fr.ListFollowing(userID, nil)
	err = s.fr.DeleteFollower(ids)
	if err != nil {
		return err
	}
	ids, err = s.fr.ListFollowers(userID, nil)
	err = s.fr.DeleteFollowing(ids)
	if err != nil {
		return err
	}
	return s.ur.Delete(userID)
}

func (s *userService) Update(userID uint, handle string) (people.User, error) {
	u, err := s.ur.Get(userID)
	if err != nil {
		return people.User{}, err
	}
	if handle != u.Handle {
		if err := s.Validate(people.AuthUser{Handle: handle}); err != nil {
			return people.User{}, err
		}
	}
	return s.ur.Update(userID, handle)
}

func (s *userService) Validate(u people.AuthUser) error {
	if err := s.v.Var(u.Handle, "alphanum"); err != nil {
		err := err.(validator.ValidationErrors)
		switch err[0].Tag() {
		case "alphanum":
			return service.NewError(people.ValidationError, "Handle cannot contain special characters")
		default:
			return errors.New("Unknown")
		}
	}
	if _, err := s.ur.GetID(u.Handle); err == nil {
		return service.NewError(people.ValidationError, "User with this handle already exists")
	}
	return nil
}
