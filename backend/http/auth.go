package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
)

func (h *handler) PostRegister(ctx context.Context, r people.PostRegisterRequestObject) (people.PostRegisterResponseObject, error) {
	ar, err := h.as.Register(*r.Body)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PostRegister400JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostRegister200JSONResponse(ar), nil
}

func (h *handler) PostLogin(ctx context.Context, r people.PostLoginRequestObject) (people.PostLoginResponseObject, error) {
	ar, err := h.as.Login(*r.Body)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PostLogin401JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostLogin200JSONResponse(ar), nil
}

func (h *handler) PostRefresh(ctx context.Context, r people.PostRefreshRequestObject) (people.PostRefreshResponseObject, error) {
	ts, err := h.as.Refresh(r.Body.RefreshToken)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.AuthError:
				return people.PostRefresh403JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostRefresh200JSONResponse(ts), nil
}

//TODO: logout (remove refresh token from db)
