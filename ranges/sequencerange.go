package ranges

//Sequencable定义了具有先后顺序的类型接口，是对可以作为SeqRange区间起点和终点的类型P的一种约束。
//Range 接口要求起点和终点的类型参数P必须是comparable类型，所以Sequencable首先应该是
//comparable,在此基础上增加三个方法要求，Equal、Before和After，这三个方法用来判断点与点
//之间的先后关系。
//之所以设计Sequencable约束并要求该约束具有上述三个方法，主要考虑对time.Time类型作为Rang起点与终点
//类型的支持。time.Time具有上述三个方法。所以，具体类型time.Time符合Sequencable[time.Time]约束(接口)要求。

//Sequencable约束中，类型参数S是指符合Sequencable[S]接口要求的类型本身，
//具体类型S一定实现了接口Sequencable[S]
//这样写，是由于GO不支持嵌套的类型参数定义，无法写Sequencable [S Sequencable[S]]，
// 所以用[S any]类型代表所有符合Sequencable约束的类型。
// 因而，对于任何满足Sequencable[S]的具体类型的值sq ,
//    s S=(interface{})(sq).(S)
// 是安全的类型转换， 同样，对于任何具体类型S的值s，
//  sq Sequencable[S]=(interface{})(s)(Sequencable[S])
//也是安全的类型转换
type Sequencable[S any] interface {
	comparable
	//Equal方法比较顺序元素与另一个元素是否相等
	Equal(ohter S) bool
	//Before方法判断顺序元素是否在另一个元素之前。
	Before(other S) bool
	//After方法判断顺序元素是否在另一个元素之后。
	After(ohter S) bool
}

//CreateSeqRange函数用于给定的两个点创建一个区间，
//无论两个点的先后顺序如何，创建出来的区间的起点都会在终点之前。
func CreateSeqRange[P Sequencable[T], T any](p1, p2 P) SeqRange[P, T] {
	t2 := typeTo[P, T](p2)
	if p1.Before(t2) || p1.Equal(t2) {
		return SeqRange[P, T]{start: p1, end: p2}
	} else {
		return SeqRange[P, T]{start: p2, end: p1}
	}
}

//SeqRange[P Sequencable[T], T any]定义了各类有顺序元素组成的区间类型，该类型的区间
//满足Range[P,SeqRange[P,T]接口。
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

/////////////////////////////////////////////////////////////////////////////
///////////////以下方法使用ranges包的函数完成////////////////////////////////////
func (sr SeqRange[P, T]) IsPoint() bool {
	return IsPoint[P, SeqRange[P, T]](sr)
}

func (sr SeqRange[P, T]) String() string {
	return RngToStr[P, SeqRange[P, T]](sr, v2s[P])
}
func (sr SeqRange[P, T]) Equal(other SeqRange[P, T]) bool {
	return Equal[P, SeqRange[P, T]](sr, other)
}

func (sr SeqRange[P, T]) IsIntersected(other SeqRange[P, T]) bool {
	return IsIntersected[P, SeqRange[P, T]](sr, other)
}

func (sr SeqRange[P, T]) Intersect(other SeqRange[P, T]) (bool, SeqRange[P, T]) {
	return Intersect[P, SeqRange[P, T]](sr, other)
}

func (sr SeqRange[P, T]) IntersectOthers(others []SeqRange[P, T]) (bool, SeqRange[P, T]) {
	return IntersectOthers[P, SeqRange[P, T]](sr, others)
}

func (sr SeqRange[P, T]) Union(other SeqRange[P, T]) (bool, SeqRange[P, T]) {
	return Union[P, SeqRange[P, T]](sr, other)
}

func (sr SeqRange[P, T]) UnionOthers(others []SeqRange[P, T]) (bool, SeqRange[P, T]) {
	return UnionOthers[P, SeqRange[P, T]](sr, others)
}
func (sr SeqRange[P, T]) Except(other SeqRange[P, T]) (r1, r2 SeqRange[P, T]) {
	return Except[P, SeqRange[P, T]](sr, other)
}
func (sr SeqRange[P, T]) IsBefore(other SeqRange[P, T]) bool {
	return IsBefore[P, SeqRange[P, T]](sr, other)
}
func (sr SeqRange[P, T]) IsAfter(other SeqRange[P, T]) bool {
	return IsAfter[P, SeqRange[P, T]](sr, other)
}
