package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/user"
)

func (h *handler) GetUsersHandle(ctx context.Context, r people.GetUsersHandleRequestObject) (people.GetUsersHandleResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	u, err := h.us.GetUser(ctx, r.Handle, userID, ok)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetUsersHandle404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetUsersHandle200JSONResponse(u), nil
}

func (h *handler) PutMeFollowingHandle(ctx context.Context, r people.PutMeFollowingHandleRequestObject) (people.PutMeFollowingHandleResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	u, err := h.us.Follow(ctx, r.Handle, userID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.PutMeFollowingHandle404JSONResponse(*e), nil
			case people.ConflictError:
				return people.PutMeFollowingHandle409JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PutMeFollowingHandle200JSONResponse(u), nil
}

func (h *handler) DeleteMeFollowingHandle(ctx context.Context, r people.DeleteMeFollowingHandleRequestObject) (people.DeleteMeFollowingHandleResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	u, err := h.us.Unfollow(ctx, r.Handle, userID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.DeleteMeFollowingHandle404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.DeleteMeFollowingHandle200JSONResponse(u), nil
}

func (h *handler) GetMeFollowing(ctx context.Context, r people.GetMeFollowingRequestObject) (people.GetMeFollowingResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	us, err := h.us.ListCurrUserFollowing(ctx, userID, user.HandlePaginationParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetMeFollowing404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetMeFollowing200JSONResponse(us), nil
}

func (h *handler) GetMeFollowers(ctx context.Context, r people.GetMeFollowersRequestObject) (people.GetMeFollowersResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	us, err := h.us.ListCurrUserFollowers(ctx, userID, user.HandlePaginationParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetMeFollowers404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetMeFollowers200JSONResponse(us), nil
}

func (h *handler) GetUsersHandleFollowing(ctx context.Context, r people.GetUsersHandleFollowingRequestObject) (people.GetUsersHandleFollowingResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	us, err := h.us.ListFollowing(ctx, r.Handle, userID, ok, user.HandlePaginationParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetUsersHandleFollowing404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetUsersHandleFollowing200JSONResponse(us), nil
}

func (h *handler) GetUsersHandleFollowers(ctx context.Context, r people.GetUsersHandleFollowersRequestObject) (people.GetUsersHandleFollowersResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	us, err := h.us.ListFollowers(ctx, r.Handle, userID, ok, user.HandlePaginationParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetUsersHandleFollowers404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetUsersHandleFollowers200JSONResponse(us), nil
}

func (h *handler) GetPostsPostIDLikes(ctx context.Context, r people.GetPostsPostIDLikesRequestObject) (people.GetPostsPostIDLikesResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	ur, err := h.us.ListPostLikes(ctx, r.PostID, userID, ok, user.HandlePaginationParams(r.Params))
	if err != nil {
		return nil, err
	}
	return people.GetPostsPostIDLikes200JSONResponse(ur), nil
}

func (h *handler) GetUsersSearch(ctx context.Context, r people.GetUsersSearchRequestObject) (people.GetUsersSearchResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	ur, err := h.us.ListMatches(ctx, r.Params.Query, userID, ok, user.HandlePaginationParams{Limit: r.Params.Limit, Before: r.Params.Before, After: r.Params.After})
	if err != nil {
		return nil, err
	}
	return people.GetUsersSearch200JSONResponse(ur), nil
}

func (h *handler) PutMe(ctx context.Context, r people.PutMeRequestObject) (people.PutMeResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	u, err := h.us.Update(userID, r.Body.Handle)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PutMe401JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PutMe200JSONResponse(u), nil
}
