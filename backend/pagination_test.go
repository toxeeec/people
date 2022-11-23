package people_test

import (
	"testing"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/stretchr/testify/assert"
	people "github.com/toxeeec/people/backend"
	"github.com/toxeeec/people/backend/service/post"
	"github.com/toxeeec/people/backend/service/user"
)

// TestPaginationSelect uses Feed but it can be replaced with any method that uses Pagination
func TestPaginationSelect(t *testing.T) {
	db, _ := people.PostgresConnect()
	db.MustExec("TRUNCATE user_profile CASCADE")
	db.MustExec("TRUNCATE post CASCADE")
	us := user.NewService(db)
	ps := post.NewService(db)
	var user people.AuthUser
	gofakeit.Struct(&user)
	userID, _ := us.Create(user)

	var oldest uint
	var before uint
	var after uint
	var newest uint

	count := 7
	// - after - - before - -
	for i := 0; i < count; i++ {
		var u people.AuthUser
		gofakeit.Struct(&u)
		id, _ := us.Create(u)
		us.Follow(userID, u.Handle)
		var post people.PostBody
		gofakeit.Struct(&post)
		p, _ := ps.Create(id, post)
		switch i {
		case 0:
			oldest = p.ID
		case 1:
			after = p.ID
		case 4:
			before = p.ID
		case count - 1:
			newest = p.ID
		}
	}

	tests := map[string]struct {
		pagination people.Pagination[uint]
		oldest     uint
		newest     uint
		count      int
	}{
		"no pagination":             {people.NewPagination[uint](nil, nil, nil), oldest, newest, count},
		"before":                    {people.NewPagination(&before, nil, nil), oldest, before - 1, 4},
		"after":                     {people.NewPagination(nil, &after, nil), after + 1, newest, 5},
		"after greater than before": {people.NewPagination(&after, &before, nil), 0, 0, 0},
		"before and after":          {people.NewPagination(&before, &after, nil), after + 1, before - 1, 2},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			res, _ := ps.Feed(userID, tc.pagination)
			assert.Equal(t, tc.count, len(res.Data))
			if len(res.Data) == 0 {
				assert.Nil(t, res.Meta)
			} else {
				assert.Equal(t, tc.oldest, res.Meta.Oldest)
				assert.Equal(t, tc.newest, res.Meta.Newest)
			}
		})
	}
}
