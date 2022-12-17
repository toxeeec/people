package post

import (
	"context"
	"errors"
	"strings"

	"github.com/go-playground/validator/v10"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/service/user"
	"golang.org/x/sync/errgroup"
)

func TrimContent(p people.NewPost) people.NewPost {
	p.Content = strings.TrimSpace(p.Content)
	return p
}

type IDPaginationParams struct {
	Limit  *uint
	Before *uint
	After  *uint
}

type Service interface {
	Create(ctx context.Context, np people.NewPost, userID uint, repliesTo *uint) (people.PostResponse, error)
	Get(ctx context.Context, postID, userID uint, auth bool) (people.PostResponse, error)
	Delete(postID, userID uint) error
	ListUserPosts(ctx context.Context, handle string, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error)
	ListFeed(ctx context.Context, userID uint, params IDPaginationParams) (people.PostsResponse, error)
	ListReplies(ctx context.Context, postID, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error)
	Like(postID, userID uint) (people.PostResponse, error)
	Unlike(postID, userID uint) (people.PostResponse, error)
	ListUserLikes(ctx context.Context, handle string, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error)
}

type postService struct {
	pr repository.Post
	ur repository.User
	fr repository.Follow
	lr repository.Like
	us user.Service
	v  *validator.Validate
}

func NewService(v *validator.Validate, pr repository.Post, ur repository.User, fr repository.Follow, lr repository.Like, us user.Service) Service {
	s := postService{}
	s.v = v
	s.pr = pr
	s.ur = ur
	s.fr = fr
	s.lr = lr
	s.us = us
	return &s
}

func (s *postService) validate(p people.NewPost) error {
	if err := s.v.Var(p.Content, "min=1"); err != nil {
		err := err.(validator.ValidationErrors)
		switch err[0].Tag() {
		case "min":
			return service.NewError(people.ValidationError, "Content cannot be empty")
		default:
			return errors.New("Unknown")
		}
	}
	return nil
}

func (s *postService) Create(ctx context.Context, np people.NewPost, userID uint, repliesTo *uint) (people.PostResponse, error) {
	np = TrimContent(np)
	if err := s.validate(np); err != nil {
		return people.PostResponse{}, err
	}
	var p people.Post
	var u people.User
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		p, err = s.pr.Create(np, userID, repliesTo)
		if err != nil {
			return service.NewError(people.NotFoundError, "Post not found")
		}
		return nil
	})
	g.Go(func() error {
		var err error
		u, err = s.us.GetUserWithStatus(context.Background(), userID, userID, true)
		return err
	})
	if err := g.Wait(); err != nil {
		return people.PostResponse{}, err
	}
	return people.PostResponse{Data: p, User: u}, nil
}

func (s *postService) Get(ctx context.Context, postID uint, userID uint, auth bool) (people.PostResponse, error) {
	return s.getPostResponseWithStatuses(ctx, postID, userID, auth)
}

func (s *postService) Delete(postID, userID uint) error {
	err := s.pr.Delete(postID, userID)
	if err != nil {
		return err
	}
	return nil
}

func (s *postService) ListUserPosts(ctx context.Context, handle string, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error) {
	p := pagination.New(params.Before, params.After, params.Limit)
	id, err := s.ur.GetID(handle)
	if err != nil {
		return people.PostsResponse{}, service.NewError(people.NotFoundError, "User not found")
	}
	var ps []people.Post
	var u people.User
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		ps, err = s.pr.ListUserPosts(id, p)
		return err
	})
	g.Go(func() error {
		var err error
		u, err = s.us.GetUserWithStatus(context.Background(), id, userID, auth)
		return err
	})
	if err := g.Wait(); err != nil {
		return people.PostsResponse{}, err
	}

	if auth {
		lss, err := s.listStatus(Slice(ps).IDs(), userID)
		if err != nil {
			return people.PostsResponse{}, nil
		}
		Slice(ps).AddStatus(lss)
	}
	prs := make([]people.PostResponse, len(ps))
	for i, p := range ps {
		prs[i] = people.PostResponse{Data: p, User: u}
	}
	return pagination.NewResults[people.PostResponse, uint](prs), nil
}

func (s *postService) ListFeed(ctx context.Context, userID uint, params IDPaginationParams) (people.PostsResponse, error) {
	p := pagination.New(params.Before, params.After, params.Limit)
	us, err := s.fr.ListFollowing(userID, p)
	if err != nil {
		return people.PostsResponse{}, err
	}
	ids := user.Slice(us).IDs()
	ids = append(ids, userID)
	var ps []people.Post
	var fss map[uint]people.FollowStatus
	var u people.User
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		ps, err = s.pr.ListFeed(ids, userID, p)
		return err
	})
	g.Go(func() error {
		var err error
		fss, err = s.us.ListStatus(context.Background(), ids, userID)
		return err
	})
	g.Go(func() error {
		var err error
		u, err = s.us.GetUserWithStatus(context.Background(), userID, 0, false)
		return err
	})
	if err := g.Wait(); err != nil {
		return people.PostsResponse{}, err
	}

	lss, err := s.listStatus(Slice(ps).IDs(), userID)
	if err != nil {
		return people.PostsResponse{}, nil
	}
	Slice(ps).AddStatus(lss)
	user.Slice(us).AddStatus(fss)
	um := user.Slice(append(us, u)).ToMap()
	prs := make([]people.PostResponse, len(ps))
	for i, p := range ps {
		prs[i] = people.PostResponse{Data: p, User: um[p.UserID]}
	}
	return pagination.NewResults[people.PostResponse, uint](prs), nil
}

func (s *postService) ListReplies(ctx context.Context, postID uint, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error) {
	p := pagination.New(params.Before, params.After, params.Limit)
	ps, err := s.pr.ListReplies(postID, p)
	if err != nil {
		return people.PostsResponse{}, service.NewError(people.NotFoundError, "Post not found")
	}
	ids := Slice(ps).UserIDs()
	us, err := s.us.ListUsersWithStatus(context.Background(), ids, userID, auth)
	if err != nil {
		return people.PostsResponse{}, err
	}
	if auth {
		lss, err := s.listStatus(Slice(ps).IDs(), userID)
		if err != nil {
			return people.PostsResponse{}, nil
		}
		Slice(ps).AddStatus(lss)
	}
	um := user.Slice(us).ToMap()
	prs := make([]people.PostResponse, len(ps))
	for i, p := range ps {
		prs[i] = people.PostResponse{Data: p, User: um[p.UserID]}
	}
	return pagination.NewResults[people.PostResponse, uint](prs), nil
}

func (s *postService) Like(postID uint, userID uint) (people.PostResponse, error) {
	err := s.lr.Create(postID, userID)
	if err != nil {
		if errors.Is(err, repository.ErrPostNotFound) {
			return people.PostResponse{}, service.NewError(people.NotFoundError, err.Error())
		}
		if errors.Is(err, repository.ErrAlreadyLiked) {
			return people.PostResponse{}, service.NewError(people.ConflictError, err.Error())
		}
		return people.PostResponse{}, err
	}
	pr, err := s.getPostResponseWithStatuses(context.Background(), postID, userID, true)
	if err != nil {
		return people.PostResponse{}, err
	}

	return pr, nil
}

func (s *postService) Unlike(postID uint, userID uint) (people.PostResponse, error) {
	err := s.lr.Delete(postID, userID)
	if err != nil {
		return people.PostResponse{}, service.NewError(people.NotFoundError, "Post not found")
	}
	pr, err := s.getPostResponseWithStatuses(context.Background(), postID, userID, true)
	if err != nil {
		return people.PostResponse{}, err
	}
	return pr, nil
}

func (s *postService) listStatus(ids []uint, userID uint) (map[uint]people.LikeStatus, error) {
	lss := make(map[uint]people.LikeStatus, len(ids))
	liked, err := s.lr.ListStatusLiked(ids, userID)
	if err != nil {
		return nil, err
	}
	for _, id := range ids {
		_, likedOk := liked[id]
		lss[id] = people.LikeStatus{IsLiked: likedOk}
	}
	return lss, nil
}

func (s *postService) ListUserLikes(ctx context.Context, handle string, userID uint, auth bool, params IDPaginationParams) (people.PostsResponse, error) {
	p := pagination.New(params.Before, params.After, params.Limit)
	id, err := s.ur.GetID(handle)
	if err != nil {
		return people.PostsResponse{}, service.NewError(people.NotFoundError, "User not found")
	}
	ps, err := s.lr.ListUserLikes(id, p)
	if err != nil {
		return people.PostsResponse{}, err
	}

	var us []people.User
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		us, err = s.us.ListUsersWithStatus(context.Background(), Slice(ps).UserIDs(), userID, auth)
		if err != nil {
			return err
		}
		return nil
	})
	if auth {
		g.Go(func() error {
			lss, err := s.listStatus(Slice(ps).IDs(), userID)
			if err != nil {
				return err
			}
			Slice(ps).AddStatus(lss)
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return people.PostsResponse{}, nil
	}
	um := user.Slice(us).ToMap()
	prs := make([]people.PostResponse, len(ps))
	for i, p := range ps {
		prs[i] = people.PostResponse{Data: p, User: um[p.UserID]}
	}
	return pagination.NewResults[people.PostResponse, uint](prs), nil
}

func (s *postService) getPostResponseWithStatuses(ctx context.Context, postID uint, userID uint, auth bool) (people.PostResponse, error) {
	pc := make(chan people.Post, 1)
	var u people.User
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		p, err := s.pr.Get(postID)
		if err != nil {
			return service.NewError(people.NotFoundError, "Post not found")
		}
		select {
		case pc <- p:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if auth {
		g.Go(func() error {
			ls := s.lr.Status(postID, userID)
			select {
			case p := <-pc:
				p.Status = &ls
				pc <- p
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return people.PostResponse{}, err
	}

	p := <-pc
	u, err := s.us.GetUserWithStatus(context.Background(), p.UserID, userID, auth)
	if err != nil {
		return people.PostResponse{}, err
	}
	return people.PostResponse{Data: p, User: u}, nil
}
