package model

import (
	"database/sql/driver"
	"fmt"
	_ "github.com/lib/pq"
	"time"
)

//时间类型
type Time struct {
	time.Time
}

// MarshalJSON 序列化为JSON
func (t Time) MarshalJSON() ([]byte, error) {

	if t.IsZero() {
		return []byte("\"\""), nil
	}
	stamp := fmt.Sprintf("\"%s\"", t.Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

// UnmarshalJSON反序列化为JSON
func (t *Time) UnmarshalJSON(data []byte) error {

	var err error
	t.Time, err = time.Parse("2006-01-02 15:04:05", string(data)[1:20])
	return err
}

//读取插入
func (t *Time) Scan(v interface{}) error {

	value, ok := v.(time.Time)

	if ok {
		*t = Time{Time: value}
		return nil
	}
	return fmt.Errorf("can not convert %v to timestamp", v)
}

//插入
func (t Time) Value() (driver.Value, error) {
	str := t.Format("2006-01-02 15:04:05")
	if str == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return str, nil
}
