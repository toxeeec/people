package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/toxeeec/people/backend/pagination"
)

type QueryBuilder struct {
	base  string
	join  string
	conds []string
	args  []any
	end   string
	err   error
}

func NewQuery(base string) *QueryBuilder {
	var q QueryBuilder
	q.base = base
	q.conds = make([]string, 1, 5)
	q.args = make([]any, 0, 5)
	// required if no conditions
	q.conds[0] = "1 = 1"
	return &q
}

func (q *QueryBuilder) Where(cond string, args ...any) *QueryBuilder {
	args = Flatten(args)
	if strings.Contains(cond, "IN ") {
		var err error
		cond, args, err = sqlx.In(cond, args)
		if err != nil {
			q.err = err
		}
	}
	q.conds = append(q.conds, cond)
	q.args = append(q.args, args...)
	return q
}

func (q *QueryBuilder) Join(field string, cond string) *QueryBuilder {
	q.join = " JOIN " + field + " ON " + cond
	return q
}

func (q *QueryBuilder) OrderBy(asc bool, cols ...string) *QueryBuilder {
	q.end = " ORDER BY " + strings.Join(cols, ", ")
	if asc {
		q.end += " ASC"
	} else {
		q.end += " DESC"
	}
	return q
}

func (q *QueryBuilder) Limit(n uint) *QueryBuilder {
	q.end += fmt.Sprintf(" LIMIT %d", n)
	return q
}

func (q *QueryBuilder) Paginate(p pagination.ID, field string, val string, args ...any) *QueryBuilder {
	if p.Before != nil {
		q.Where(field + " < " + val)
		q.args = append(q.args, *p.Before)
		q.args = append(q.args, args...)
	}
	if p.After != nil {
		q.Where(field + " > " + val)
		q.args = append(q.args, *p.After)
		q.args = append(q.args, args...)
	}
	q.OrderBy(false, field)
	q.Limit(p.Limit)
	return q
}

func (q *QueryBuilder) Build() (string, []any, error) {
	if q.err != nil {
		return "", nil, fmt.Errorf("Query.Build %w", q.err)
	}
	q.base += q.join
	q.base += " WHERE " + strings.Join(q.conds, " AND ") + q.end
	return sqlx.Rebind(bindType, q.base), q.args, nil
}

var bindType = sqlx.BindType("postgres")
