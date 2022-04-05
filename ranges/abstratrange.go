/**
ranges包抽象了所有由起点和终点（不包括在内）所构成的左闭右开区间—— 形如[Pstart, Pend)
任何具体类型，只要实现了如下几个简单的方法：
    //Range方法用给定的起点与终点创建一个新的Range
     Range(start, end P) R
	// DeRange方法将自身析构为起点与终点
	DeRange() (start, end P)
	//IsIncludedPoint方法判断给定的点p是否在区间之内
	IsIncludedPoint(p P) bool
	//IsBeforePoint方法判断区间是否在给定的点p之前
	IsBeforePoint(p P) bool
	//IsAfterPoint方法判断区间是否在给定的点P之后
	IsAfterPoint(p P) bool
   只要所有具体类型的Range实现了上述的简单方法，ranges包就可以利用这些共性，编写函数，对Range进行更加复杂的操作。
   //而具体Range类型，利用ranges包提供的这些函数，可以简化以下Range方法的实现：

    //IsPoint方法判断给定的区间是否是一个点，当区间的起点与终点相同时，区间就是一个点。
	IsPoint() bool
	//将区间化为字符串，缺省的格式[startString,endstring)
	String() string
	//以下方法可以借助ranges包提供相应函数，帮助具体类型简化这些方法的实现逻辑。
	// Union方法求区间与另一个区间(other)的并集，并判断并集是否由两个相继（相交）的区间构成，
	Union(other R) (bool, R)
	//UnionOthers计算一个区间与另外一些区间的
	UnionOthers(others []R) (bool, R)
	IsIntersected(other R) bool
	Intersect(other R) (bool, R)
	IntersectOthers(others []R) (bool, R)
	IsBefore(other R) bool
	IsAfter(other R) bool
**/
package ranges

import "fmt"

//Range是指是由类型参数P的值作为起点和终点（不包括在内）的左闭右开区间。
//数据结构形如，struct {start P,end P},数学表示形式如， [start,end)
//从数学上讲，start永远小于end，end永远在start之后，如果二者相等，区间实际上就是一个点。
//Range接口中，类型参数P是构成Range起点与终点的点元素的类型,
//而类型参数R则是实现了Range[P,R]接口的具体类型。
//这样写，是由于GO不支持嵌套的类型参数定义，无法写Range[P comparable,Range[P,Range]]，
//所有用[R any]来代替Range[P,R]，虽然，go1.18不允许直接对泛型的类型参数使用类型断言,比如，r T=P.(T)
//但是，对于任何满足Range[P，R]的具体类型的值rp，则
//r R=(interface{})(rp).(R)
//都是安全的类型转换。
//反之，对于任何满足R[P]的具体类型的值r，则
// rp Range[P,R]=(interface{})(r).(Range[P,R])
//也是安全的类型转换。
//这里，还需要注意的是，“零值区间”是一类特殊的区间，该区间的起点和终点都是参数类型P的零值。
type Range[P comparable, R any] interface {
	//以下方法需要实现者依靠自己去实现
	// Range方法用给定的起点与终点创建一个新的Range
	Range(start, end P) R
	// DeRange方法将自身析构为起点与终点
	DeRange() (start, end P)
	//IsIncludedPoint方法判断给定的点p是否在区间之内
	IsIncludedPoint(p P) bool
	//IsBeforePoint方法判断区间是否在给定的点p之前
	IsBeforePoint(p P) bool
	//IsAfterPoint方法判断区间是否在给定的点P之后
	IsAfterPoint(p P) bool

	////////////////////////////////////////////////////////////////

	//以下方法可以借助ranges包提供相应函数，帮助具体类型简化这些方法的实现逻辑。
	//IsPoint方法判断给定的区间是否是一个点，当区间的起点与终点相同时，区间就是一个点。
	IsPoint() bool
	//将区间化为字符串，缺省的格式[startString,endstring)
	String() string

	//Equal方法判断区间是否与另一个区间（other）相等。也就是起点与起点相等，终点与终点相等。
	Equal(other R) bool
	// Union方法求区间与另一个区间(other)的并集，也就是最小的起点与最大的终点所构成的区间，返回false表明结果区间是不相邻区间构成的。
	Union(other R) (bool, R)
	//UnionOthers方法计算一个区间与另外一些区间（others）的并集，如果存在两个区间不相邻，则返回false。
	UnionOthers(others []R) (bool, R)
	//IsIntersected方法计算是否与其他区间other相交，返回false表示不相交。
	IsIntersected(other R) bool
	//Intersect方法计算区间与另一个区间的交集，如果不相交，返回false，并且，结果区间“零值区间”，即：区间起点终点都是零值。
	Intersect(other R) (bool, R)
	// IntersectOthers计算区间与另一些区间的交集，如果都不相交，返回false，并且，结果区间“零值区间”。
	IntersectOthers(others []R) (bool, R)
	//Except方法求区间与other差集，也就是去掉other的剩余部分。
	//返回结果最多是两个区间，r1和r2。
	//当other是this的真子集时，差集分为两段，r1和r2都是非零值区间。
	//如果两个集合有部分交集，或者完全不相交，只有一个差集，结果r1是二者的差集，此时r2为零值区间。
	//如果两个区间相等，或者this是other的真子集，则差集是零值区间，因而，结果区间r1和r2都是零值区间。
	Except(other R) (r1, r2 R)
	//IsBefore计算区间是否在另一个区间(other)之前，也就是区间的终点是否在other区间之前。
	IsBefore(other R) bool
	//IsAfter计算区间是否在另一个区间（other）之后，也就是区间的起点是否在other区间之后。
	IsAfter(other R) bool
}

//typeTo函数将给定源类型S的值s，转换目标类型D的值。
//由于range包中，需要转换的两个参数类型实际应用中都会被赋予相同的类型，所以使用typeTo是安全的。
//此方法仅用于ranges包内部使用。
func typeTo[S, D any](s S) D {
	result, ok := (interface{})(s).(D)
	if !ok {
		panic("bad type conversion")
	}
	return result
}

// v2s函数把任何一个类型的变量转换为字符串
// 此方法仅用于ranges包内部使用。
func v2s[V any](v V) string {
	return fmt.Sprintf("%v", v)
}

///////////////////下面是ranges包提供的辅助函数,可以帮助接口的实现者快速实现功能/////////////////////////
//IntersectOthers函数计算两个区间，this与other是否相交。
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
func Equal[P comparable, R any](this, other Range[P, R]) bool {
	thisStart, thisEnd := this.DeRange()
	otherStart, otherEnd := other.DeRange()
	return thisStart == otherStart && thisEnd == otherEnd

}

//IsIncludedPoint方法判断other区间是否在this区间之内，是this区间的子集。
func IsIncluded[P comparable, R any](this, other Range[P, R]) bool {
	otherStart, otherEnd := other.DeRange()
	_, thisEnd := this.DeRange()
	var result bool = this.IsIncludedPoint(otherStart) && (this.IsIncludedPoint(otherEnd) || thisEnd == otherEnd)
	return result
}

//Intersect函数计算this区间与other区间的交集，如果不相交，返回false，并且，结果区间“零值区间”，即：区间起点终点都是零值。
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

// IntersectOthers函数计算区间this与另一些区间(others)的交集，如果与所有其他区间都不相交，返回false，并且，结果区间“零值区间”。
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

// Union函数求this区间与other区间(other)的并集，也就是最小的起点与最大的终点所构成的区间。false表明结果区间是不相邻区间构成的。
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

//UnionOthers函数计算this区间与另外一些区间（others）的并集，（最小的起点与最大的终点构成的区间。）如果存在两个区间不相邻，则返回false。
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

//Except函数求区间this与other差集，也就是this区间去掉other的剩余部分。
//返回结果最多是两个区间，r1和r2。
//当other是this的真子集时，差集分为两段，r1和r2都是非零值区间。
//如果两个集合有部分交集，或者完全不相交，只有一个差集，结果r1是二者的差集，此时r2为零值区间。
//如果两个区间相等，或者this是other的真子集，则差集是零值区间，因而，结果区间r1和r2都是零值区间。
func Except[P comparable, R any](this, other Range[P, R]) (r1, r2 R) {
	//r1,r2都是空值
	// 如果二者不相交，则，range1区间是自身，second是零值区间
	if !IsIntersected(this, other) {
		r1 = this.Range(this.DeRange())
		return
	}
	//如果相等或被other所完全包含,则头尾都是零值
	if Equal(this, other) || IsIncluded(other, this) {
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

//IsBefore函数判断this区间是否在另一个区间(other)之前，也就是this区间的终点是否在other区间之前。
func IsBefore[P comparable, R any](this, other Range[P, R]) bool {
	otherStart, _ := other.DeRange()
	return this.IsBeforePoint(otherStart)
}

//IsAfter函数计算this区间是否在另一个区间（other）之后，也就是this区间的起点是否在other区间之后。
func IsAfter[P comparable, R any](this, other Range[P, R]) bool {
	_, otherEnd := other.DeRange()
	return this.IsAfterPoint(otherEnd)
}

//RngToStr函数用来辅助将区间r按照固有格式[startStr,endStr）转换为字符串。这里，输入参数中的f函数
//负责将类型P的值转换为string。
func RngToStr[P comparable, R any](r Range[P, R], f func(P) string) string {
	start, end := r.DeRange()
	return "[" + f(start) + "," + f(end) + ")"
}
