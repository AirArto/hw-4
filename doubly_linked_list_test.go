package linkedlist

import (
	"errors"
	"testing"
)

func TestCount(t *testing.T) {
	newList := List(&ExampleList{})
	newCache := Cache(&ExampleCache{Queue: newList, Capacity: 6})

	err := error(nil)
	slice := newCache.(*ExampleCache).Queue.(*ExampleList).ToSlice()
	targetSlice := []interface{}{}

	for i, elem := range targetSlice {
		if elem != slice[i] {
			err = errors.New("Expected and recieved list are different")
		}
	}

	if err != nil {
		t.Errorf("\n\t%s", err)
	} else {
		newCache.Set("1", "odin")
		newCache.Set("1", "odin")
		newCache.Set("1", "odin")
		newCache.Set("1", "odin")

		slice = newCache.(*ExampleCache).Queue.(*ExampleList).ToSlice()
		targetSlice = []interface{}{"odin"}
		for i, elem := range targetSlice {
			if elem != slice[i] {
				err = errors.New("Expected and recieved list are different")
			}
		}
	}

	if err != nil {
		t.Errorf("\n\t%s", err)
	} else {
		newCache.Set("2", "dva")
		newCache.Set("3", "tri")
		newCache.Set("4", "chetyre")
		newCache.Set("2", "dva")
		newCache.Set("3", "tri")
		newCache.Set("5", "piat")
		newCache.Set("6", "shest")
		newCache.Set("1", "odin")
		newCache.Set("1", "odin")
		newCache.Set("1", "odin")
		newCache.Set("7", "sem")

		slice = newCache.(*ExampleCache).Queue.(*ExampleList).ToSlice()
		targetSlice = []interface{}{"sem", "odin", "shest", "piat", "tri", "dva"}
		for i, elem := range targetSlice {
			if elem != slice[i] {
				err = errors.New("Expected and recieved list are different")
			}
		}
	}

	if err != nil {
		t.Errorf("\n\t%s", err)
	} else {
		newCache.Get("1")
		newCache.Set("7", "sem")
		newCache.Get("1")

		slice = newCache.(*ExampleCache).Queue.(*ExampleList).ToSlice()
		targetSlice = []interface{}{"odin", "sem", "shest", "piat", "tri", "dva"}
		for i, elem := range targetSlice {
			if elem != slice[i] {
				err = errors.New("Expected and recieved list are different")
			}
		}
	}
}
