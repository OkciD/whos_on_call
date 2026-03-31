package duration

import (
	"encoding/json"
	"fmt"
	"time"
)

type MarshallableDuration struct {
	time.Duration
}

func (d MarshallableDuration) MarshalJSON() ([]byte, error) {
	return json.Marshal(d.String())
}

func (d *MarshallableDuration) UnmarshalJSON(b []byte) error {
	var v any
	if err := json.Unmarshal(b, &v); err != nil {
		return err
	}
	switch value := v.(type) {
	case float64:
		d.Duration = time.Duration(value)
		return nil
	case string:
		var err error
		d.Duration, err = time.ParseDuration(value)
		if err != nil {
			return err
		}
		return nil
	default:
		return fmt.Errorf("failed to parse duration %v", v)
	}
}
