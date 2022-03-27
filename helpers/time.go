package helpers

import (
	"math"
	"strings"
	"time"
)

// 获取当前的时间 - 字符串
func GetCurrentDate() string {
	return time.Now().Format("2006-01-02 15:04:05")
}

// 获取当前的时间 - Unix时间戳
func GetCurrentUnix() int64 {
	return time.Now().Unix()
}

// 获取当前的时间 - 毫秒级时间戳
func GetCurrentMilliUnix() int64 {
	return time.Now().UnixNano() / 1000000
}

// 获取当前的时间 - 纳秒级时间戳
func GetCurrentNanoUnix() int64 {
	return time.Now().UnixNano()
}
func GetTimeDate(_format string) (date string) {
	if len(_format) == 0 {
		_format = "YmdHisMS"
	}
	date = _format

	// 时区
	//timeZone, _ := time.LoadLocation(ServerInfo["timezone"])
	timeZone := time.FixedZone("CST", 8*3600) // 东八区

	timer := time.Now().In(timeZone)

	var year int64 = int64(timer.Year())
	var month int64 = int64(timer.Month())
	var day int64 = int64(timer.Day())
	var hour int64 = int64(timer.Hour())
	var minute int64 = int64(timer.Minute())
	var second int64 = int64(timer.Second())
	var week int64 = int64(timer.Weekday())
	var ms int64 = int64(timer.UnixNano() / 1e6)
	var ns int64 = int64(timer.UnixNano() / 1e9)
	msTmp := IntToString(int64(math.Floor(float64(ms / 1000))))
	nsTmp := IntToString(int64(math.Floor(float64(ns / 1000000))))

	var _year string
	var _month string
	var _day string
	var _hour string
	var _minute string
	var _second string
	var _week string // 英文星期
	var _Week string // 中文星期
	var _ms string   // 毫秒
	var _ns string   // 纳秒

	_year = IntToString(year)
	if month < 10 {
		_month = IntToString(month)
		_month = "0" + _month
	} else {
		_month = IntToString(month)
	}
	if day < 10 {
		_day = IntToString(day)
		_day = "0" + _day
	} else {
		_day = IntToString(day)
	}
	if hour < 10 {
		_hour = IntToString(hour)
		_hour = "0" + _hour
	} else {
		_hour = IntToString(hour)
	}
	if minute < 10 {
		_minute = IntToString(minute)
		_minute = "0" + _minute
	} else {
		_minute = IntToString(minute)
	}
	if second < 10 {
		_second = IntToString(second)
		_second = "0" + _second
	} else {
		_second = IntToString(second)
	}
	_week = IntToString(week)
	WeekZh := [...]string{"日", "一", "二", "三", "四", "五", "六"} // 默认从"日"开始
	_Week = WeekZh[week]
	_ms = strings.Replace(IntToString(ms), msTmp, "", -1)
	_ns = strings.Replace(IntToString(ns), nsTmp, "", -1)

	// 替换关键词
	date = strings.Replace(date, "MS", _ms, -1)
	date = strings.Replace(date, "NS", _ns, -1)
	date = strings.Replace(date, "Y", _year, -1)
	date = strings.Replace(date, "m", _month, -1)
	date = strings.Replace(date, "d", _day, -1)
	date = strings.Replace(date, "H", _hour, -1)
	date = strings.Replace(date, "i", _minute, -1)
	date = strings.Replace(date, "s", _second, -1)
	date = strings.Replace(date, "W", _Week, -1)
	date = strings.Replace(date, "w", _week, -1)

	return
}

// DateToTimeS 秒日期时间戳转时间戳，s
func DateToTimeS(_date string, format string) int64 {
	var date string
	if len(_date) == 0 { //给一个默认值
		date = GetTimeDate("YmdHis")
	} else {
		date = _date
	}

	var layout string
	if format == "YmdHis" || format == "" {
		layout = "20060102150405" // 转化所需内定模板
	} else if format == "Y-m-d H:i:s" {
		layout = "2006-01-02 15:04:05"
	} else if format == "Y年m月d日 H:i:s" {
		layout = "2006年01月02日 15:04:05"
	} else {
		layout = "20060102150405"
	}

	//日期转化为时间戳
	loc, _ := time.LoadLocation("Local") //获取时区
	tmp, _ := time.ParseInLocation(layout, date, loc)
	timestamp := tmp.Unix() //转化为时间戳 类型是int64

	return timestamp
}

func GetDatesBetweenDay(startDate string, endDate string, format string) (day int64) {
	// 日期转秒时间戳
	startTime := DateToTimeS(startDate, format)
	endTime := DateToTimeS(endDate, format)
	// 获取日期秒之差
	dayTime := endTime - startTime
	// 时分秒将被忽略，只取天的部分
	day = int64(math.Floor(float64(dayTime / (24 * 60 * 60))))
	return
}
