package pagination

import (
	"context"

	people "github.com/toxeeec/people/backend"
	"golang.org/x/sync/errgroup"
)

type Mode uint

const (
	ModeNone        Mode = 0
	ModeBefore           = 0b01
	ModeAfter            = 0b10
	ModeBeforeAfter      = ModeBefore | ModeAfter
)

type Pagination[T any] struct {
	Before *T
	After  *T
	Limit  uint
}

type ID = Pagination[uint]
type Handle Pagination[string]

func New[T any](before, after *T, limit *uint) Pagination[T] {
	const limitDefault, limitMax = 20, 100
	p := Pagination[T]{Limit: limitDefault}
	if limit != nil && *limit <= limitMax {
		p.Limit = *limit
	}
	if before != nil {
		p.Before = before
	}
	if after != nil {
		p.After = after
	}
	return p
}

func NewResults[T people.Identifier[U], U any](data []T) people.PaginatedResults[T, U] {
	res := people.PaginatedResults[T, U]{Data: data}
	if len(res.Data) > 0 {
		res.Meta = &people.PaginationMeta[U]{
			Newest: res.Data[0].Identify(),
			Oldest: res.Data[len(res.Data)-1].Identify(),
		}
	}
	if res.Data == nil {
		res.Data = []T{}
	}
	return res
}

type GetIDFn func(string) (uint, error)

func (hp Handle) IDPagination(ctx context.Context, getIDfn GetIDFn) (ID, error) {
	p := ID{Limit: hp.Limit}
	g, ctx := errgroup.WithContext(ctx)
	if hp.Before != nil {
		g.Go(func() error {
			id, err := getIDfn(*hp.Before)
			if err != nil {
				return err
			}
			p.Before = &id
			return nil
		})
	}
	if hp.After != nil {
		g.Go(func() error {
			id, err := getIDfn(*hp.After)
			if err != nil {
				return err
			}
			p.After = &id
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return ID{}, err
	}
	return p, nil
}
