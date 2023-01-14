package user

import (
	"context"

	people "github.com/toxeeec/people/backend"
	"golang.org/x/sync/errgroup"
)

func (s *userService) GetUserWithStatus(ctx context.Context, srcID uint, userID uint, auth bool) (people.User, error) {
	uc := make(chan people.User, 1)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		u, err := s.ur.Get(srcID)
		if err != nil {
			return err
		}
		select {
		case uc <- u:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if auth {
		g.Go(func() error {
			fs := s.status(srcID, userID)
			select {
			case u := <-uc:
				u.Status = &fs
				uc <- u
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return people.User{}, err
	}
	return <-uc, nil
}

func (s *userService) ListStatus(ctx context.Context, userIDs []uint, srcID uint) (map[uint]people.FollowStatus, error) {
	var following map[uint]struct{}
	var followed map[uint]struct{}
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		var err error
		following, err = s.fr.ListStatusFollowing(userIDs, srcID)
		return err
	})
	g.Go(func() error {
		var err error
		followed, err = s.fr.ListStatusFollowed(userIDs, srcID)
		return err
	})
	if err := g.Wait(); err != nil {
		return nil, err
	}
	fss := make(map[uint]people.FollowStatus, len(userIDs))
	for _, id := range userIDs {
		_, followingOk := following[id]
		_, followedOk := followed[id]
		fss[id] = people.FollowStatus{IsFollowing: followingOk, IsFollowed: followedOk}
	}
	return fss, nil
}

func (s *userService) ListUsersWithStatus(ctx context.Context, userIDs []uint, srcID uint, auth bool) ([]people.User, error) {
	if len(userIDs) == 0 {
		return []people.User{}, nil
	}
	usc := make(chan []people.User, 1)
	g, ctx := errgroup.WithContext(ctx)
	g.Go(func() error {
		us, err := s.ur.List(userIDs)
		if err != nil {
			return err
		}
		select {
		case usc <- us:
		case <-ctx.Done():
			return ctx.Err()
		}
		return nil
	})
	if auth {
		g.Go(func() error {
			fss, err := s.ListStatus(context.Background(), userIDs, srcID)
			if err != nil {
				return err
			}
			select {
			case us := <-usc:
				AddStatuses(us, fss)
				usc <- us
			case <-ctx.Done():
				return ctx.Err()
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return nil, err
	}
	return <-usc, nil
}

func (s *userService) status(srcID, userID uint) people.FollowStatus {
	followingc := make(chan bool)
	followedc := make(chan bool)
	go func() {
		followingc <- s.fr.GetStatusFollowing(srcID, userID)
	}()
	go func() {
		followedc <- s.fr.GetStatusFollowed(srcID, userID)
	}()
	return people.FollowStatus{IsFollowing: <-followingc, IsFollowed: <-followedc}
}
