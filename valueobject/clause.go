package valueobject

import "strconv"

type Clause struct {
	Page    int
	PerPage int
	Order   string
}

func NewClause(page string, perPage string, order string) (*Clause, error) {
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
	clause := &Clause{
		Page:    p,
		PerPage: pp,
		Order:   order,
	}
	return clause, nil
}
