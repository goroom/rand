package rand

import (
	"math/rand"
	"time"
)

var (
	g_rand *rand.Rand
)

func init() {
	g_rand = rand.New(rand.NewSource(time.Now().UnixNano()))
}

func Int() int {
	return g_rand.Int()
}

func Intn(n int) int {
	return g_rand.Intn(n)
}

func Int32() int32 {
	return g_rand.Int31()
}

func Int32n(n int32) int32 {
	return g_rand.Int31n(n)
}

func Int64() int64 {
	return g_rand.Int63()
}

func Int64n(n int64) int64 {
	return g_rand.Int63n(n)
}

func Float32() float32 {
	return g_rand.Float32()
}

func Float64() float64 {
	return g_rand.Float64()
}

func Perm(count int) []int {
	return g_rand.Perm(count)
}

//范围随机
func RandRange(min int, max int) int {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return Intn(max-min+1) + min
}

func RandRangeInt32(min int32, max int32) int {
	return RandRange(int(min), int(max))
}

func RandRangeInt64(min int64, max int64) int64 {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return Int64n(max-min+1) + min
}

func Bool() bool {
	if g_rand.Intn(2) == 0 {
		return true
	}
	return false
}
