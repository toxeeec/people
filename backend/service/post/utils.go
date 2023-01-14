package post

import people "github.com/toxeeec/people/backend"

func UserIDs(ps []people.Post) []uint {
	ids := make([]uint, len(ps))
	for i, p := range ps {
		ids[i] = p.UserID
	}
	return ids
}

func IDs(ps []people.Post) []uint {
	ids := make([]uint, len(ps))
	for i, p := range ps {
		ids[i] = p.ID
	}
	return ids
}

func AddStatus(ps []people.Post, lss map[uint]people.LikeStatus) {
	for i, p := range ps {
		ls := lss[p.ID]
		ps[i].Status = &ls
	}
}
