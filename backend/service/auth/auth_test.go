package auth_test

import (
	"testing"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/repository/inmem"
	"github.com/toxeeec/people/backend/service/auth"
	"github.com/toxeeec/people/backend/service/image"
	"github.com/toxeeec/people/backend/service/user"
)

type AuthSuite struct {
	suite.Suite
	as auth.Service
	ur repository.User
	us user.Service
}

func (s *AuthSuite) TestRegister() {
	var taken people.AuthUser
	var u people.AuthUser
	gofakeit.Struct(&taken)
	gofakeit.Struct(&u)
	s.ur.Create(taken)

	validationError := people.ValidationError

	tests := map[string]struct {
		au    people.AuthUser
		valid bool
		kind  *people.ErrorKind
	}{

		"taken handle": {taken, false, &validationError},
		"valid":        {u, true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			au, err := s.as.Register(tc.au)
			if tc.valid {
				assert.Equal(s.T(), tc.au.Handle, au.User.Handle)
				assert.NoError(s.T(), err)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *AuthSuite) TestLogin() {
	var invalidPassword people.AuthUser
	var unknownHandle people.AuthUser
	var u people.AuthUser
	gofakeit.Struct(&invalidPassword)
	gofakeit.Struct(&unknownHandle)
	gofakeit.Struct(&u)
	invalidPassword.Handle = u.Handle
	s.as.Register(u)

	validationError := people.ValidationError

	tests := map[string]struct {
		au    people.AuthUser
		valid bool
		kind  *people.ErrorKind
	}{
		"unknown handle":   {unknownHandle, false, &validationError},
		"invalid password": {invalidPassword, false, &validationError},
		"valid":            {u, true, nil},
	}
	for name, tc := range tests {
		s.Run(name, func() {
			au, err := s.as.Login(tc.au)
			if tc.valid {
				assert.Equal(s.T(), tc.au.Handle, au.User.Handle)
				assert.NoError(s.T(), err)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)
			}
		})
	}
}

func (s *AuthSuite) TestRefresh() {
	// sleep is used so every token is different
	var u people.AuthUser
	gofakeit.Struct(&u)
	ar, _ := s.as.Register(u)
	ts := ar.Tokens

	time.Sleep(time.Second)
	ts, err := s.as.Refresh(ts.RefreshToken)
	assert.NoError(s.T(), err)

	// create new token
	time.Sleep(time.Second)
	newTS, err := s.as.Refresh(ts.RefreshToken)
	assert.NoError(s.T(), err)

	authError := people.AuthError
	var e *people.Error
	// try using the previous token
	time.Sleep(time.Second)
	_, err = s.as.Refresh(ts.RefreshToken)
	assert.ErrorAs(s.T(), err, &e)
	assert.Equal(s.T(), authError, *e.Kind)

	// new token is also invalidated now
	time.Sleep(time.Second)
	_, err = s.as.Refresh(newTS.RefreshToken)
	assert.ErrorAs(s.T(), err, &e)
	assert.Equal(s.T(), authError, *e.Kind)
}

func (s *AuthSuite) TestLogout() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	ar1, _ := s.as.Register(au)
	ts1 := ar1.Tokens
	ar2, _ := s.as.Login(au)
	ts2 := ar2.Tokens

	ts1, err := s.as.Refresh(ts1.RefreshToken)
	assert.NoError(s.T(), err)
	ts2, err = s.as.Refresh(ts2.RefreshToken)
	assert.NoError(s.T(), err)

	logoutFromAll := true
	err = s.as.Logout(ts1.RefreshToken, &logoutFromAll)
	assert.NoError(s.T(), err)

	_, err = s.as.Refresh(ts1.RefreshToken)
	assert.Error(s.T(), err)
	_, err = s.as.Refresh(ts2.RefreshToken)
	assert.Error(s.T(), err)
}

func (s *AuthSuite) TestDelete() {
	var au people.AuthUser
	gofakeit.Struct(&au)
	ar, _ := s.as.Register(au)

	validationError := people.ValidationError

	tests := map[string]struct {
		password string
		valid    bool
		kind     *people.ErrorKind
	}{

		"invalid password": {gofakeit.Password(true, true, true, true, true, 8), false, &validationError},
		"valid":            {au.Password, true, nil},
	}

	for name, tc := range tests {
		s.Run(name, func() {
			err := s.as.Delete(ar.User.ID, tc.password)
			assert.Equal(s.T(), tc.valid, err == nil)
			if tc.valid {
				_, err := s.ur.Get(ar.User.ID)
				assert.Error(s.T(), err)
			} else {
				var e *people.Error
				assert.ErrorAs(s.T(), err, &e)
				assert.Equal(s.T(), *tc.kind, *e.Kind)

			}
		})
	}
}

func (s *AuthSuite) SetupTest() {
	um := map[uint]people.User{}
	tsm := map[uuid.UUID]people.RefreshToken{}
	fm := map[inmem.FollowKey]time.Time{}
	lm := map[inmem.LikeKey]struct{}{}
	pm := map[uint]people.Post{}
	im := map[uint]people.Image{}
	v := validator.New()
	s.ur = inmem.NewUserRepository(um)
	tr := inmem.NewTokenRepository(tsm)
	fr := inmem.NewFollowRepository(fm, um)
	lr := inmem.NewLikeRepository(lm, pm, um)
	ir := inmem.NewImageRepository(im)
	is := image.NewService(ir)
	s.us = user.NewService(v, s.ur, fr, lr, is)
	s.as = auth.NewService(s.ur, tr, s.us)
}

func TestAuthSuite(t *testing.T) {
	suite.Run(t, new(AuthSuite))
}
