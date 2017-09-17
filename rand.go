package rand

import (
	"math/rand"
	"sync"
	"time"
)

type Rand struct {
	sync.Mutex
	rand *rand.Rand
}

var (
	gRand *Rand
)

func init() {
	gRand = &Rand{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (r *Rand) ExpFloat64() float64        { r.Lock(); defer r.Unlock(); return r.rand.ExpFloat64() }
func (r *Rand) Float32() float32           { r.Lock(); defer r.Unlock(); return r.rand.Float32() }
func (r *Rand) Float64() float64           { r.Lock(); defer r.Unlock(); return r.rand.Float64() }
func (r *Rand) Int() int                   { r.Lock(); defer r.Unlock(); return r.rand.Int() }
func (r *Rand) Int31() int32               { r.Lock(); defer r.Unlock(); return r.rand.Int31() }
func (r *Rand) Int31n(n int32) int32       { r.Lock(); defer r.Unlock(); return r.rand.Int31n(n) }
func (r *Rand) Int63() int64               { r.Lock(); defer r.Unlock(); return r.rand.Int63() }
func (r *Rand) Int63n(n int64) int64       { r.Lock(); defer r.Unlock(); return r.rand.Int63n(n) }
func (r *Rand) Intn(n int) int             { r.Lock(); defer r.Unlock(); return r.rand.Intn(n) }
func (r *Rand) NormFloat64() float64       { r.Lock(); defer r.Unlock(); return r.rand.NormFloat64() }
func (r *Rand) Perm(n int) []int           { r.Lock(); defer r.Unlock(); return r.rand.Perm(n) }
func (r *Rand) Read(p []byte) (int, error) { r.Lock(); defer r.Unlock(); return r.rand.Read(p) }
func (r *Rand) Seed(seed int64)            { r.Lock(); defer r.Unlock(); r.rand.Seed(seed) }
func (r *Rand) Uint32() uint32             { r.Lock(); defer r.Unlock(); return r.rand.Uint32() }
func (r *Rand) Uint64() uint64             { r.Lock(); defer r.Unlock(); return r.rand.Uint64() }

func GetRand() *Rand {
	return gRand
}

//func ExpFloat64() float64        { return gRand.ExpFloat64() }
//func Float32() float32           { return gRand.Float32() }
//func Float64() float64           { return gRand.Float64() }
//func Int() int                   { return gRand.Int() }
//func Int31() int32               { return gRand.Int31() }
//func Int31n(n int32) int32       { return gRand.Int31n(n) }
//func Int63() int64               { return gRand.Int63() }
//func Int63n(n int64) int64       { return gRand.Int63n(n) }
//func Intn(n int) int             { return gRand.Intn(n) }
//func NormFloat64() float64       { return gRand.NormFloat64() }
//func Perm(n int) []int           { return gRand.Perm(n) }
//func Read(p []byte) (int, error) { return gRand.Read(p) }
//func Seed(seed int64)            { gRand.Seed(seed) }
//func Uint32() uint32             { return gRand.Uint32() }
//func Uint64() uint64             { return gRand.Uint64() }

func SetDefaultRand(r *Rand) {
	gRand = r
}

//范围随机
func (r *Rand) RandRange(min int, max int) int {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return r.rand.Intn(max-min+1) + min
}

func (r *Rand) RandRangeInt32(min int32, max int32) int {
	return r.RandRange(int(min), int(max))
}

func (r *Rand) RandRangeInt64(min int64, max int64) int64 {
	if max <= min {
		if max == min {
			return min
		}
		return 0
	}
	return r.rand.Int63n(max-min+1) + min
}

func (r *Rand) Bool() bool {
	gRand.Lock()
	defer gRand.Unlock()
	if gRand.rand.Intn(2) == 0 {
		return true
	}
	return false
}
