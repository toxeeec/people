package people

import "github.com/jmoiron/sqlx"

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

func SeekPaginationQueries(base, end, before, after, beforeAfter string) [4]string {
	return [4]string{
		base + end,
		base + before + end,
		base + after + end,
		base + beforeAfter + end,
	}
}

type SeekPaginationResult[T Identifier] struct {
	Data []T
	Meta *SeekPaginationMeta
}

func SeekPaginationSelect[T Identifier](db *sqlx.DB, queries *[4]string, p SeekPagination, params ...any) (SeekPaginationResult[T], error) {
	var res SeekPaginationResult[T]
	res.Data = make([]T, 0, p.Limit)
	query := queries[p.Mode]
	switch p.Mode {
	case PaginationModeNone:
		params = append(params, p.Limit)
	case PaginationModeBefore:
		params = append(params, p.Limit, p.Before)
	case PaginationModeAfter:
		params = append(params, p.Limit, p.After)
	case PaginationModeBeforeAfter:
		params = append(params, p.Limit, p.Before, p.After)
	}
	err := db.Select(&res.Data, query, params...)
	if err != nil {
		return SeekPaginationResult[T]{}, err
	}

	if len(res.Data) > 0 {
		res.Meta = &SeekPaginationMeta{
			NewestID: res.Data[0].Identify(),
			OldestID: res.Data[len(res.Data)-1].Identify(),
		}
	}
	return res, nil
}
