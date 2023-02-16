package post

import people "github.com/toxeeec/people/backend"

func userIDs(ps []people.Post) []uint {
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
