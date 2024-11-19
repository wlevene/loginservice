package util

import "time"

func TimeStamp2Time(t int64) time.Time {
	return time.Unix(t, 0)
}

func Time2TimeStamp(t time.Time) int64 {
	return t.Unix()
}

func Time2String(t *time.Time) string {
	return t.String()
}

func String2Time(t string) time.Time {
	layout := "2006-01-02 15:04:05.999999999 -0700 MST"
	ret_time, _ := time.Parse(layout, t)
	return ret_time
}
