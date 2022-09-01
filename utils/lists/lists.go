package lists

import (
	"github.com/Caisin/caisin-go/utils/math"
	"sort"
)

type ConstSet interface {
	string | int |
		int64 | int32 | int8 |
		uint64 | uint32 | uint8 |
		float64 | float32
}

func Sort(list *[]any) {
	sort.Slice(list, func(i, j int) bool {
		return false
	})
}

func List[T any]() []T {
	return make([]T, 0)
}
func Split500[T any](list []T) (ret [][]T) {
	return Split(list, 500)
}
func Split[T any](list []T, bachSize int) (ret [][]T) {
	size := len(list)
	if size > 0 {
		for i := 0; i < size; i++ {
			start := math.Min(i, size)
			cs := list[start:math.Min(start+bachSize, size)]
			ret = append(ret, cs)
			i += bachSize - 1
		}
	}
	return ret
}

func List2Map[K ConstSet, V any](list []V, keyFun func(v V) K) (ret map[K]V) {
	size := len(list)
	if size > 0 {
		ret = make(map[K]V)
		for _, v := range list {
			ret[keyFun(v)] = v
		}
	}
	return
}

func GenList2Map[K ConstSet, V, D any](list []D, kvFun func(v D) (K, V)) (ret map[K]V) {
	size := len(list)
	if size > 0 {
		ret = make(map[K]V)
		for _, v := range list {
			key, value := kvFun(v)
			ret[key] = value
		}
	}
	return
}

func Trans[A, B any](list []A, transFun func(v A) B) (ret []B) {
	size := len(list)
	if size > 0 {
		ret = make([]B, 0)
		for _, v := range list {
			b := transFun(v)
			ret = append(ret, b)
		}
	}
	return
}

func BoolMap[T ConstSet](list []T) (ret map[T]bool) {
	if len(list) > 0 {
		ret = make(map[T]bool)
		for _, t := range list {
			ret[t] = true
		}
	}
	return
}
func Group[K ConstSet, V any](list []V, kvFun func(v V) K) (ret map[K][]V) {
	size := len(list)
	if size > 0 {
		ret = make(map[K][]V)
		for _, v := range list {
			key := kvFun(v)
			vs, ok := ret[key]
			if ok {
				ret[key] = append(vs, v)
			} else {
				ret[key] = []V{v}
			}
		}
	}
	return
}

func DelRepeat[T float64 | int64 | uint64 | int | uint | int32 | uint32 | int8 | uint8 | float32 | string](list []T) []T {
	m := make(map[T]any)
	for _, uid := range list {
		m[uid] = true
	}
	list = make([]T, 0)
	for uid := range m {
		list = append(list, uid)
	}
	return list
}
