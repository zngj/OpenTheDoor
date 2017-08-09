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