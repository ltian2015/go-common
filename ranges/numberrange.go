package ranges

type number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 | ~float32 | ~float64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64
}

//CreateNumberRange函数用于给定的两个点创建一个区间，
//无论两个点的大小顺序如何，创建出来的区间的起点都会小于终点。
func CreateNumberRange[P number](p1, p2 P) NumberRange[P] {
	if p1 <= p2 {
		return NumberRange[P]{start: p1, end: p2}
	} else {
		return NumberRange[P]{p2, p1}
	}

}

//NumberRange[P number] 定义了各种数字类型元素组成的区间类型，该类型的区间
//满足Range[P, NumberRange[P]接口。
type NumberRange[P number] struct {
	start P
	end   P
}

func (nr NumberRange[P]) Range(start, end P) NumberRange[P] {
	result := CreateNumberRange(start, end)
	return result
}

func (nr NumberRange[P]) DeRange() (start, end P) {
	return nr.start, nr.end
}

func (nr NumberRange[P]) IsIncludedPoint(p P) bool {
	return p >= nr.start && p < nr.end
}
func (nr NumberRange[P]) IsBeforePoint(p P) bool {
	return p >= nr.end
}
func (nr NumberRange[P]) IsAfterPoint(p P) bool {
	return p < nr.start
}

////////////////////////////////////////////////////////////////////
func (nr NumberRange[P]) IsPoint() bool {
	return IsPoint[P, NumberRange[P]](nr)
}
func (nr NumberRange[P]) String() string {
	return RngToStr[P, NumberRange[P]](nr, v2s[P])
}
func (nr NumberRange[P]) Equal(other NumberRange[P]) bool {
	return Equal[P, NumberRange[P]](nr, other)
}
func (nr NumberRange[P]) IsIntersected(other NumberRange[P]) bool {
	return IsIntersected[P, NumberRange[P]](nr, other)
}

func (nr NumberRange[P]) Intersect(other NumberRange[P]) (bool, NumberRange[P]) {
	return Intersect[P, NumberRange[P]](nr, other)
}

func (nr NumberRange[P]) IntersectOthers(others []NumberRange[P]) (bool, NumberRange[P]) {
	return IntersectOthers[P, NumberRange[P]](nr, others)
}

func (nr NumberRange[P]) Union(other NumberRange[P]) (bool, NumberRange[P]) {
	return Union[P, NumberRange[P]](nr, other)
}

func (nr NumberRange[P]) UnionOthers(others []NumberRange[P]) (bool, NumberRange[P]) {
	return UnionOthers[P, NumberRange[P]](nr, others)

}
func (nr NumberRange[P]) Except(other NumberRange[P]) (r1, r2 NumberRange[P]) {
	return Except[P, NumberRange[P]](nr, other)
}
func (nr NumberRange[P]) IsBefore(other NumberRange[P]) bool {
	return IsBefore[P, NumberRange[P]](nr, other)
}
func (nr NumberRange[P]) IsAfter(other NumberRange[P]) bool {

	return IsAfter[P, NumberRange[P]](nr, other)
}

///////////////////////////////////////////////////////////////////////////////

type NumCycle[P number] struct {
	Count int
	Unit  P
}

func (nc NumCycle[P]) GetCount() int {
	return nc.Count
}
func (nc NumCycle[P]) GetUnit() P {
	return nc.Unit
}

//
type NRCycleFunc[P number] struct{}

func (nc *NRCycleFunc[P]) OfCycles(t NumberRange[P], n int, c NumCycle[P]) NumberRange[P] {
	rStart, rEnd := t.DeRange()
	var mount = P(n) * P(c.Count) * c.Unit
	start := rStart + mount
	end := rEnd + mount
	return t.Range(start, end)
}

//
type NPCycleFunc[P number] struct{}

func (nc *NPCycleFunc[P]) OfCycles(t P, n int, c NumCycle[P]) P {
	result := t + P(n)*P(c.Count)*c.Unit
	return result
} //
