package util

import "time"

func TimeNow() *time.Time {
	t := time.Now()
	return &t
}

//func TimeFromUnixMilli(millisec int64) time.Time {
//	return time.Unix(0, millisec*int64(time.Millisecond))
//}
//
//func TimeToUnixMilli(t time.Time) int64 {
//	return t.UnixNano() / int64(time.Millisecond)
//}

func GetTodayInterval() (begin, end time.Time) {
	now := time.Now()
	begin = time.Date(now.Year(), now.Month(), now.Day(), 0 ,0,0 ,0, time.Local)
	end = begin.AddDate(0, 0, 1)
	return
}

func NowDate() time.Time {
	now := time.Now()
	return time.Date(now.Year(), now.Month(), now.Day(), 0 ,0,0 ,0, time.Local)
}