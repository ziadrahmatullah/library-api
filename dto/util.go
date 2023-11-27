package dto

import "time"

func newResponsesFromEntities[T1 any, T2 any](t []*T1, mapper func(*T1) *T2) []*T2 {
	responses := make([]*T2, 0)
	for _, t1 := range t {
		responses = append(responses, mapper(t1))
	}
	return responses
}

func toDateString(date time.Time) string {
	return date.Format("2006-01-02")
}
