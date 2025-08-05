package MyTime

import (
	"database/sql/driver"
	"strconv"
	"strings"
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

// StrToTime
// @Description: 时间字符串转Time
// @param value
// @return Time
func StrToTime(value string) Time {
	t, err := time.ParseInLocation(TimeFormat, value, time.Local)
	if err != nil {
		return Time{}
	}

	return Time(t)
}

// Format
// @Description: 格式化时间
// @receiver t
// @param layout 时间格式 如 2006-01-02 15:04:05
// @return string
func (t Time) Format(layout string) string {
	return time.Time(t).Format(layout)
}

// ParseToSecond
// @Description: 时间字符串转int
// @param timeStr
// @return int
func ParseToSecond(timeStr string) int {
	if timeStr == "" {
		return 0
	}
	parts := strings.Split(timeStr, ":")
	switch len(parts) {
	case 3: // HH:MM:SS
		hours, err1 := strconv.Atoi(parts[0])
		minutes, err2 := strconv.Atoi(parts[1])
		seconds, err3 := strconv.Atoi(parts[2])
		if err1 != nil || err2 != nil || err3 != nil {
			return 0 // 如果转换失败，返回默认值
		}
		return hours*3600 + minutes*60 + seconds
	case 2: // MM:SS
		minutes, err1 := strconv.Atoi(parts[0])
		seconds, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			return 0 // 如果转换失败，返回默认值
		}
		return minutes*60 + seconds
	default:
		return 0 // 如果格式不正确，返回默认值
	}
}
