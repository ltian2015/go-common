package ranges

import "time"

type TimeInterval = SeqRange[time.Time, time.Time]

func CreateTimeInterval(t1, t2 time.Time) TimeInterval {
	return CreateSeqRange[time.Time, time.Time](t1, t2)
}

func Tintvl2Str(ti TimeInterval) string {
	var layout = "2006-01-02 15:04:05"
	return FmtTintvl(ti, layout)
}
func FmtTintvl(ti TimeInterval, layout string) string {
	var f = func(t time.Time) string {
		return t.Format(layout)
	}
	return RngToStr[time.Time, TimeInterval](ti, f)
}
