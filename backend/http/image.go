package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
)

func (h *handler) PostImages(ctx context.Context, r people.PostImagesRequestObject) (people.PostImagesResponseObject, error) {
	userID, _ := people.FromContext(ctx, people.UserIDKey)
	ir, err := h.is.Create(userID, r.Body)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PostImages400JSONResponse(*e), nil
			}
		}
		return nil, err
	}

	return people.PostImages200JSONResponse(ir), nil
}
