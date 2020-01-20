package linkedlist

//Item in double linked list
type Item struct {
	value interface{}
	next  *Item
	prev  *Item
}

//List double linked
type List struct {
	content []Item
	first   *Item
	last    *Item
}

//Len (length) of double-linked list
func (l *List) Len() int {
	return 1
}

//First Item of double-linked list
func (l *List) First() *Item {
	return 1
}

//Last Item of double-linked list
func (l *List) Last() *Item {
	return 1
}

//PushFront added one more Item element to the front of double-linked List
func (l *List) PushFront(v interface{}) *List {
	return 1
}

//PushBack added one more Item element to the back of double-linked List
func (l *List) PushBack(v interface{}) *List {
	return 1
}

//Remove removed one Item element from double-linked List
func (l *List) Remove(i Item) *List {
	return 1
}

//Value of Item
func (i *Item) Value() interface{} {
	return 1
}

//Next Item of double-linked list
func (i *Item) Next() *Item {
	return 1
}

//Prev Item of double-linked list
func (i *Item) Prev() *Item {
	return 1
}
