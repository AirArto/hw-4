package linkedlist

//List double linked
type List struct {
	length int
	first  *Item
	last   *Item
}

//Len (length) of double-linked list
func (l *List) Len() int {
	return l.length
}

//First Item of double-linked list
func (l *List) First() *Item {
	return l.first
}

//Last Item of double-linked list
func (l *List) Last() *Item {
	return l.last
}

//PushFront added one more Item element to the front of double-linked List
func (l *List) PushFront(v interface{}) *List {
	item := Item{v, nil, l.first}
	if l.length != 0 {
		l.first.prev = &item
	} else {
		l.last = &item
	}
	l.first = &item
	l.length++
	return l
}

//PushBack added one more Item element to the back of double-linked List
func (l *List) PushBack(v interface{}) *List {
	item := Item{v, l.last, nil}
	if l.length != 0 {
		l.last.next = &item
	} else {
		l.first = &item
	}
	l.last = &item
	l.length++
	return l
}

//Remove removed one Item element from double-linked List
func (l *List) Remove(i Item) *List {
	if i.next != nil {
		i.next.prev = i.prev
	}
	if i.prev != nil {
		i.prev.next = i.next
	}
	l.length--
	return l
}

//ToSlice returned List converted to slice
func (l *List) ToSlice() []interface{} {
	var retVal []interface{}
	item := l.first
	for i := 0; i < l.length; i++ {
		retVal = append(retVal, item.value)
		item = item.Next()
	}
	return retVal
}

//Item in double linked list
type Item struct {
	value interface{}
	prev  *Item
	next  *Item
}

//Value of Item
func (i *Item) Value() interface{} {
	return i.value
}

//Next Item of double-linked list
func (i *Item) Next() *Item {
	if i.next != nil {
		return i.next
	}
	return i
}

//Prev Item of double-linked list
func (i *Item) Prev() *Item {
	if i.prev != nil {
		return i.prev
	}
	return i
}
