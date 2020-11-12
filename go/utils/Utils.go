package utils

import "time"

var location *time.Location

func init() {
	l, _ := time.LoadLocation("Asia/Shanghai")
	location = l
}

func GetNowStr() string {

	in := time.Now().In(location)
	timeFormat := in.Format("2006-01-02 15:04:05")
	return timeFormat
}

func String2Time(timeStr string) time.Time {
	tempTime, err := time.ParseInLocation("2006-01-02 15:04:05", timeStr, location)

	panic(err)

	return tempTime
}
