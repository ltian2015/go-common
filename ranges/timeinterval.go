package ranges

import "time"

type TimeInterval = SeqRange[time.Time, time.Time]
type MyTime time.Time

func (mt MyTime) Before(other MyTime) bool {
	return mt.Before(other)
}
func (mt MyTime) After(other MyTime) bool {
	return mt.After(other)
}
func (mt MyTime) Equal(other MyTime) bool {
	return mt.Equal(other)
}

type MyTimeInterval = SeqRange[MyTime, MyTime]
