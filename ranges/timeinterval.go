package ranges

import "time"

type TimeInterval = SeqRange[time.Time, time.Time]

func CreateTimeInterval(t1, t2 time.Time) TimeInterval {
	return CreateSeqRange[time.Time, time.Time](t1, t2)
}

/**
由于TimeInterval是由泛型类型SeqRange实例化产生的类型，
go.18版本不允许有分型类型实例化所产生的类型定义自己的方法，比如：

func (tintvl TimeInterval) String() string

所以，定义若干函数来操作TimeInterval
**/
// Tintvl2Str函数使用格式"YYYY-MM-DD hh:mm:ss" 来格式化时间段的起止时间
const TIME_LAYOUT_SECOND = "2006-01-02 15:04:05"

func Tintvl2Str(ti TimeInterval) string {
	return FmtTintvl(ti, TIME_LAYOUT_SECOND)
}

// FmtTintvl函数使用给定的格式参数layout来格式化时间段的起止时间

func FmtTintvl(ti TimeInterval, layout string) string {
	var f = func(t time.Time) string {
		return t.Format(layout)
	}
	return RngToStr[time.Time, TimeInterval](ti, f)
}

//TimeCycle定义了时间周期类型
type TimeCycle struct {
	Count int
	Unit  time.Duration
}

func (tc TimeCycle) GetCount() int {
	return tc.Count
}
func (tc TimeCycle) GetUnit() time.Duration {
	return tc.Unit
}

//TPCycleFunc是时间间隔（TimeInterval）的周期计算函数
type TICycleFunc struct{}

func (nc *TICycleFunc) OfCycles(t TimeInterval, n int, c TimeCycle) TimeInterval {
	rStart, rEnd := t.DeRange()
	var duration = time.Duration(n * c.GetCount() * int(c.GetUnit()))
	start := rStart.Add(duration)
	end := rEnd.Add(duration)
	return t.Range(start, end)
}

//TPCycleFunc是时间点（Time Point）的周期计算函数
type TPCycleFunc struct{}

func (tp *TPCycleFunc) OfCycles(t time.Time, n int, c TimeCycle) time.Time {
	var duration = time.Duration(n * c.GetCount() * int(c.GetUnit()))
	result := t.Add(duration)
	return result
} //
