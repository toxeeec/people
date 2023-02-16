package postgres

import (
	"fmt"
	"strings"

	"github.com/jmoiron/sqlx"
	"github.com/toxeeec/people/backend/pagination"
)

type QueryBuilder struct {
	base   string
	join   string
	group  string
	args   []any
	conds  []string
	having []string
	vals   []string
	end    string
	err    error
}

func NewQuery(base string) *QueryBuilder {
	var q QueryBuilder
	q.base = base
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

func (q *QueryBuilder) Join(table string, cond string) *QueryBuilder {
	q.join = " JOIN " + table + " ON " + cond
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

func (q *QueryBuilder) GroupBy(cols ...string) *QueryBuilder {
	q.group = " GROUP BY " + strings.Join(cols, ", ")
	return q
}

func (q *QueryBuilder) Limit(n uint) *QueryBuilder {
	q.end += fmt.Sprintf(" LIMIT %d", n)
	return q
}

func (q *QueryBuilder) Having(cond string, args ...any) *QueryBuilder {
	args = Flatten(args)
	if strings.Contains(cond, "IN ") {
		var err error
		cond, args, err = sqlx.In(cond, args)
		if err != nil {
			q.err = err
		}
	}
	q.having = append(q.having, cond)
	q.args = append(q.args, args...)
	return q
}

// TODO: simplify (use field only)
func (q *QueryBuilder) Paginate(p pagination.ID, field string, val string, args ...any) *QueryBuilder {
	if p.Before != nil {
		if q.group == "" {
			q.Where(field + " < " + val)
		} else {
			q.Having(field + " < " + val)
		}
		q.args = append(q.args, *p.Before)
		q.args = append(q.args, args...)
	}
	if p.After != nil {
		if q.group == "" {
			q.Where(field + " > " + val)
		} else {
			q.Having(field + " > " + val)
		}
		q.args = append(q.args, *p.After)
		q.args = append(q.args, args...)
	}
	q.OrderBy(false, field)
	q.Limit(p.Limit)
	return q
}

func (q *QueryBuilder) Values(vals ...string) *QueryBuilder {
	q.vals = append(q.vals, vals...)
	return q
}

func (q *QueryBuilder) Build() (string, []any, error) {
	if q.err != nil {
		return "", nil, fmt.Errorf("Query.Build %w", q.err)
	}
	if len(q.vals) > 0 {
		q.base += " VALUES " + strings.Join(q.vals, ", ")
		return sqlx.Rebind(bindType, q.base), q.args, nil
	}
	q.base += q.join
	if len(q.conds) > 0 {
		q.base += " WHERE " + strings.Join(q.conds, " AND ")
	}
	q.base += q.group
	if len(q.having) > 0 {
		q.base += " HAVING " + strings.Join(q.having, " AND ")
	}
	q.base += q.end
	return sqlx.Rebind(bindType, q.base), q.args, nil
}

var bindType = sqlx.BindType("postgres")
