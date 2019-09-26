package common

import (
	"fmt"
	"github.com/go-sql-driver/mysql"
	"time"
)

type MyTime time.Time

func (t MyTime) String() string {
	return time.Time(t).Format("2006-01-02 15:04:05")
}

func (this MyTime) MarshalJSON() ([]byte, error) {
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}

type MyNullTime mysql.NullTime

func (this MyNullTime) String() string {
	return time.Time(this.Time).Format("2006-01-02 15:04:05")
}

func (this MyNullTime) MarshalJSON() ([]byte, error) {
	if !this.Valid {
		return []byte("null"), nil
	}
	var stamp = fmt.Sprintf("\"%s\"", time.Time(this.Time).Format("2006-01-02 15:04:05"))
	return []byte(stamp), nil
}
