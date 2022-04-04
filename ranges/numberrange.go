package ranges

import "fmt"

func CreateNumberRange[P number](p1, p2 P) NumberRange[P] {
	if p1 <= p2 {
		return NumberRange[P]{start: p1, end: p2}
	} else {
		return NumberRange[P]{p2, p1}
	}

}

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

func (nr NumberRange[P]) IsPoint() bool {
	return nr.start == nr.end
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

func (nr NumberRange[P]) String() string {
	return fmt.Sprintf("NumberRange[%v,%v)", nr.start, nr.end)
}

func (nr NumberRange[P]) IsIntersected(other NumberRange[P]) bool {
	var rgThis Range[P, NumberRange[P]] = nr
	var rgOther Range[P, NumberRange[P]] = other
	return IsIntersected(rgThis, rgOther)
}

func (nr NumberRange[P]) Intersect(other NumberRange[P]) (bool, NumberRange[P]) {
	var rgThis Range[P, NumberRange[P]] = nr
	var rgOther Range[P, NumberRange[P]] = other
	return Intersect(rgThis, rgOther)
}

func (nr NumberRange[P]) IntersectOthers(others []NumberRange[P]) (bool, NumberRange[P]) {
	/*
		var result NumberRange[P] = nr
		var isAllIntersected, intersected bool = true, true
		if len(others) == 0 {
			return isAllIntersected, result
		}
		for _, other := range others {
			intersected, result = result.Intersect(other)
			isAllIntersected = isAllIntersected && intersected
		}
		return isAllIntersected, result
	*/
	var rgThis Range[P, NumberRange[P]] = nr
	return IntersectOthers(rgThis, others)
}

func (nr NumberRange[P]) Union(other NumberRange[P]) (bool, NumberRange[P]) {
	var rgThis Range[P, NumberRange[P]] = nr
	var rgOther Range[P, NumberRange[P]] = other
	return Union(rgThis, rgOther)
}

func (nr NumberRange[P]) UnionOthers(others []NumberRange[P]) (bool, NumberRange[P]) {
	/*
		var result NumberRange[P] = nr
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
	*/
	var rgThis Range[P, NumberRange[P]] = nr
	return UnionOthers(rgThis, others)

}
func (nr NumberRange[P]) IsBefore(other NumberRange[P]) bool {
	var rgThis Range[P, NumberRange[P]] = nr
	var rgOther Range[P, NumberRange[P]] = other
	return IsBefore(rgThis, rgOther)
}
func (nr NumberRange[P]) IsAfter(other NumberRange[P]) bool {
	var rgThis Range[P, NumberRange[P]] = nr
	var rgOther Range[P, NumberRange[P]] = other
	return IsAfter(rgThis, rgOther)
}
