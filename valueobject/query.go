package valueobject

type Query struct {
	Page       int
	PerPage    int
	OrderedBy  string
	Conditions []*Condition
	With       []string
	Locked     bool
}

func NewQuery() *Query {
	query := &Query{}
	query.Conditions = make([]*Condition, 0)
	return query
}

func (q *Query) Condition(field string, operator Operator, value any) *Query {
	condition := NewCondition(field, operator, value)
	q.Conditions = append(q.Conditions, condition)
	return q
}

func (q *Query) Paginate(page int, perPage int) *Query {
	q.Page = page
	q.PerPage = perPage
	return q
}

func (q *Query) Order(orderedBy string) *Query {
	q.OrderedBy = orderedBy
	return q
}

func (q *Query) Lock() *Query {
	q.Locked = true
	return q
}
