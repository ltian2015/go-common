package ranges

import (
	"fmt"
	"time"
)

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

//CreateSeqRange函数用于创建一个
//在类型上，P==T
func CreateSeqRange[P Sequencable[T], T any](p1, p2 P) SeqRange[P, T] {
	t2 := typeTo[P, T](p2)
	if p1.Before(t2) || p1.Equal(t2) {
		return SeqRange[P, T]{start: p1, end: p2}
	} else {
		return SeqRange[P, T]{start: p2, end: p1}
	}
}

//这里的S，D在类型上永远相等，所以typeTo是安全的。
func typeTo[S, D any](s S) D {
	return (interface{})(s).(D)
}

type SeqRange[P Sequencable[T], T any] struct {
	start P
	end   P
}

func (sr SeqRange[P, T]) Range(start, end P) SeqRange[P, T] {
	return CreateSeqRange[P, T](start, end)
}
func (sr SeqRange[P, T]) DeRange() (start, end P) {
	return sr.start, sr.end
}
func (sr SeqRange[P, T]) IsPoint() bool {
	var endT T = typeTo[P, T](sr.end)
	return sr.start.Equal(endT)
}
func (sr SeqRange[P, T]) IsIncludedPoint(p P) bool {
	var pt T = typeTo[P, T](p)
	return (sr.start.Before(pt) || sr.start.Equal(pt)) && sr.end.After(pt)
}
func (sr SeqRange[P, T]) IsBeforePoint(p P) bool {
	var pt T = typeTo[P, T](p)
	return sr.end.Equal(pt) || sr.end.Before(pt)
}
func (sr SeqRange[P, T]) IsAfterPoint(p P) bool {
	var pt T = typeTo[P, T](p)
	return sr.start.After(pt)
}

func (sr SeqRange[P, T]) String() string {
	return fmt.Sprintf("NumberRange[%v,%v)", sr.start, sr.end)
}

func (sr SeqRange[P, T]) IsIntersected(other SeqRange[P, T]) bool {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return IsIntersected(rgThis, rgOther)
}

func (sr SeqRange[P, T]) Intersect(other SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return Intersect(rgThis, rgOther)
}

func (sr SeqRange[P, T]) IntersectOthers(others []SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var result SeqRange[P, T] = sr
	var isAllIntersected, intersected bool = true, true
	if len(others) == 0 {
		return isAllIntersected, result
	}
	for _, other := range others {
		intersected, result = result.Intersect(other)
		isAllIntersected = isAllIntersected && intersected
	}
	return isAllIntersected, result
}

func (sr SeqRange[P, T]) Union(other SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return Union(rgThis, rgOther)
}

func (sr SeqRange[P, T]) UnionOthers(others []SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var result SeqRange[P, T] = sr
	var isAllSuccessive, successived bool = true, true
	if len(others) == 0 {
		return true, result
	}
	for _, other := range others {
		successived, result = result.Union(other)
		if successived == false {
			isAllSuccessive = false
		}
	}
	return isAllSuccessive, result
}
func (sr SeqRange[P, T]) IsBefore(other SeqRange[P, T]) bool {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return IsBefore(rgThis, rgOther)
}
func (sr SeqRange[P, T]) IsAfter(other SeqRange[P, T]) bool {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return IsAfter(rgThis, rgOther)
}
