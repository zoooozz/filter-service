package time

import (
	"database/sql/driver"
	"strconv"
	xtime "time"
)

type Time int64

func (jt *Time) Scan(src interface{}) (err error) {
	switch sc := src.(type) {
	case xtime.Time:
		*jt = Time(sc.Unix())
	case string:
		var i int64
		i, err = strconv.ParseInt(sc, 10, 64)
		*jt = Time(i)
	}
	return
}

func (jt Time) Value() (driver.Value, error) {
	return xtime.Unix(int64(jt), 0), nil
}

func (jt Time) Time() xtime.Time {
	return xtime.Unix(int64(jt), 0)
}

type Duration xtime.Duration

func (d *Duration) UnmarshalText(text []byte) error {
	tmp, err := xtime.ParseDuration(string(text))
	if err == nil {
		*d = Duration(tmp)
	}
	return err
}
