package valueobject

import "strconv"

type Query struct {
	Page       int
	PerPage    int
	Order      string
	Conditions []*Condition
	With       []string
	Lock       bool
}

func NewQuery(page string, perPage string, order string) (*Query, error) {
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
		Order:   order,
	}
	return query, nil
}
