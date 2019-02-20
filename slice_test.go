package rand

import (
	"testing"
	"time"
)

func TestSlice(t *testing.T) {
	t.Log(time.Now())
	m := map[string]int{}
	for i := 0; i < 300; i++ {
		v := GetRand().WeightMapOnce(map[interface{}]int64{
			"a": 1,
			"b": 2,
		})
		m[v.(string)]++
	}
	t.Log(m)
}
