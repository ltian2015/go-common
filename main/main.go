package main

import (
	"fmt"
	"time"

	"com.example/common/cycle"
	"com.example/common/ranges"
)

func main() {
	var floatRng = ranges.CreateNumberRange(1.1, 5.1)
	var int32Rng ranges.NumberRange[int32] = ranges.CreateNumberRange[int32](1, 43)
	println(floatRng.String())
	println(int32Rng.String())

	var intCycle = ranges.NumCycle[int32]{Count: 10, Unit: 1}
	f := &ranges.NRCycleFunc[int32]{}
	var cc = cycle.NewCycleCalculator[ranges.NumberRange[int32], ranges.NumCycle[int32]](int32Rng, intCycle, f)
	_, int32Rng2 := cc.Next()
	_, int32Rng3 := cc.Next()
	println(int32Rng2.String())
	println(int32Rng3.String())
	var i int = 12
	f2 := &ranges.NPCycleFunc[int]{}
	var intCycle2 = ranges.NumCycle[int]{Count: 10, Unit: 1}
	var c2 = cycle.NewCycleCalculator[int, ranges.NumCycle[int]](i, intCycle2, f2)
	var index1, j = c2.Next()
	index2, z := c2.Next()
	println("------------------------")
	println("J=", j, index1)
	println("Z=", z, index2)
	var k int = 0
	f3 := &ranges.NPCycleFunc[int]{}
	var c3 = cycle.NewCycleCalculator[int, ranges.NumCycle[int]](k, intCycle2, f3)
	var _, n = c3.Next()
	var intRng = ranges.CreateNumberRange(k, n)
	println(intRng.String())
	f5 := &ranges.NRCycleFunc[int]{}
	var c4 = cycle.NewCycleCalculator[ranges.NumberRange[int], ranges.NumCycle[int]](intRng, intCycle2, f5)
	for i = 0; i < 10; i++ {
		_, r := c4.Next()
		//println(r.String())
		_ = r
	}
	const TIME_LAYOUT = "2006-01-02 15:04:05"
	var t1, _ = time.Parse(TIME_LAYOUT, "2022-01-01 00:00:00")
	var tc = ranges.TimeCycle{Count: 15, Unit: 1 * time.Minute}
	var tpcc = cycle.NewCycleCalculator[time.Time, ranges.TimeCycle](t1, tc, &ranges.TPCycleFunc{})
	var _, t2 = tpcc.Next()
	var ti1 ranges.TimeInterval = ranges.CreateTimeInterval(t1, t2)
	var ticc = cycle.NewCycleCalculator[ranges.TimeInterval, ranges.TimeCycle](ti1, tc, &ranges.TICycleFunc{})
	//fmt.Println(ti1.String())
	for i = 0; i < 96; i++ {
		_, ti := ticc.Pre()
		fmt.Println(ti.String())
	}
	//ticc.Reset()
	println("-----------------------")
	curIndex, curValue := ticc.Current()
	println(curIndex, " - ", curValue.String())
}
