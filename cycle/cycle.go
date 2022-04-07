package cycle

type Cycle[U any] interface {
	GetCount() int
	GetUnit() U
}

/**
**/
// CycleFunc接口定义了一个泛型表示的周期函数，
//该接口只定义了一个方法，用于计算给定值t经过n个周期c后的值。
//其中，类型参数T表示给定值t的类型，而类型参数C表示周期值c的类型。
//以接口形式进行定义，主要是因为GO1.18不支持如下带有类型参数的泛型类型：
// type CycleFunc func [T,C any](t T,n int,c C) T
type CycleFunc[T, C any] interface {
	OfCycles(t T, n int, c C) T
}

//CycleCalculator定义了通用的周期计算器，
//只要给定初始值，周期量值，周期计算函数，就可以计算下一个（Next）和上一个（Pre）周期下的值。
//CycleCalculator不是线程安全类型，请注意不要在多线程环境下使用
type CycleCalculator[T, C any] struct {
	origin     T               //初始值
	cycle      C               //周期量值
	cycleFunc  CycleFunc[T, C] //   周期函数
	cycleIndex int             // 周期序号
}

//Next方法计算下一个周期值，返回下一周期序号和周期值。
func (rc *CycleCalculator[T, C]) Next() (int, T) {
	rc.cycleIndex += 1
	return rc.cycleIndex, rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
}

//Pre方法计算上一个周期值，返回上一周期序号和周期值。
func (rc *CycleCalculator[T, C]) Pre() (int, T) {
	rc.cycleIndex -= 1
	return rc.cycleIndex, rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
}

//Value方法返回当前周期序号的下值。
func (rc *CycleCalculator[T, C]) Current() (int, T) {
	if rc.cycleIndex == 0 {
		return 0, rc.origin
	} else {
		return rc.cycleIndex, rc.cycleFunc.OfCycles(rc.origin, rc.cycleIndex, rc.cycle)
	}
}

// Reset方法重置周期序号为0
func (rc *CycleCalculator[T, C]) Reset() {
	rc.cycleIndex = 0
}

//NewCycleCalculator函数用给定的初始值o，周期值c和 周期计算函数构造一个周期计算器，并返回其指针。
func NewCycleCalculator[T, C any](o T, c C, clFunc CycleFunc[T, C]) *CycleCalculator[T, C] {
	return &CycleCalculator[T, C]{origin: o, cycle: c, cycleFunc: clFunc}
}
