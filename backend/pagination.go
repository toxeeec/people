package people

import "github.com/jmoiron/sqlx"

type PaginationMode int

const (
	PaginationModeNone PaginationMode = iota
	PaginationModeBefore
	PaginationModeAfter
	PaginationModeBeforeAfter
)

type PaginationMeta[T any] struct {
	Oldest T `json:"oldest"`
	Newest T `json:"newest"`
}

type Pagination[T any] struct {
	Mode   PaginationMode
	Before T
	After  T
	Limit  uint
}

type IDPagination = Pagination[uint]
type HandlePagination = Pagination[string]

const (
	offsetDefault = 0
	limitDefault  = 20
	limitMax      = 100
)

func NewPagination[T any](before, after *T, limit *uint) Pagination[T] {
	p := Pagination[T]{Mode: PaginationModeNone, Limit: limitDefault}
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

func PaginationQueries(base, end, before, after, beforeAfter string) [4]string {
	return [4]string{
		base + end,
		base + before + end,
		base + after + end,
		base + beforeAfter + end,
	}
}

type PaginationResult[T Identifier[U], U any] struct {
	Data []T                `json:"data"`
	Meta *PaginationMeta[U] `json:"meta,omitempty"`
}

func PaginationSelect[T Identifier[U], U any](db *sqlx.DB, queries *[4]string, p Pagination[U], params ...any) (PaginationResult[T, U], error) {
	var res PaginationResult[T, U]
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
		return PaginationResult[T, U]{}, err
	}

	if len(res.Data) > 0 {
		res.Meta = &PaginationMeta[U]{
			Newest: res.Data[0].Identify(),
			Oldest: res.Data[len(res.Data)-1].Identify(),
		}
	}
	return res, nil
}
