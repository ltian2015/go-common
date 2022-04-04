package ranges

import "fmt"

//Range是指是由类型参数P的值作为起点和终点（不包括在内）的左闭右开区间。
//数据结构形如，struct {start P,end P},数学表示形式如， [start,end)
//从数学上讲，start永远小于end，end永远在start之后，如果二者相等，区间实际上就是一个点。
//Range接口中，类型参数P是构成Range起点与终点的类型,而类型参数R则Range自身类型。
//也就是二者类型完全相同,R==Range[P,R]。
//这样写，是由于GO不支持嵌套的类型参数定义，无法写Range[P comparable,Range[any]]，
//所有用R来代替Range[P,R]，但是泛型类型不允许直接使用类型断言,比如，r T=P.(T)
//因而，对于任何满足Range[P，R]的具体类型的值rp，则 r R=(interface{})(rp).(R)都是安全的类型转换。
//反之，对于任何满足R[P]的具体类型的值r，则 rp Range[P,R]=(interface{})(r).(Range[P,R])都是安全的类型转换。
type Range[P comparable, R any] interface {
	//以下方法需要实现者依靠自己去实现
	Range(start, end P) R
	DeRange() (start, end P)
	IsIncludedPoint(p P) bool
	IsBeforePoint(p P) bool
	IsAfterPoint(p P) bool
	IsPoint() bool
	String() string
	//以下方法可以借助ranges包提供相应函数，帮助具体类型简化这些方法的实现逻辑。
	Union(other R) (bool, R)
	UnionOthers(others []R) (bool, R)
	IsIntersected(other R) bool
	Intersect(other R) (bool, R)
	IntersectOthers(other []R) (bool, R)
	IsBefore(other R) bool
	IsAfter(other R) bool
}

//typeTo函数将给定源类型S的值s，转换目标类型D的值。
//由于range包中，需要转换的两个参数类型实际应用中都会被赋予相同的类型，所以使用typeTo是安全的。
func typeTo[S, D any](s S) D {
	return (interface{})(s).(D)
}

///////////////////下面是ranges包提供的辅助函数,可以帮助接口的实现者快速实现功能/////////////////////////
//IsIntersected函数判断this h与other是否相交
//两个区间相交，那么必须有
func IsIntersected[P comparable, R any](this, other Range[P, R]) bool {
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	isIntersected := (this.IsIncludedPoint(otherStart) || this.IsIncludedPoint(otherEnd) ||
		other.IsIncludedPoint(thisStart) || other.IsIncludedPoint(thisEnd)) &&
		thisStart != otherEnd && otherStart != thisEnd
	return isIntersected
}

//IsPoint函数判断给定的区间值r是否是一个点。
//如果给定区间值r的起点与终点相等，则返回true,否则返回false
func IsPoint[P comparable, R any](r Range[P, R]) bool {
	start, end := r.DeRange()
	return start == end
}

//Equql函数判断this与other是否相等。
func Equql[P comparable, R any](this, other Range[P, R]) bool {
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	return thisStart == otherStart && thisEnd == otherEnd

}
func IsIncluded[P comparable, R any](this, other Range[P, R]) bool {
	otherStart, otherEnd := other.DeRange()
	_, thisEnd := this.DeRange()
	var result bool = this.IsIncludedPoint(otherStart) && (this.IsIncludedPoint(otherEnd) || thisEnd == otherEnd)
	return result
}
func Intersect[P comparable, R any](this, other Range[P, R]) (bool, R) {
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	isIntersected := IsIntersected(this, other)
	var start, end P // 零值
	if !isIntersected {

		return isIntersected, this.Range(start, end)
	}
	start, end = thisStart, thisEnd
	if this.IsIncludedPoint(otherStart) {
		start = otherStart
	}
	if this.IsIncludedPoint(otherEnd) {
		end = otherEnd
	}
	return isIntersected, this.Range(start, end)
}
func IntersectOthers[P comparable, R any](this Range[P, R], others []R) (bool, R) {
	var existIntersection bool = false
	var result R
	if len(others) == 0 {
		return false, result
	}
	var intersectResult Range[P, R] = this
	var isIntersected bool
	for _, r := range others {
		isIntersected, result = Intersect(intersectResult, typeTo[R, Range[P, R]](r))
		if isIntersected {
			existIntersection = true
		}
		intersectResult = typeTo[R, Range[P, R]](result)
	}
	result = typeTo[Range[P, R], R](intersectResult)
	return existIntersection, result
}

func Union[P comparable, R any](this, other Range[P, R]) (bool, R) {
	isIntersected := IsIntersected(this, other)
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	isSuccessive := isIntersected || thisStart == otherEnd || thisEnd == otherStart
	start, end := thisStart, thisEnd
	if this.IsAfterPoint(otherStart) {
		start = otherStart
	}
	if this.IsBeforePoint(otherEnd) {
		end = otherEnd
	}
	return isSuccessive, this.Range(start, end)
}
func UnionOthers[P comparable, R any](this Range[P, R], others []R) (bool, R) {
	var isAllSuccessive bool = true
	var result R = typeTo[Range[P, R], R](this)
	if len(others) == 0 {
		return isAllSuccessive, result
	}
	var unionResult Range[P, R] = this
	var isSuccessive bool
	for _, r := range others {
		isSuccessive, result = Union(unionResult, typeTo[R, Range[P, R]](r))
		if !isSuccessive {
			isAllSuccessive = false
		}
		unionResult = typeTo[R, Range[P, R]](result)
	}
	return isAllSuccessive, result
}

//Except函数求this与other差集，也就是this中去掉other内容的剩余部分
//返回结果最多是两个区间，r1和r2。
//当this完全包含了other，且没有端点重合的时候。r1和r2都是非零值区间。
//如果两个区间相等，或者this被other所包含，则结果区间r1和r2都是零值。
//其他情况下，结果r1非零值，此时r2为零值。
func Except[P comparable, R any](this, other Range[P, R]) (r1, r2 R) {
	//r1,r2都是空值
	// 如果二者不相交，则，range1区间是自身，second是零值区间
	if !IsIntersected(this, other) {
		r1 = this.Range(this.DeRange())
		return
	}
	//如果相等或被other所完全包含,则头尾都是零值
	if Equql(this, other) || IsIncluded(other, this) {
		return
	}
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	//如果完全包含了对方，且没有端点重合时，则first区间是以this的开始为开始，以other开始为结束。
	//second区间则是以other的结束为开始，this的结束为结束。
	if IsIncluded(this, other) && thisStart != otherStart && thisEnd != otherEnd {
		r1 = this.Range(thisStart, otherStart)
		r2 = this.Range(otherEnd, thisEnd)
		return
		//完全包含对方，起点重合，则r1是二者终点构成的区间
	} else if IsIncluded(this, other) && thisStart == otherStart {
		r1 = this.Range(otherEnd, thisEnd)
		return
		//完全包含对方，终点重合，则r1是二者起点构成的区间
	} else if IsIncluded(this, other) && thisEnd == otherEnd {
		r1 = this.Range(thisStart, otherStart)
		return
	}
	//如果只包含对方一个点，则分两种情况
	//如果other只有起点在this之内，则r1区间是二者起点构成的区间。
	if this.IsIncludedPoint(otherStart) {
		r1 = this.Range(thisStart, otherStart)
		return
	}
	//如果ohter只有终点在this之内，则r1区间是二者的终点构成。
	if this.IsIncludedPoint(otherEnd) {
		r1 = this.Range(otherEnd, thisEnd)
		return
	}
	return
}

func IsBefore[P comparable, R any](this, other Range[P, R]) bool {
	otherStart, _ := other.DeRange()
	return this.IsBeforePoint(otherStart)
}

func IsAfter[P comparable, R any](this, other Range[P, R]) bool {
	_, otherEnd := other.DeRange()
	return this.IsAfterPoint(otherEnd)
}

func ToString[P comparable, R any](r Range[P, R]) string {
	start, end := r.DeRange()
	return fmt.Sprintf("[%v,%v)", start, end)
}
