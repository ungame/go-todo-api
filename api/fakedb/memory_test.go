package fakedb

import (
	"reflect"
	"testing"
)

func TestMemory_Set_integer(t *testing.T) {
	mem := New()

	var (
		key = "1"
		val = 1
	)
	mem.Set(key, val)

	got := mem.Get(key)

	if got.(int) != val {
		t.Errorf("expected: %v, got: %v", val, got.(int))
	}
}

func TestMemory_Set_struct(t *testing.T) {
	mem := New()

	type (
		Struct struct {
			ID    string
			Value int
		}
	)

	var (
		item = &Struct{ID: "1", Value: 1}
	)

	mem.Set(item.ID, item)

	got := mem.Get(item.ID)
	isEquals := reflect.DeepEqual(item, got.(*Struct))

	if !isEquals {
		t.Errorf("expected: %v, got: %v", item, got.(*Struct))
	}
}

func TestMemory_Get(t *testing.T) {
	mem := New()

	var key interface{}

	got := mem.Get(key)

	if got != nil {
		t.Errorf("expected: %v, got: %v", nil, got)
	}
}

func TestMemory_Has(t *testing.T) {
	mem := New()

	var (
		key = "1"
		val = 1
	)
	mem.Set(key, val)

	got := mem.Has(key)

	if !got {
		t.Errorf("expected: %v, got: %v", true, got)
	}
}

func TestMemory_Len(t *testing.T) {
	mem := New()

	var (
		key = "1"
		val = 1
	)
	mem.Set(key, val)

	got := mem.Len()

	if got != 1 {
		t.Errorf("expected: %v, got: %v", 1, got)
	}
}

func TestMemory_Del(t *testing.T) {
	mem := New()

	var (
		key = "1"
		val = 1
	)
	mem.Set(key, val)
	mem.Del(key)

	got := mem.Has(key)

	if got {
		t.Errorf("expected: %v, got: %v", false, got)
	}
}

func TestMemory_ForEach(t *testing.T) {
	mem := New()

	type (
		Item struct {
			Key string
			Val int
		}
	)

	var (
		items = []*Item{
			{Key: "1", Val: 1},
			{Key: "2", Val: 2},
		}
	)

	for index := range items {
		mem.Set(items[index].Key, items[index])
	}

	var (
		got = make([]*Item, 0, len(items))
	)

	mem.ForEach(func(item interface{}) bool {
		got = append(got, item.(*Item))
		return true
	})

	isEquals := reflect.DeepEqual(items, got)

	if !isEquals {
		t.Errorf("expected: %v, got: %v", items, got)
	}
}
