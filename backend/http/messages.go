package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

func (h *handler) GetUsersHandleThread(ctx context.Context, r people.GetUsersHandleThreadRequestObject) (people.GetUsersHandleThreadResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	thread, err := h.ms.GetUsersThread(ctx, userID, r.Handle)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetUsersHandleThread404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetUsersHandleThread200JSONResponse(thread), nil
}

func (h *handler) GetThreadsThreadID(ctx context.Context, r people.GetThreadsThreadIDRequestObject) (people.GetThreadsThreadIDResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	thread, err := h.ms.GetThread(ctx, userID, r.ThreadID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.AuthError:
				return people.GetThreadsThreadID403JSONResponse(*e), nil
			case people.NotFoundError:
				return people.GetThreadsThreadID404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetThreadsThreadID200JSONResponse(thread), nil
}

func (h *handler) GetThreadsThreadIDMessages(ctx context.Context, r people.GetThreadsThreadIDMessagesRequestObject) (people.GetThreadsThreadIDMessagesResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	msgs, err := h.ms.ListThreadMessages(ctx, r.ThreadID, userID, pagination.IDParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.AuthError:
				return people.GetThreadsThreadIDMessages403JSONResponse(*e), nil
			case people.NotFoundError:
				return people.GetThreadsThreadIDMessages404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetThreadsThreadIDMessages200JSONResponse(msgs), nil
}

func (h *handler) GetThreads(ctx context.Context, r people.GetThreadsRequestObject) (people.GetThreadsResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	is, err := h.ms.ListThreads(ctx, userID, pagination.IDParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetThreads404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetThreads200JSONResponse(is), nil
}
