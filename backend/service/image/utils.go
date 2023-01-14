package image

import people "github.com/toxeeec/people/backend"

func IDs(is []people.Image) []uint {
	ids := make([]uint, len(is))
	for i, p := range is {
		ids[i] = p.ID
	}
	return ids
}
