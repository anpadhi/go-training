package main

import "math"

func main()  {
	maxInt8 := math.MaxInt8
	minInt8 := math.MinInt8
	maxInt16 := math.MaxInt16
	minInt16 := math.MinInt16
	maxInt32 := math.MaxInt32
	minInt32 := math.MinInt32
	maxInt64 := math.MaxInt64
	minInt64 := math.MinInt64

	maxUint8 := math.MaxUint8
	maxUint16 := math.MaxUint16
	maxUint32 := math.MaxUint32
	maxUint64 := math.MaxUint16

	maxFloat32 := math.MaxFloat32
	maxFloat64 := math.MaxFloat64

	println("Range of Int8::", minInt8, " to ", maxInt8)
	println("Range of Int16::", minInt16, " to ", maxInt16)
	println("Range of Int32::", minInt32, " to ", maxInt32)
	println("Range of Int64::", minInt64, " to ", maxInt64)
	println("Max Uint8::", maxUint8)
	println("Max Uint16::", maxUint16)
	println("Max Uint32::", maxUint32)
	println("Max Uint64::", maxUint64)
	println("Max Float32::", maxFloat32)
	println("Max Float64::", maxFloat64)
}