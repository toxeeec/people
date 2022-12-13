package post

import people "github.com/toxeeec/people/backend"

type Slice []people.Post

func (s Slice) UserIDs() []uint {
	ids := make([]uint, len(s))
	for i, p := range s {
		ids[i] = p.UserID
	}
	return ids
}

func (s Slice) IDs() []uint {
	ids := make([]uint, len(s))
	for i, p := range s {
		ids[i] = p.ID
	}
	return ids
}

func (s Slice) AddStatus(lss map[uint]people.LikeStatus) {
	for i, p := range s {
		ls := lss[p.ID]
		s[i].Status = &ls
	}
}
