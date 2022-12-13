package user

import people "github.com/toxeeec/people/backend"

type Slice []people.User

func (s Slice) AddStatus(fss map[uint]people.FollowStatus) {
	for i, u := range s {
		fs := fss[u.ID]
		s[i].Status = &fs
	}
}

func (s Slice) IDs() []uint {
	ids := make([]uint, len(s))
	for i, u := range s {
		ids[i] = u.ID
	}
	return ids
}

func (s Slice) ToMap() map[uint]people.User {
	m := make(map[uint]people.User, len(s))
	for _, v := range s {
		m[v.ID] = v
	}
	return m
}
