package linkedlist

import (
	"errors"
	"testing"
)

func TestCount(t *testing.T) {
	newList := List{}
	err := error(nil)
	slice := newList.ToSlice()
	targetSlice := []interface{}{}
	for i, elem := range targetSlice {
		if elem != slice[i] {
			err = errors.New("Expected and recieved list are different")
		}
	}
	newList.PushFront("third")
	newList.PushFront("second")
	newList.PushFront("first")
	newList.PushBack("fourth")
	newList.PushBack("fifth")
	newList.PushBack("last")
	slice = newList.ToSlice()
	targetSlice = []interface{}{"first", "second", "third", "fourth", "fifth", "last"}
	for i, elem := range targetSlice {
		if elem != slice[i] {
			err = errors.New("Expected and recieved list are different")
		}
	}

	if err != nil {
		t.Errorf("\n\t%s", err)
	} else {
		newList.Remove(*newList.First().Prev().Prev().Next().Next().Next().Next().Prev().Prev().Next().Next())
		slice = newList.ToSlice()
		targetSlice = []interface{}{"first", "second", "third", "fourth", "last"}
		for i, elem := range targetSlice {
			if elem != slice[i] {
				err = errors.New("Expected and recieved list are different")
			}
		}
	}

	if err != nil {
		t.Errorf("\n\t%s", err)
	} else {
		newList.Remove(*newList.Last())
		slice = newList.ToSlice()
		targetSlice = []interface{}{"first", "second", "third", "fourth"}
		for i, elem := range targetSlice {
			if elem != slice[i] {
				err = errors.New("Expected and recieved list are different")
			}
		}
	}

	if newList.First().Next().Value() != "second" {
		t.Errorf("\n\t%s", "Wrong value")
	}
}
