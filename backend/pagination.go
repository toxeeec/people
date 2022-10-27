package people

type Pagination struct {
	Offset uint
	Limit  uint
}

type PaginationMode int

const (
	PaginationModeNone PaginationMode = iota
	PaginationModeBefore
	PaginationModeAfter
	PaginationModeBeforeAfter
)

type SeekPagination struct {
	Mode   PaginationMode
	Before uint
	After  uint
	Limit  uint
}

const (
	offsetDefault = 0
	limitDefault  = 20
	limitMax      = 100
)

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

func NewSeekPagination(before, after, limit *uint) SeekPagination {
	p := SeekPagination{PaginationModeNone, 0, 0, limitDefault}
	if limit != nil && *limit <= limitMax {
		p.Limit = *limit
	}
	if before != nil {
		p.Before = *before
		p.Mode = PaginationModeBefore
	}
	if after != nil {
		p.After = *after
		if p.Mode != PaginationModeNone {
			p.Mode = PaginationModeBeforeAfter
		} else {
			p.Mode = PaginationModeAfter
		}
	}
	return p
}
