package ranges

import (
	"testing"
	"time"
)

func TestGenericNumberRange(t *testing.T) {
	nr0 := CreateNumberRange(0, 3)
	nr1 := CreateNumberRange(1, 6)
	nr2 := CreateNumberRange(3, 5)
	nr3 := CreateNumberRange(3, 10)
	nr4 := CreateNumberRange(11, 15)
	nrs := []NumberRange[int]{nr0, nr1, nr2, nr3, nr4}

	var r0, r1, r2, r3, r4 Range[int, NumberRange[int]] = nr0, nr1, nr2, nr3, nr4
	_, b := Intersect[int, NumberRange[int]](nr1, nr2)
	_ = b

	var isIntersect1, ir1 = Intersect(r1, r2)
	var isIntersect2, ir2 = Intersect(r1, r3)
	var isIntersect3, ir3 = Intersect(r1, r4)
	println("r0: ", r0.String())
	println("r1: ", r1.String())
	println("r2: ", r2.String())
	println("r3: ", r3.String())
	println("r4: ", r4.String())
	println("r1*r2 : ", ir1.String(), isIntersect1)
	println("r1*r3 : ", ir2.String(), isIntersect2)
	println("r1*r4 : ", ir3.String(), isIntersect3)
	var isSuccessive1, ur1 = Union(r1, r2)
	var isSuccessive2, ur2 = Union(r1, r3)
	var isSuccessive3, ur3 = Union(r1, r4)

	println("r1+r2: ", ur1.String(), isSuccessive1)
	println("r1+r3: ", ur2.String(), isSuccessive2)
	println("r1+r4: ", ur3.String(), isSuccessive3)
	ok, nr5 := nr1.UnionOthers(nrs)
	println("r1+..r4 ", nr5.String(), ok)
	var ok2, nr6 = nr1.IntersectOthers(nrs[:3])
	println("r0*..r3 ", nr6.String(), ok2)
	for _, v := range nrs[:3] {
		println(v.String())
	}
	var er1, er2 = Except(r1, r1)
	println("r1-r1: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r1, r2)
	println("r1-r2: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r1, r3)
	println("r1-r3: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r1, r4)
	println("r1-r4: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r3, r1)
	println("r3-r1: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r3, r2)
	println("r3-r2: result1=", er1.String(), "result2=", er2.String())
	er1, er2 = Except(r2, r1)
	println("r2-r1: result1=", er1.String(), "result2=", er2.String())
}
func TestSequenceRange(t *testing.T) {
	var t1 = time.Now()
	var t2 = t1.Add(24 * time.Hour)
	var t3 = t1.Add(36 * time.Hour)
	var t4 = t1.Add(48 * time.Hour)

	var ti1 TimeInterval = CreateTimeInterval(t1, t2)
	var ti2 TimeInterval = CreateTimeInterval(t1, t3)
	var ti3 TimeInterval = CreateTimeInterval(t3, t4)
	var ti4 TimeInterval = CreateSeqRange[time.Time, time.Time](t1, t4)
	var yes bool
	var resultTi TimeInterval
	yes, resultTi = ti1.Intersect(ti2)
	println(Tintvl2Str(ti1), "+", Tintvl2Str(ti2), "=", Tintvl2Str(resultTi), yes)
	yes, resultTi = ti1.Intersect(ti3)
	println(Tintvl2Str(ti1), "+", Tintvl2Str(ti3), "=", Tintvl2Str(resultTi), yes)

	yes, resultTi = ti1.Intersect(ti4)
	println(Tintvl2Str(ti1), "+", Tintvl2Str(ti4), "=", Tintvl2Str(resultTi), yes)
	yes, resultTi = ti2.Intersect(ti3)
	println(Tintvl2Str(ti2), "+", Tintvl2Str(ti3), "=", Tintvl2Str(resultTi), yes)

}
