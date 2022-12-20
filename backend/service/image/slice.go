package image

import people "github.com/toxeeec/people/backend"

type Slice []people.Image

func (s Slice) IDs() []uint {
	ids := make([]uint, len(s))
	for i, p := range s {
		ids[i] = p.ID
	}
	return ids
}
