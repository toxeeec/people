package http

import (
	"context"
	"errors"

	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/post"
)

func (h *handler) PostPosts(ctx context.Context, r people.PostPostsRequestObject) (people.PostPostsResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	pr, err := h.ps.Create(ctx, *r.Body, userID, nil)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PostPosts400JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostPosts200JSONResponse(pr), nil
}

func (h *handler) GetPostsPostID(ctx context.Context, r people.GetPostsPostIDRequestObject) (people.GetPostsPostIDResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	pr, err := h.ps.Get(ctx, r.PostID, userID, ok)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.GetPostsPostID404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetPostsPostID200JSONResponse(pr), nil
}

func (h *handler) DeletePostsPostID(ctx context.Context, r people.DeletePostsPostIDRequestObject) (people.DeletePostsPostIDResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	err := h.ps.Delete(r.PostID, userID)
	if err != nil {
		return nil, err
	}
	return people.DeletePostsPostID204Response{}, nil
}

func (h *handler) GetUsersHandlePosts(ctx context.Context, r people.GetUsersHandlePostsRequestObject) (people.GetUsersHandlePostsResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	prs, err := h.ps.ListUserPosts(ctx, r.Handle, userID, ok, post.IDPaginationParams(r.Params))
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.GetUsersHandlePosts404JSONResponse(*e), nil
			case people.NotFoundError:
				return people.GetUsersHandlePosts404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.GetUsersHandlePosts200JSONResponse(prs), nil
}

func (h *handler) GetMeFeed(ctx context.Context, r people.GetMeFeedRequestObject) (people.GetMeFeedResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	prs, err := h.ps.ListFeed(ctx, userID, post.IDPaginationParams(r.Params))
	if err != nil {
		return nil, err
	}
	return people.GetMeFeed200JSONResponse(prs), nil
}

func (h *handler) PostPostsPostIDReplies(ctx context.Context, r people.PostPostsPostIDRepliesRequestObject) (people.PostPostsPostIDRepliesResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	pr, err := h.ps.Create(ctx, *r.Body, userID, &r.PostID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.ValidationError:
				return people.PostPostsPostIDReplies400JSONResponse(*e), nil
			case people.NotFoundError:
				return people.PostPostsPostIDReplies404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PostPostsPostIDReplies200JSONResponse(pr), nil
}

func (h *handler) GetPostsPostIDReplies(ctx context.Context, r people.GetPostsPostIDRepliesRequestObject) (people.GetPostsPostIDRepliesResponseObject, error) {
	userID, ok := fromContext(ctx, userIDKey)
	ps, err := h.ps.ListReplies(ctx, r.PostID, userID, ok, post.IDPaginationParams(r.Params))
	if err != nil {
		return nil, err
	}
	return people.GetPostsPostIDReplies200JSONResponse(ps), nil
}

func (h *handler) PutPostsPostIDLikes(ctx context.Context, r people.PutPostsPostIDLikesRequestObject) (people.PutPostsPostIDLikesResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	pr, err := h.ps.Like(ctx, r.PostID, userID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.PutPostsPostIDLikes404JSONResponse(*e), nil
			case people.ConflictError:
				return people.PutPostsPostIDLikes409JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.PutPostsPostIDLikes200JSONResponse(pr), nil
}

func (h *handler) DeletePostsPostIDLikes(ctx context.Context, r people.DeletePostsPostIDLikesRequestObject) (people.DeletePostsPostIDLikesResponseObject, error) {
	userID, _ := fromContext(ctx, userIDKey)
	pr, err := h.ps.Unlike(ctx, r.PostID, userID)
	if err != nil {
		var e *people.Error
		if errors.As(err, &e) {
			switch *e.Kind {
			case people.NotFoundError:
				return people.DeletePostsPostIDLikes404JSONResponse(*e), nil
			}
		}
		return nil, err
	}
	return people.DeletePostsPostIDLikes200JSONResponse(pr), nil
}
