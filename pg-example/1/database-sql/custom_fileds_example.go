package main

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

type MyCustomType struct {
	Number int
	Valid  bool
}

// наш тип MyCustomType удовлетовряет интерфейсу sql.Scanner
var _ sql.Scanner = (*MyCustomType)(nil)

// Scan implements the Scanner interface.
func (n *MyCustomType) Scan(src interface{}) error {
	// The src value will be of one of the following types:
	//
	//    int64
	//    float64
	//    bool
	//    []byte
	//    string
	//    time.Time
	//    nil - for NULL values
	if src == nil {
		n.Number, n.Valid = 0, false
		return nil
	}
	n.Valid = true

	// some fantastic logic here
	switch src := src.(type) {
	case int64:
		n.Number = int(src)
	case bool:
		n.Number = 1
	default:
		return fmt.Errorf("can't scan %#v into MyCustomType", src)
	}

	return nil
}

// Если мы хотим наш тип как-то хитро мапить в null/value, то реализуем Value()

// наш тип MyCustomType удовлетовряет интерфейсу driver.Valuer
var _ driver.Valuer = (*MyCustomType)(nil)

// Value implements the driver Valuer interface.
func (n MyCustomType) Value() (driver.Value, error) {
	if !n.Valid {
		return nil, nil
	}
	return int64(n.Number), nil
}
