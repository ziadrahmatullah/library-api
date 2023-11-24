package valueobject

type Operator string

const (
	Equal Operator = "="
	Ilike          = "ILIKE"
)

type Condition struct {
	Field     string
	Operation Operator
	Value     string
}

func NewCondition(field string, operation Operator, value string) *Condition {
	if value == "" {
		return nil
	}
	if operation == Ilike {
		value = "%" + value + "%"
	}
	return &Condition{
		Field:     field,
		Operation: operation,
		Value:     value,
	}
}
