package MyTime

import (
	"database/sql/driver"
	"time"
)

const TimeFormat = "2006-01-02 15:04:05"

const (
	SecondsPerMinute = 60
	SecondsPerHour   = 60 * SecondsPerMinute
	SecondsPerDay    = 24 * SecondsPerHour
	SecondsPerWeek   = 7 * SecondsPerDay
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
	// 0001-01-01 00:00:00 属于空值，遇到空值解析成 null 即可
	if t.String() == "0001-01-01 00:00:00" {
		return nil, nil
	}
	return []byte(time.Time(t).Format(TimeFormat)), nil
}

// Scan 检出 mysql 时调用
func (t *Time) Scan(v interface{}) error {
	// mysql 内部日期的格式可能是 2006-01-02 15:04:05 +0800 CST 格式，所以检出的时候还需要进行一次格式化
	tTime, _ := time.Parse(TimeFormat+" +0800 CST", v.(time.Time).String())
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
