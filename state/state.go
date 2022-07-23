package state

import (
	"errors"
	"reflect"
	"sync"
)

var (
	state = make(map[string]any)
	mu    = sync.RWMutex{}
)

func Get[T any]() (*T, error) {
	t := new(T)
	typeOf := reflect.TypeOf(*t)
	if typeOf == nil || typeOf.Name() == "" {
		return nil, errors.New("类型错误")
	}
	name := typeOf.Name()
	ret, ok := state[name]
	if !ok {
		return nil, errors.New(name + ":not exists")

	}
	return ret.(*T), nil
}

func Add(t *any) bool {
	typeOf := reflect.TypeOf(*t)
	if typeOf == nil || typeOf.Name() == "" {
		return false
	}
	name := typeOf.Name()
	mu.Lock()
	defer mu.Unlock()
	state[name] = t
	return true
}
