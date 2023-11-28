package valueobject

import "strconv"

type Query struct {
	Page       int
	PerPage    int
	OrderedBy  string
	Conditions []*Condition
	With       []string
	Locked     bool
}

func NewQuery(page string, perPage string) (*Query, error) {
	p, err := strconv.Atoi(page)
	if err != nil {
		return nil, err
	}
	pp, err := strconv.Atoi(perPage)
	if err != nil {
		return nil, err
	}
	if p < 1 {
		p = 1
	}
	query := &Query{
		Page:    p,
		PerPage: pp,
	}
	query.Conditions = make([]*Condition, 0)
	return query, nil
}
