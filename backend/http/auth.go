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

func (h *handler) PostLogout(ctx context.Context, r people.PostLogoutRequestObject) (people.PostLogoutResponseObject, error) {
	err := h.as.Logout(r.Body.RefreshToken, r.Body.LogoutFromAll)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.AuthError:
				return people.PostLogout403JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostLogout204Response{}, nil
}

func (h *handler) DeleteMe(ctx context.Context, r people.DeleteMeRequestObject) (people.DeleteMeResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	err := h.as.Delete(userID, r.Body.Password, r.Body.RefreshToken)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.DeleteMe401JSONResponse(*e), nil
			case people.AuthError:
				return people.DeleteMe403JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.DeleteMe204Response{}, nil
}
