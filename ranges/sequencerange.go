package ranges

//Sequencable约束中，类型参数S是指符合Sequencable接口要求的类型本身，
//也就是S==any Type impl Sequencable, 二者类型完全相同,S==Sequencable[S]
//这样写，是由于GO不支持嵌套的类型参数定义，无法写Sequencable [S Sequencable[any]]，
// 所以用S any类型代表所有符合Sequencable约束的类型。
// 因而，对于任何满足Sequencable[S]的具体类型的值sq ,则 s S=(interface{})(sq).(S)都是安全的类型转换。
type Sequencable[S any] interface {
	comparable //Range 接口要求类型参数必须是可比较的类型。需要在可比较类型基础上增加三个方法约束即可。
	Equal(s S) bool
	Before(s S) bool
	After(s S) bool
}

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
	return RngToStr[P, SeqRange[P, T]](sr, v2s[P])
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
	var rgThis Range[P, SeqRange[P, T]] = sr
	return IntersectOthers(rgThis, others)
}

func (sr SeqRange[P, T]) Union(other SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var rgThis Range[P, SeqRange[P, T]] = sr
	var rgOther Range[P, SeqRange[P, T]] = other
	return Union(rgThis, rgOther)
}

func (sr SeqRange[P, T]) UnionOthers(others []SeqRange[P, T]) (bool, SeqRange[P, T]) {
	var rgThis Range[P, SeqRange[P, T]] = sr
	return UnionOthers(rgThis, others)
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
