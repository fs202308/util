package xos

import "time"

func TodayWithFormat(format string) *time.Time {
	t, err := time.Parse(format, time.Now().Format(format))
	if nil != err {
		t = time.Time{}
		return &t
	}

	return &t
}
