package people

const (
	offsetDefault = 0
	limitDefault  = 20
	limitMax      = 100
)

type Pagination struct {
	Offset uint
	Limit  uint
}

func NewPagination(page, limit *uint) Pagination {
	p := Pagination{offsetDefault, limitDefault}
	if limit != nil && *limit <= limitMax {
		p.Limit = *limit
	}
	if page != nil {
		p.Offset = (*page - 1) * p.Limit
	}
	return p
}
