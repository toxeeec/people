package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/pagination"
)

func (h *handler) GetMessagesHandle(ctx context.Context, r people.GetMessagesHandleRequestObject) (people.GetMessagesHandleResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	ums, err := h.ms.ListUserMessages(r.Handle, userID, pagination.IDParams{Limit: r.Params.Limit, Before: r.Params.Before, After: r.Params.After})
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetMessagesHandle404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetMessagesHandle200JSONResponse(ums), nil
}
