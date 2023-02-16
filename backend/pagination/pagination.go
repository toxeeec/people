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

type Params[T any] struct {
	Limit  *uint
	Before *T
	After  *T
}

type HandleParams = Params[string]
type IDParams = Params[uint]

// TODO: Remove pointers (use math.Min and math.Max if values are missing)
type Pagination[T any] struct {
	Before *T
	After  *T
	Limit  uint
}

type ID = Pagination[uint]
type Handle = Pagination[string]

func New[T any](params Params[T]) Pagination[T] {
	const limitDefault, limitMax = 20, 100
	p := Pagination[T]{Limit: limitDefault}
	if params.Limit != nil && *params.Limit <= limitMax {
		p.Limit = *params.Limit
	}
	if params.Before != nil {
		p.Before = params.Before
	}
	if params.After != nil {
		p.After = params.After
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

type GetIDFn[T any] func(T) (uint, error)

func IntoID[T any](ctx context.Context, p Pagination[T], getIDfn GetIDFn[T]) (ID, error) {
	res := ID{Limit: p.Limit}
	g, ctx := errgroup.WithContext(ctx)
	if p.Before != nil {
		g.Go(func() error {
			id, err := getIDfn(*p.Before)
			if err != nil {
				return err
			}
			res.Before = &id
			return nil
		})
	}
	if p.After != nil {
		g.Go(func() error {
			id, err := getIDfn(*p.After)
			if err != nil {
				return err
			}
			res.After = &id
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return ID{}, err
	}
	return res, nil
}
