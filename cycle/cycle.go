package cycle

type Cycle[U any] interface {
	GetCount() int
	GetUnit() U
}

/**
**/
// CycleFunc接口定义了一个泛型表示的周期函数，
//以接口形式进行定义，主要是因为GO1.18不支持如下带有类型参数的泛型类型：
// type CycleFunc func [T,C any](t T,n int,c C) T
type CycleFunc[T, C any] interface {
	OfCycles(t T, n int, c C) T
}
type CycleCalculator[T, C any] struct {
	origin     T               //初始值
	cycle      C               //周期量值
	cycleFunc  CycleFunc[T, C] //   周期函数
	cycleIndex int             // 周期序号
}

func (rc *CycleCalculator[T, C]) Next() (int, T) {
	rc.cycleIndex += 1
	return rc.cycleIndex, rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
}
func (rc *CycleCalculator[T, C]) Pre() (int, T) {
	rc.cycleIndex -= 1
	return rc.cycleIndex, rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
}
func (rc *CycleCalculator[T, C]) Value() T {
	if rc.cycleIndex == 0 {
		return rc.origin
	} else {
		return rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
	}
}
func (rc *CycleCalculator[T, C]) Cycles() int {
	return rc.cycleIndex
}

func (rc *CycleCalculator[T, C]) Reset() {
	rc.cycleIndex = 0
}

func NewCycleCalculator[T, C any](o T, c C, clAdder CycleFunc[T, C]) *CycleCalculator[T, C] {
	return &CycleCalculator[T, C]{origin: o, cycle: c, cycleFunc: clAdder}
}
