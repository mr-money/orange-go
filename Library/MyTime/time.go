package MyTime

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

const (
	Nanosecond       time.Duration = 1
	Microsecond                    = 1000 * Nanosecond
	Millisecond                    = 1000 * Microsecond
	Second                         = 1000 * Millisecond
	SecondsPerMinute               = 60 * Second
	SecondsPerHour                 = 60 * SecondsPerMinute
	SecondsPerDay                  = 24 * SecondsPerHour
)

type Time time.Time

func (t *Time) UnmarshalJSON(data []byte) (err error) {
	// 空值不进行解析
	if len(data) == 2 {
		*t = Time(time.Time{})
		return
	}

	// 指定解析的格式
	now, err := time.Parse(`"`+TimeFormat+`"`, string(data))
	*t = Time(now)
	return
}

func (t Time) MarshalJSON() ([]byte, error) {
	b := make([]byte, 0, len(TimeFormat)+2)
	b = append(b, '"')
	b = time.Time(t).AppendFormat(b, TimeFormat)
	b = append(b, '"')
	return b, nil
}

// Value 写入 mysql 时调用
func (t Time) Value() (driver.Value, error) {
	if time.Time(t).IsZero() {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// Scan 检出 mysql 时调用
func (t *Time) Scan(v interface{}) error {
	tValue, _ := v.(time.Time)

	location, _ := time.LoadLocation("Asia/Shanghai")

	tTime, _ := time.ParseInLocation(
		TimeFormat,
		tValue.Format(TimeFormat),
		location,
	)

	*t = Time(tTime)

	return nil
}

// 用于 fmt.Println 和后续验证场景
func (t Time) String() string {
	return time.Time(t).Format(TimeFormat)
}

//
// StrToTime
// @Description: 时间字符串转Time
// @param value
// @return Time
//
func StrToTime(value string) Time {
	t, err := time.Parse(TimeFormat, value)
	if err != nil {
		return Time{}
	}

	return Time(t)
}
