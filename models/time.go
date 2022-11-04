package models

/**
  @author:pandi
  @date:2022-11-03
  @note:
**/
import (
	"database/sql/driver"
	"encoding/json"
	"time"
)

var marshalFormat = "2006-01-02 15:04:05"

type MyTime time.Time

func (t MyTime) Value() (driver.Value, error) {
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(marshalFormat)), nil
}

func (mt *MyTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation(marshalFormat, s, time.UTC)
	if err != nil {
		return err
	}
	*mt = MyTime(t)
	return nil
}

func (mt *MyTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(*mt).Format(marshalFormat))
}

func (mt MyTime) String() string {
	return time.Time(mt).Format(marshalFormat)
}
