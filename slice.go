package rand

import (
	"fmt"
	"reflect"
)

func (r *Rand) WeightStringInt64MapOnce(m map[string]int64) string {
	length := len(m)
	weight := make([]int64, 0, length)
	key := make([]string, 0, length)
	for k, v := range m {
		key = append(key, k)
		weight = append(weight, v)
	}

	ls, _ := r.WeightSliceNoRepeat(key, func(i int) int64 {
		return weight[i]
	}, 1)
	return key[ls[0]]
}

func (r *Rand) WeightMapOnce(m map[interface{}]int64) interface{} {
	length := len(m)
	weight := make([]int64, 0, length)
	key := make([]interface{}, 0, length)
	for k, v := range m {
		key = append(key, k)
		weight = append(weight, v)
	}

	ls, _ := r.WeightSliceNoRepeat(key, func(i int) int64 {
		return weight[i]
	}, 1)
	return key[ls[0]]
}

func (r *Rand) WeightSliceNoRepeat(slice interface{}, calcWeight func(i int) int64, count int) ([]int, error) {
	if count <= 0 {
		return nil, nil
	}

	// slice is not slice, return error
	v := reflect.ValueOf(slice)
	if v.Kind() != reflect.Slice {
		return nil, fmt.Errorf("util.RandWeightSlice(): error type %s, need slice", v.Kind())
	}

	// slice is empty, return nil
	sLen := v.Len()
	if sLen <= 0 {
		return nil, nil
	}

	// slice not enough, return all item if weight > 0
	if count >= sLen {
		var rt []int
		for i := 0; i < sLen; i++ {
			if calcWeight(i) > 0 {
				rt = append(rt, i)
			}
		}
		return rt, nil
	}

	// get weight
	weightList, weightSum := r.getRandSliceWeight(sLen, calcWeight)
	if weightSum <= 0 {
		return nil, nil
	}

	return r.randWeight(count, weightList, weightSum), nil
}

func (r *Rand) getRandSliceWeight(length int, calcWeight func(i int) int64) ([]int64, int64) {
	weightList := make([]int64, length)
	var weightSum int64 = 0
	for i := 0; i < length; i++ {
		weightList[i] = int64(calcWeight(i))
		if weightList[i] < 0 {
			weightList[i] = 0
		}
		weightSum += weightList[i]
	}
	return weightList, weightSum
}

func (r *Rand) randWeight(count int, weightList []int64, weightSum int64) []int {
	sLen := len(weightList)
	var rt []int
	for i := 0; i < count && weightSum > 0; i++ {
		r := int64(r.Int()) % weightSum
		var twa int64
		for j := 0; j < sLen; j++ {
			twa += weightList[j]
			if r < twa {
				weightSum -= weightList[j]
				weightList[j] = 0
				rt = append(rt, j)
				break
			}
		}
	}

	return rt
}
