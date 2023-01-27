package auth

import (
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/service/user"
	"golang.org/x/crypto/bcrypt"
)

type Service interface {
	Register(au people.AuthUser) (people.AuthResponse, error)
	Login(au people.AuthUser) (people.AuthResponse, error)
	Refresh(refreshToken string) (people.Tokens, error)
	Logout(rtString string, logoutFromAll *bool) error
	Delete(userID uint, password string) error
}

type authService struct {
	ur repository.User
	tr repository.Token
	us user.Service
}

func NewService(ur repository.User, tr repository.Token, us user.Service) Service {
	return &authService{ur, tr, us}
}

func (s *authService) Register(au people.AuthUser) (people.AuthResponse, error) {
	err := s.us.Validate(au)
	if err != nil {
		return people.AuthResponse{}, err
	}

	bytes, err := bcrypt.GenerateFromPassword([]byte(au.Password), bcrypt.DefaultCost)
	if err != nil {
		return people.AuthResponse{}, err
	}
	au.Password = string(bytes)
	u, err := s.ur.Create(au)
	if err != nil {
		return people.AuthResponse{}, err
	}
	ts, err := s.newTokens(u.ID)
	if err != nil {
		go s.ur.Delete(u.ID)
		return people.AuthResponse{}, err
	}
	return people.AuthResponse{User: u, Tokens: ts}, nil
}

func (s *authService) Login(au people.AuthUser) (people.AuthResponse, error) {
	id, err := s.ur.GetID(au.Handle)
	if err != nil {
		return people.AuthResponse{}, service.NewError(people.ValidationError, "Invalid handle or password")
	}
	hash, err := s.ur.GetHash(id)
	if err != nil {
		return people.AuthResponse{}, err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(au.Password))
	if err != nil {
		return people.AuthResponse{}, service.NewError(people.ValidationError, "Invalid handle or password")
	}

	ts, err := s.newTokens(id)
	if err != nil {
		return people.AuthResponse{}, err
	}
	u, err := s.ur.Get(id)
	if err != nil {
		return people.AuthResponse{}, err
	}
	return people.AuthResponse{User: u, Tokens: ts}, nil
}

func (s *authService) Refresh(rtString string) (people.Tokens, error) {
	rt, err := s.checkRefreshToken(rtString)
	if err != nil {
		return people.Tokens{}, err
	}

	at, err := NewAccessToken(rt.UserID)
	if err != nil {
		return people.Tokens{}, err
	}
	newRT, err := NewRefreshToken(rt.UserID, &rt.ID)
	if err != nil {
		return people.Tokens{}, err
	}
	err = s.tr.Update(newRT)
	if err != nil {
		go s.tr.Delete(rt.ID)
		return people.Tokens{}, err
	}
	return people.Tokens{AccessToken: at, RefreshToken: newRT.Value}, nil
}

func (s *authService) Logout(rtString string, logoutFromAll *bool) error {
	rt, err := parseRefreshToken(rtString)
	if err != nil {
		return service.NewError(people.AuthError, "Malformed refresh token")
	}
	if _, err := s.tr.Get(rt.Value); err != nil {
		// token doesn't exist
		go s.tr.Delete(rt.ID)
		return service.NewError(people.AuthError, "Invalid refresh token")
	}

	if logoutFromAll != nil && *logoutFromAll == true {
		return s.tr.DeleteAll(rt.UserID)
	}

	return s.tr.Delete(rt.ID)
}

func (s *authService) Delete(userID uint, password string) error {
	hash, err := s.ur.GetHash(userID)
	if err != nil {
		return err
	}
	err = bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	if err != nil {
		return service.NewError(people.ValidationError, "Invalid password")
	}
	return s.us.Delete(userID)
}
