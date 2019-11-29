package utils

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

func (this *MyTime) UnMarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	t,err := time.Parse("2006-01-02 15:04:05",string(data))
	*this = MyTime(t)
	return err
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

func InNight() bool {
	hour := time.Now().Hour()
	return (hour >= 0 && hour < 6) || (hour >= 19)
}