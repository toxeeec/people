package user

import people "github.com/toxeeec/people/backend"

func AddStatuses(us []people.User, fss map[uint]people.FollowStatus) {
	for i, u := range us {
		fs := fss[u.ID]
		us[i].Status = &fs
	}
}

func IDs(us []people.User) []uint {
	ids := make([]uint, len(us))
	for i, u := range us {
		ids[i] = u.ID
	}
	return ids
}

func IntoMap(us []people.User) map[uint]people.User {
	m := make(map[uint]people.User, len(us))
	for _, v := range us {
		m[v.ID] = v
	}
	return m
}
