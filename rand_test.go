package rand

import (
	"testing"
)

func TestRand(t *testing.T) {
	t.Log(Int(), Intn(10))
	t.Log(Int32(), Int32n(10))
	t.Log(Int64(), Int64n(10))
	t.Log(Float32(), Float64())
	t.Log(Perm(10))
	t.Log(RandRange(1, 3),
		RandRangeInt32(1291212412, 2091212412),
		RandRangeInt64(129121241212912124, 229121241212912124))
	t.Log(Bool())

	t.Log(String(32, RST_LOWER))
	t.Log(String(32, RST_UPPER))
	t.Log(String(32, RST_NUMBER))
	t.Log(String(32, RST_SYMBOL))
	t.Log(String(32, 0))

	t.Log(StringArray(&[]string{"张三", "李四", "王五", "赵六"}))
	t.Log(CarPlate())
	t.Log(CarPlateNewEnergy())
	t.Log(ChineseName())

	t.Log(Phone())

}
