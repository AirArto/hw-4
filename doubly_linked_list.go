package linkedlist

type List interface {
	Len() int                             // длина списка
	Front() *ExampleItem                  // первый ExampleItem
	Back() *ExampleItem                   // последний ExampleItem
	PushFront(v interface{}) *ExampleItem // добавить значение в начало
	PushBack(v interface{}) *ExampleItem  // добавить значение в конец
	Remove(i *ExampleItem)                // удалить элемент
	MoveToFront(i *ExampleItem)           // переместить элемент в начало
}

//ExampleList
type ExampleList struct {
	length int
	first  *ExampleItem
	last   *ExampleItem
}

//ExampleItem in double linked list
type ExampleItem struct {
	Value    interface{}
	CacheKey string
	Prev     *ExampleItem
	Next     *ExampleItem
}

//Len (length) of double-linked list
func (l *ExampleList) Len() int {
	return l.length
}

//First ExampleItem of double-linked list
func (l *ExampleList) Front() *ExampleItem {
	return l.first
}

//Last ExampleItem of double-linked list
func (l *ExampleList) Back() *ExampleItem {
	return l.last
}

//PushFront added one more ExampleItem element to the front of double-linked ExampleList
func (l *ExampleList) PushFront(v interface{}) *ExampleItem {
	item := ExampleItem{
		Value: v,
		Prev:  nil,
		Next:  l.first,
	}
	if l.length != 0 {
		l.first.Prev = &item
	} else {
		l.last = &item
	}
	l.first = &item
	l.length++
	return &item
}

//PushBack added one more ExampleItem element to the back of double-linked ExampleList
func (l *ExampleList) PushBack(v interface{}) *ExampleItem {
	item := ExampleItem{
		Value: v,
		Prev:  l.last,
		Next:  nil,
	}
	if l.length != 0 {
		l.last.Next = &item
	} else {
		l.first = &item
	}
	l.last = &item
	l.length++
	return &item
}

//Remove removed one ExampleItem element from double-linked ExampleList
func (l *ExampleList) Remove(i *ExampleItem) {
	if i.Next != nil {
		i.Next.Prev = i.Prev
	}
	if i.Prev != nil {
		i.Prev.Next = i.Next
	}
	if l.last == i {
		l.last = i.Prev
	}
	if l.first == i {
		l.first = i.Next
	}
	l.length--
}

//MoveToFront
func (l *ExampleList) MoveToFront(i *ExampleItem) {
	l.Remove(i)

	i.Prev = nil
	i.Next = l.first
	if l.length != 0 {
		l.first.Prev = i
	} else {
		l.last = i
	}
	l.first = i
	l.length++
}

//ToSlice returned ExampleList converted to slice
func (l *ExampleList) ToSlice() []interface{} {
	var retVal []interface{}
	item := l.first
	for i := 0; i < l.length; i++ {
		retVal = append(retVal, item.Value)
		item = item.Next
	}
	return retVal
}

type Cache interface {
	Set(key string, value interface{}) bool // Добавить значение в кэш по ключу
	Get(key string) (interface{}, bool)     // Получить значение из кэша по ключу
	Clear()                                 // Очистить кэш
}

//ExampleCache
type ExampleCache struct {
	Capacity   int
	Queue      List
	Dictionary map[string]*ExampleItem
}

//Set
func (c *ExampleCache) Set(key string, value interface{}) bool {
	if len(c.Dictionary) == 0 {
		c.Dictionary = map[string]*ExampleItem{}
	}
	elem, isExist := c.Dictionary[key]
	if isExist {
		c.Dictionary[key].Value = value
		c.Queue.MoveToFront(elem)
		return isExist
	}
	c.Dictionary[key] = c.Queue.PushFront(value)
	c.Dictionary[key].CacheKey = key
	if c.Queue.Len() > c.Capacity {
		delete(c.Dictionary, c.Queue.Back().CacheKey)
		c.Queue.Remove(c.Queue.Back())
	}
	return isExist
}

//Get
func (c *ExampleCache) Get(key string) (interface{}, bool) {
	elem, isExist := c.Dictionary[key]
	if isExist {
		c.Queue.MoveToFront(elem)
		return c.Dictionary[key].Value, isExist
	}
	return nil, isExist
}

//Clear
func (c *ExampleCache) Clear() {
	c.Capacity = 0
	c.Queue = *(new(List))
	c.Dictionary = map[string]*ExampleItem{}
}
