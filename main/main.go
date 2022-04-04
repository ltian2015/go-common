package main

import "com.example/common/ranges"

func main() {
	var floatRng = ranges.CreateNumberRange(1.1, 5.1)
	var int32Rng ranges.NumberRange[int32] = ranges.CreateNumberRange[int32](1, 43)
	println(floatRng.String())
	println(int32Rng.String())
}
