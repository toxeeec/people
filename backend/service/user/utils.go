package user

import people "github.com/toxeeec/people/backend"

func AddStatuses(us []people.User, fss map[uint]people.FollowStatus) {
	for i, u := range us {
		fs := fss[u.ID]
		us[i].Status = &fs
	}
}

func AddImages(us []people.User, imgs map[uint]*string) {
	for i, u := range us {
		img := imgs[u.ID]
		us[i].Image = img
	}
}

// TODO: use people.IntoIDs instead
func IDs(us []people.User) []uint {
	ids := make([]uint, len(us))
	for i, u := range us {
		ids[i] = u.ID
	}
	return ids
}

// TODO: use people.IntoMap instead
func IntoMap(us []people.User) map[uint]people.User {
	m := make(map[uint]people.User, len(us))
	for _, v := range us {
		m[v.ID] = v
	}
	return m
}
