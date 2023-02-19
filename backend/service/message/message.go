package message

import (
	"context"
	"encoding/json"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/service/notification"
	"github.com/toxeeec/people/backend/service/user"
	"github.com/toxeeec/people/backend/set"
	"golang.org/x/exp/slices"
	"golang.org/x/sync/errgroup"
)

type Service interface {
	ReadMessage(ctx context.Context, from uint, data []byte) error
	GetUsersThread(ctx context.Context, userID uint, users ...string) (people.Thread, error)
	GetThread(ctx context.Context, userID, threadID uint) (people.Thread, error)
	ListThreadMessages(ctx context.Context, threadID uint, userID uint, params pagination.IDParams) (people.Messages, error)
	ListThreads(ctx context.Context, userID uint, params pagination.IDParams) (people.Threads, error)
}

type messageService struct {
	mr repository.Message
	ur repository.User
	ns notification.Service
	us user.Service
}

func NewService(mr repository.Message, ur repository.User, ns notification.Service, us user.Service) Service {
	return &messageService{mr, ur, ns, us}
}

func (s *messageService) ReadMessage(ctx context.Context, fromID uint, data []byte) error {
	var um people.UserMessage
	if err := json.Unmarshal(data, &um); err != nil {
		return service.NewError(people.ValidationError, "Invalid message format")
	}
	um = trim(um)
	if err := validate(um); err != nil {
		return err
	}
	userIDs, err := s.mr.GetThreadUsers(um.ThreadID)
	if err != nil {
		return service.NewError(people.NotFoundError, "Thread not found")
	}
	if !slices.Contains(userIDs, fromID) {
		return service.NewError(people.AuthError, "You do not have permission to send messages to this thread")
	}

	g, ctx := errgroup.WithContext(ctx)
	dbmsgc := make(chan people.DBMessage, 1)
	fromc := make(chan people.User, 1)
	g.Go(func() error {
		dbm, err := s.mr.Create(um.ThreadID, um.Content, fromID)
		if err != nil {
			return err
		}
		select {
		case dbmsgc <- dbm:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	g.Go(func() error {
		from, err := s.ur.Get(fromID)
		if err != nil {
			return err
		}
		select {
		case fromc <- from:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return err
	}
	msg := IntoMessage(<-dbmsgc, <-fromc)
	s.ns.Notify(people.MessageNotification, msg, userIDs...)
	return nil
}

func (s *messageService) GetUsersThread(ctx context.Context, userID uint, handles ...string) (people.Thread, error) {
	ids, err := s.ur.ListIDs(handles...)
	if err != nil || len(ids) != len(handles) {
		return people.Thread{}, service.NewError(people.NotFoundError, "User not found")
	}
	if !slices.Contains(ids, userID) {
		ids = append(ids, userID)
	}
	threadID, err := s.mr.GetThreadID(ids...)
	if err != nil {
		// thread doesn't exist
		threadID, err = s.mr.CreateThread(ids...)
		if err != nil {
			return people.Thread{}, service.InternalServerError
		}
	}
	g, ctx := errgroup.WithContext(ctx)
	thread := people.Thread{ID: threadID}
	g.Go(func() error {
		thread.Latest, err = s.getLatestMessage(threadID)
		if err != nil {
			return err
		}
		return nil
	})
	g.Go(func() error {
		users, err := s.getThreadUsers(threadID)
		if err != nil {
			return err
		}
		thread.Users = users
		return nil
	})
	if err := g.Wait(); err != nil {
		return people.Thread{}, err
	}
	return thread, nil
}

func (s *messageService) GetThread(ctx context.Context, userID uint, threadID uint) (people.Thread, error) {
	thread := people.Thread{ID: threadID}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		thread.Users, err = s.getThreadUsers(threadID)
		if err != nil {
			return err
		}
		found := false
		for _, u := range thread.Users {
			if u.ID == userID {
				found = true
				break
			}
		}
		if !found {
			return service.NewError(people.AuthError, "You do not have permission to view this thread")
		}
		return nil
	})
	g.Go(func() error {
		thread.Latest, _ = s.getLatestMessage(threadID)
		return nil
	})
	if err := g.Wait(); err != nil {
		println(err.Error())
		return people.Thread{}, err
	}
	return thread, nil
}

func (s *messageService) ListThreadMessages(ctx context.Context, threadID uint, userID uint, params pagination.IDParams) (people.Messages, error) {
	p := pagination.New(params)
	g, ctx := errgroup.WithContext(ctx)
	dbmsgsc := make(chan []people.DBMessage, 1)
	g.Go(func() error {
		dbmsgs, err := s.mr.ListThreadMessages(threadID, p)
		if err != nil {
			return err
		}
		select {
		case dbmsgsc <- dbmsgs:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	g.Go(func() error {
		ids, err := s.mr.GetThreadUsers(threadID)
		if err != nil {
			return err
		}
		if !slices.Contains(ids, userID) {
			return service.NewError(people.AuthError, "You do not have permission to view this thread")
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return people.Messages{}, err
	}
	dbmsgs := <-dbmsgsc
	ids := userIDs(dbmsgs)
	users, err := s.ur.List(ids)
	if err != nil {
		return people.Messages{}, err
	}
	usersMap := make(map[uint]people.User, len(users))
	for _, u := range users {
		usersMap[u.ID] = u
	}
	msgs := IntoMessages(dbmsgs, usersMap)
	return pagination.NewResults[people.Message, uint](msgs), nil
}

func (s *messageService) ListThreads(ctx context.Context, userID uint, params pagination.IDParams) (people.Threads, error) {
	p := pagination.New(params)
	ids, err := s.mr.ListThreadIDs(userID, p)
	if err != nil {
		return people.Threads{}, service.NewError(people.AuthError, "Thread not found")
	}
	tUsersc := make(chan []people.ThreadUser, 1)
	dbmsgsc := make(chan []people.DBMessage, 1)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		tUsers, err := s.mr.ListThreadUsers(ids...)
		if err != nil {
			return err
		}
		select {
		case tUsersc <- tUsers:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	g.Go(func() error {
		dbmsgs, err := s.mr.ListLatestMessages(ids...)
		if err != nil {
			return err
		}
		select {
		case dbmsgsc <- dbmsgs:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if err := g.Wait(); err != nil {
		return people.Threads{}, nil
	}
	tUsers := <-tUsersc
	userIDs := make([]uint, len(tUsers))
	for i, tu := range tUsers {
		userIDs[i] = tu.UserID
	}
	userIDs = set.FromSlice(userIDs).Slice()
	users, err := s.ur.List(userIDs)
	if err != nil {
		return people.Threads{}, err
	}

	usersMap := make(map[uint]people.User, len(users))
	for _, u := range users {
		usersMap[u.ID] = u
	}
	tUsersMap := make(map[uint][]people.User)
	for _, tu := range tUsers {
		us := tUsersMap[tu.ID]
		us = append(us, usersMap[tu.UserID])
		tUsersMap[tu.ID] = us
	}
	dbmsgs := <-dbmsgsc
	msgs := IntoMessages(dbmsgs, usersMap)
	msgsMap := make(map[uint]people.Message, len(dbmsgs))
	for _, msg := range msgs {
		msgsMap[msg.ThreadID] = msg
	}
	threads := make([]people.Thread, len(ids))
	for i := range threads {
		id := ids[i]
		threads[i].ID = id
		threads[i].Users = tUsersMap[id]
		msg, ok := msgsMap[id]
		if ok {
			threads[i].Latest = &msg
		}
	}
	return pagination.NewResults[people.Thread, uint](threads), nil
}

func (s *messageService) getThreadUsers(threadID uint) ([]people.User, error) {
	ids, err := s.mr.GetThreadUsers(threadID)
	if err != nil {
		return nil, err
	}
	return s.ur.List(ids)
}

func (s *messageService) getLatestMessage(threadID uint) (*people.Message, error) {
	dbmsg, err := s.mr.GetLatestMessage(threadID)
	if err != nil {
		return nil, nil
	}
	u, err := s.ur.Get(dbmsg.FromID)
	if err != nil {
		return nil, err
	}
	msg := IntoMessage(dbmsg, u)
	return &msg, nil
}
