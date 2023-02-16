package inmem

import (
	"errors"
	"fmt"
	"math"
	"time"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
	"github.com/toxeeec/people/backend/repository"
	"golang.org/x/exp/slices"
)

type messageRepo struct {
	msgs             map[uint][]people.DBMessage
	threads          map[uint]struct{}
	threadUsers      map[uint][]uint
	um               map[uint]people.User
	lastMID          uint
	lastThreadID     uint
	lastThreadUserID uint
}

func (r *messageRepo) newMsgID() uint {
	r.lastMID++
	return r.lastMID
}
func (r *messageRepo) newThreadID() uint {
	r.lastThreadID++
	return r.lastThreadID
}
func (r *messageRepo) newThreadUserID() uint {
	r.lastThreadUserID++
	return r.lastThreadUserID
}

func NewMessageRepository(msgs map[uint][]people.DBMessage, threads map[uint]struct{}, threadUsers map[uint][]uint, um map[uint]people.User) repository.Message {
	return &messageRepo{msgs: msgs, threads: threads, threadUsers: threadUsers, um: um}
}

func (r *messageRepo) Create(threadID uint, content string, fromID uint) (people.DBMessage, error) {
	if _, ok := r.threads[threadID]; !ok {
		return people.DBMessage{}, fmt.Errorf("Message.Create: %w", errors.New("Thread not found"))
	}
	msgID := r.newMsgID()
	msg := people.DBMessage{ID: msgID, ThreadID: threadID, Content: content, FromID: fromID, SentAt: time.Now()}
	msgs := r.msgs[threadID]
	msgs = append(msgs, msg)
	r.msgs[threadID] = msgs
	return msg, nil
}

func (r *messageRepo) CreateThread(userIDs ...uint) (uint, error) {
	for _, userID := range userIDs {
		if _, ok := r.um[userID]; !ok {
			return 0, fmt.Errorf("Message.CreateThread: %w", errors.New("User not found"))
		}
	}
	threadID := r.newThreadID()
	r.threads[threadID] = struct{}{}
	for _, userID := range userIDs {
		tUsers := r.threadUsers[threadID]
		tUsers = append(tUsers, userID)
		r.threadUsers[threadID] = tUsers
	}
	return threadID, nil
}

func (r *messageRepo) GetThreadID(userIDs ...uint) (uint, error) {
	for k, thread := range r.threadUsers {
		if len(userIDs) != len(thread) {
			continue
		}
		subslice := true
		for _, userID := range userIDs {
			if !slices.Contains(thread, userID) {
				subslice = false
				break
			}
		}
		if subslice {
			return k, nil
		}
	}
	return 0, fmt.Errorf("Message.GetThreadID: %w", errors.New("Thread not found"))
}

func (r *messageRepo) ListThreadIDs(userID uint, p pagination.ID) ([]uint, error) {
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.After != nil {
		after = *p.After
	}
	beforeLatest, err := r.GetLatestMessage(before)
	if err != nil {
		beforeLatest.ID = math.MaxUint
	}
	afterLatest, err := r.GetLatestMessage(after)
	if err != nil {
		afterLatest.ID = 0
	}
	var threads []uint
	for threadID := range r.threads {
		latest, err := r.GetLatestMessage(threadID)
		if err != nil {
			continue
		}
		if slices.Contains(r.threadUsers[threadID], userID) && latest.ID < beforeLatest.ID && latest.ID > afterLatest.ID {
			threads = append(threads, threadID)
			if len(threads) == int(p.Limit) {
				return threads, nil
			}
		}
	}
	return threads, nil
}

func (r *messageRepo) GetThreadUsers(threadID uint) ([]uint, error) {
	tUsers, ok := r.threadUsers[threadID]
	if !ok {
		return nil, fmt.Errorf("Message.GetThreadUsers: %w", errors.New("Thread not found"))
	}
	return tUsers, nil
}

func (r *messageRepo) ListThreadUsers(threadIDs ...uint) ([]people.ThreadUser, error) {
	var users []people.ThreadUser
	for _, threadID := range threadIDs {
		tUsers, _ := r.threadUsers[threadID]
		for _, tUser := range tUsers {
			users = append(users, people.ThreadUser{ID: threadID, UserID: tUser})
		}
	}
	return users, nil
}

func (r *messageRepo) GetLatestMessage(threadID uint) (people.DBMessage, error) {
	if _, ok := r.threads[threadID]; !ok {
		return people.DBMessage{}, fmt.Errorf("Message.GetLatestMessage: %w", errors.New("Thread not found"))
	}
	msgs := r.msgs[threadID]
	var latest people.DBMessage
	for _, msg := range msgs {
		if msg.ID > latest.ID {
			latest = msg
		}
	}
	return latest, nil
}

func (r *messageRepo) ListLatestMessages(threadIDs ...uint) ([]people.DBMessage, error) {
	var msgs []people.DBMessage
	for _, threadID := range threadIDs {
		msgs = append(msgs, r.msgs[threadID]...)
	}
	return msgs, nil
}

func (r *messageRepo) ListThreadMessages(threadID uint, p pagination.ID) ([]people.DBMessage, error) {
	if _, ok := r.threads[threadID]; !ok {
		return nil, fmt.Errorf("Message.ListThreadMessages: %w", errors.New("Thread not found"))
	}
	before := uint(math.MaxUint)
	if p.Before != nil {
		before = *p.Before
	}
	after := uint(0)
	if p.After != nil {
		after = *p.After
	}
	var msgs []people.DBMessage
	for _, msg := range r.msgs[threadID] {
		if msg.ID < before && msg.ID > after {
			msgs = append(msgs, msg)
			if len(msgs) == int(p.Limit) {
				return msgs, nil
			}
		}
	}
	return msgs, nil
}
