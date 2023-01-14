package post_test

import (
	"context"
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/service/image"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
)

type PostSuite struct {
	suite.Suite
	ps post.Service
	pr repository.Post
	ur repository.User
}

func (s *PostSuite) TestCreate() {
	emptyContent := people.NewPost{Content: "\t\n \n\t"}
	var np people.NewPost
	var au people.AuthUser
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)
	invalidID := p.ID + 5

	validationError := people.ValidationError
	notFoundError := people.NotFoundError

	tests := map[string]struct {
		np        people.NewPost
		repliesTo *uint
		valid     bool
		kind      *people.ErrorKind
	}{
		"empty content":       {emptyContent, nil, false, &validationError},
		"invalid parent post": {np, &invalidID, false, &notFoundError},
		"valid reply":         {np, &p.ID, true, nil},
		"valid":               {np, nil, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			pr, err := s.ps.Create(context.Background(), tc.np, u.ID, tc.repliesTo)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.np.Content, pr.Data.Content)
				assert.Equal(s.T(), u.ID, pr.User.ID)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *PostSuite) TestGet() {
	var np people.NewPost
	var au people.AuthUser
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p, _ := s.pr.Create(np, u.ID, nil)

	notFoundError := people.NotFoundError

	tests := map[string]struct {
		id    uint
		valid bool
		kind  *people.ErrorKind
	}{
		"not found": {p.ID + 5, false, &notFoundError},
		"valid":     {p.ID, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			pr, err := s.ps.Get(context.Background(), tc.id, 0, false)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.Equal(s.T(), tc.id, pr.Data.ID)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *PostSuite) TestLike() {
	var np people.NewPost
	var au people.AuthUser
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p1, _ := s.pr.Create(np, u.ID, nil)
	p2, _ := s.pr.Create(np, u.ID, nil)
	s.ps.Like(p1.ID, u.ID)

	conflictError := people.ConflictError

	tests := map[string]struct {
		id    uint
		valid bool
		kind  *people.ErrorKind
	}{
		"already liked": {p1.ID, false, &conflictError},
		"valid":         {p2.ID, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			pr, err := s.ps.Like(tc.id, u.ID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.True(s.T(), pr.Data.Status.IsLiked)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *PostSuite) TestUnlike() {
	var np people.NewPost
	var au people.AuthUser
	gofakeit.Struct(&np)
	gofakeit.Struct(&au)
	u, _ := s.ur.Create(au)
	p1, _ := s.pr.Create(np, u.ID, nil)
	p2, _ := s.pr.Create(np, u.ID, nil)
	s.ps.Like(p2.ID, u.ID)

	notFoundError := people.NotFoundError

	tests := map[string]struct {
		id    uint
		valid bool
		kind  *people.ErrorKind
	}{
		"not liked": {p1.ID, false, &notFoundError},
		"valid":     {p2.ID, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			pr, err := s.ps.Unlike(tc.id, u.ID)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				assert.False(s.T(), pr.Data.Status.IsLiked)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *PostSuite) SetupTest() {
	um := map[uint]people.User{}
	fm := map[inmem.FollowKey]time.Time{}
	pm := map[uint]people.Post{}
	lm := map[inmem.LikeKey]struct{}{}
	im := map[uint]people.Image{}
	v := validator.New()
	s.pr = inmem.NewPostRepository(pm)
	s.ur = inmem.NewUserRepository(um)
	fr := inmem.NewFollowRepository(fm, um)
	lr := inmem.NewLikeRepository(lm, pm, um)
	ir := inmem.NewImageRepository(im)
	us := user.NewService(v, s.ur, fr, lr)
	is := image.NewService(ir)
	s.ps = post.NewService(v, s.pr, s.ur, fr, lr, us, is)
}

func TestPostSuite(t *testing.T) {
	suite.Run(t, new(PostSuite))
}
