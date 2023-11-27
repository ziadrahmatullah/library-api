package valueobject

import (
	"database/sql"
	"encoding/json"
)

type NullTime struct {
	sql.NullTime
}

func (n NullTime) MarshalJSON() ([]byte, error) {
	if n.Valid {
		return json.Marshal(n.Time.Format("2006-01-02"))
	}
	return []byte("null"), nil
}
