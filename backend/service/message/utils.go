package message

import (
	"strings"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service"
	"github.com/toxeeec/people/backend/set"
)

func trim(m people.UserMessage) people.UserMessage {
	m.Content = strings.TrimSpace(m.Content)
	return m
}

func validate(m people.UserMessage) error {
	if len(m.Content) == 0 {
		return service.NewError(people.ValidationError, "Content cannot be empty")
	}
	if m.ThreadID == 0 {
		return service.NewError(people.ValidationError, "Invalid thread ID")
	}
	return nil
}

func userIDs(msgs []people.DBMessage) []uint {
	set := set.New[uint]()
	for _, msg := range msgs {
		set.Add(msg.FromID)
	}
	return set.Slice()
}

func IntoMessage(msg people.DBMessage, from people.User) people.Message {
	return people.Message{ID: msg.ID, Content: msg.Content, From: from, ThreadID: msg.ThreadID, SentAt: msg.SentAt}
}

func IntoMessages(dbmsgs []people.DBMessage, users map[uint]people.User) []people.Message {
	msgs := make([]people.Message, len(dbmsgs))
	for i, dbmsg := range dbmsgs {
		msgs[i] = IntoMessage(dbmsg, users[dbmsg.FromID])
	}
	return msgs
}

func IntoUserIDs(users []people.ThreadUser) []uint {
	ids := make([]uint, len(users))
	for i, v := range users {
		ids[i] = v.UserID
	}
	return ids
}
