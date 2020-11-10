package c

import "strconv"

/** TYPES **/
var (
	INT_MAX   int   = 2147483647
	INT_MIN   int   = (-INT_MAX - 1)
	LLONG_MAX int64 = 9223372036854775807
	LLONG_MIN int64 = (-LLONG_MAX - 1)
)

type Pair struct {
	First  interface{}
	Second interface{}
}

type Tuple struct {
	sliceTuple []interface{}
}

func (t *Tuple) Init(array ...interface{}) {
	t.sliceTuple = array
}
func (t Tuple) Get(index int) interface{} {
	if index >= len(t.sliceTuple) {
		panic("index out of range :(")
	}
	return t.sliceTuple[index]
}

/** QUEUE **/

type nodeForQueue struct {
	next *nodeForQueue
	data interface{}
}
type Queue struct {
	head *nodeForQueue
	tail *nodeForQueue
	size int
}

func (q *Queue) Push(v interface{}) {
	var newVal = new(nodeForQueue)
	newVal.next = nil
	newVal.data = v
	if q.head == nil {
		q.head = newVal
		q.tail = newVal
	} else {
		q.tail.next = newVal
		q.tail = newVal
	}
	q.size++
}
func (q Queue) Empty() bool {
	return (q.size == 0)
}
func (q *Queue) Pop() {
	if q.Empty() == true {
		panic("cpp/pkg/types/c.go.Pop():you are poping empty queue")
	}
	if q.head.next == nil {
		q.head = nil
		q.tail = nil
	} else {
		q.head = q.head.next
	}
	q.size--
}
func (q Queue) Front() interface{} {
	if q.Empty() == true {
		panic("cpp/pkg/types/c.go.Front():you getting first element from empty queue")
	}
	return q.head.data
}
func (q Queue) Size() int {
	return (q.size)
}

/** STACK **/

type nodeForStack struct {
	next *nodeForStack
	prev *nodeForStack
	data interface{}
}
type Stack struct {
	head *nodeForStack
	tail *nodeForStack
	size int
}

func (s *Stack) Push(v interface{}) {
	var newVal = new(nodeForStack)
	newVal.next = nil
	newVal.data = v
	if s.head == nil {
		s.head = newVal
		s.tail = newVal
		s.tail.prev = nil
	} else {
		s.tail.next = newVal
		last := s.tail
		s.tail = newVal
		s.tail.prev = last
	}
	s.size++
}
func (s Stack) Empty() bool {
	return (s.size == 0)
}
func (s *Stack) Pop() {
	if s.Empty() == true {
		panic("cpp/pkg/types/stack.go.Pop(): you are poping empty stack")
	}
	if s.tail.prev == nil {
		s.head = nil
		s.tail = nil
	} else {
		s.tail = s.tail.prev
	}
	s.size--
}
func (s Stack) Top() interface{} {
	if s.Empty() == true {
		panic("cpp/pkg/types/stack.go.Top():you getting first element from empty stack")
	}
	return s.tail.data
}
func (s Stack) Size() int {
	return (s.size)
}

/** DEQUE **/
// Node node for deque
type Node struct {
	Next *Node
	Prev *Node
	Data interface{}
}

// Deque it's like std::deque in C ++
type Deque struct {
	head *Node
	tail *Node
	size int
}

// PushBack push to the end of deque new element
func (d *Deque) PushBack(v interface{}) {
	var newVal = new(Node)
	newVal.Next = nil
	newVal.Data = v
	if d.head == nil {
		d.head = newVal
		d.tail = newVal
		d.tail.Prev = nil
	} else {
		d.tail.Next = newVal
		last := d.tail
		d.tail = newVal
		d.tail.Prev = last
	}
	d.size++
}

// PushFront push to the front of deque new element
func (d *Deque) PushFront(v interface{}) {
	var newVal = new(Node)
	newVal.Data = v
	if d.head == nil {
		d.head = newVal
		d.tail = newVal
		d.tail.Prev = nil
	} else {
		phead := d.head
		d.head.Prev = newVal
		d.head = newVal
		d.head.Next = phead
	}
	d.size++
}

// Empty check's deque is empty or not
func (d *Deque) Empty() bool {
	return (d.size == 0)
}

// PopBack pop the last element
func (d *Deque) PopBack() {
	if d.Empty() == true {
		panic("cpp/pkg/types/stack.go.PopBack(): you are poping empty stack")
	}
	if d.tail.Prev == nil {
		d.head = nil
		d.tail = nil
	} else {
		d.tail = d.tail.Prev
	}
	d.size--
}

// PopFront pop the last element
func (d *Deque) PopFront() {
	if d.Empty() == true {
		panic("cpp/pkg/types/stack.go.PopFront(): you are poping empty stack")
	}
	if d.head.Next == nil {
		d.head = nil
		d.tail = nil
	} else {
		d.head = d.head.Next
	}
	d.size--
}

// Front return first element of deque
func (d *Deque) Front() interface{} {
	if d.Empty() == true {
		panic("cpp/pkg/types/stack.go.Front(): you getting first element from empty stack")
	}
	return d.head.Data
}

// Back return last element of deque
func (d *Deque) Back() interface{} {
	if d.Empty() == true {
		panic("Back(): you getting last element from empty stack")
	}
	return d.tail.Data
}

// Size return size of deque
func (d *Deque) Size() int {
	return (d.size)
}

// Begin return pointer to the first element
func (d *Deque) Begin() *Node {
	return d.head
}

// End return pointer to the last element
func (d *Deque) End() *Node {
	return d.tail
}

// Get get element by index
func (d *Deque) Get(index int) interface{} {
	if index >= d.size {
		panic("Get(): index out of range [" + strconv.Itoa(index) + "] with length " + strconv.Itoa(d.size))
	}
	var counter int = 0
	cur := d.head
	for cur != nil {
		if counter == index {
			return cur.Data
		}
		counter++
		cur = cur.Next
	}
	return nil
}
