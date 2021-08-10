package gtools

import (
	"time"
)

var Time timeInterface
var TimeFormat = struct {
	DefaultDatetime string
	DefaultDate     string
}{
	DefaultDatetime: "2006-01-02 15:04:05",
	DefaultDate:     "2006-01-02",
}

type timeInterface interface {
	GetDay(day int) string
	GetTime() time.Time
	TimeToUnix(t time.Time) int64
	TimeToMills(t time.Time) int64
	UnixToTime(t1 int64) time.Time
	TimeToString(t time.Time, timeFormat string) string
	StringToTime(timestring string, format string) (time.Time, error)
	Compare(t1, t2 time.Time) bool
	NextHourTime(s string, n int64, timeFormat string) string
	HourDiffer(startTime, endTime time.Time) int64
	Checkhours() bool
	TimeAddDate(years int, months int, days int) time.Time
}

type selfTime struct {
}

func init() {
	Time = &selfTime{}
}

func (_time selfTime) TimeAddDate(years int, months int, days int) time.Time {
	t := time.Now()                        // 获取当前时间
	date := t.AddDate(years, months, days) // 三个参数分别是年月日，此处获取的是前一天的日期
	return date
}

/**
 * @description: 获取指定天后的日期
 * @param {*}
 * @return {*}
 */
func (_time selfTime) GetDay(day int) string {

	// GetTimeByInt
	t := time.Now().Unix()
	t = t + int64(day*86400)
	nowDay := _time.UnixToTime(t).Format("2006-01-02")

	timeResult, err := _time.StringToTime(nowDay, "2006-01-02")
	if err != nil {
		return nowDay
	}
	return timeResult.Format("2006-01-02")
}

/**
 * @description: 当前时间
 * @param {*}
 * @return {*}
 */
func (_time selfTime) GetTime() time.Time {
	return time.Now()
}

/**
 * @description: 将时间转换成时间戳
 * @param {time.Time} t
 * @return {*}
 */
func (_time selfTime) TimeToUnix(t time.Time) int64 {
	return t.Unix()
}

/**
 * @description: 将秒级时间戳转换成毫秒
 * @param {time.Time} t
 * @return {*}
 */
func (_time selfTime) TimeToMills(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

/**
 * @description: 将时间戳转换成时间
 * @param {int64} t1
 * @return {*}
 */
func (_time selfTime) UnixToTime(t1 int64) time.Time {
	return time.Unix(t1, 0)
}

/**
 * @description: 将字符串转换成时间格式
 * @param {string} timestring
 * @param {string} timeFormat 字符串的格式
 * @return {*}
 */
func (_time selfTime) TimeToString(t time.Time, timeFormat string) string {
	if timeFormat == "" {
		timeFormat = TimeFormat.DefaultDatetime
	}
	return t.Format(timeFormat)
}

/**
 * @description: 标准字符串转时间
 * @param {string} timestring
 * @param {string} normalTimeFormat
 * @return {*}
 */
func (_time selfTime) StringToTime(timestring string, format string) (time.Time, error) {
	if timestring == "" {
		return time.Time{}, nil
	}
	return time.ParseInLocation(format, timestring, time.Local)
}

/**
 * @description: 比较两个时间大小
 * @param {*} t1
 * @param {time.Time} t2
 * @return {*}
 */
func (_time selfTime) Compare(t1, t2 time.Time) bool {
	return t1.Unix() > t2.Unix()
}

/**
 * @description: 获取基于当前时间N小时候的时间字符串
 * @param {string} s
 * @param {int64} n
 * @param {string} timeFormat
 * @return {*}
 */
func (_time selfTime) NextHourTime(s string, n int64, timeFormat string) string {
	t2, _ := time.ParseInLocation(timeFormat, s, time.Local)
	t1 := t2.Add(time.Hour * time.Duration(n))
	return _time.TimeToString(t1, timeFormat)
}

/**
 * @description: 计算俩个时间差多少小时
 * @param {*} startTime
 * @param {string} endTime
 * @param {string} timeFormat
 * @return {*}
 */
func (_time selfTime) HourDiffer(startTime time.Time, endTime time.Time) int64 {
	var hour int64
	hour = ((startTime.Unix() - endTime.Unix()) / 60) / 60
	return hour
}

/**
 * @description: 判断当前时间是否是整点
 * @param {*}
 * @return {*}
 */
func (_time selfTime) Checkhours() bool {
	_, m, s := _time.GetTime().Clock()
	if m == s && m == 0 && s == 0 {
		return true
	}
	return false
}
