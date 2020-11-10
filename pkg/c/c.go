package c

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

/** ASSERT **/
//if condition is false the program will print the error and exit
func Assert(cond bool, err string) {
	if cond == false {
		var res string = ("*  Error: " + err + "  *")
		for i := 0; i < len(res); i++ {
			fmt.Print("*")
		}
		fmt.Print("\n*\n" + res + "\n*\n")
		for i := 0; i < len(res); i++ {
			fmt.Print("*")
		}
		fmt.Print("\n")
		os.Exit(1)
	}
}

/** MATH **/
//print sqrt of number
func Sqrt(n float64) float64 {
	return math.Sqrt(n)
}

//round a number up
func Ceil(n float64) float64 {
	return math.Ceil(n)
}

//round a number down to an integer
func Floor(n float64) float64 {
	return math.Floor(n)
}

//round a number based on arifmetic laws
func Round(n float64) float64 {
	return math.Round(n)
}

//get (a ^ n) when a and n is int
func Pow(a int64, n int64) int64 {
	var res int64 = 1
	for n > 0 {
		if n%2 == 1 {
			res = res * a
		}
		a = a * a
		n >>= 1
	}
	return res
}

//get (a ^ n) % mod when a and n is int
func Powmod(a int64, n int64, mod int64) int64 {
	var res int64 = 1
	a %= mod
	for n > 0 {
		if n%2 == 1 {
			res = (res * a) % mod
		}
		a = (a * a) % mod
		n >>= 1
	}
	return res
}

//get (a ^ n) % mod when a is float and n is int
func Powf(a float64, n int) float64 {
	if n == 0 {
		return 1
	}
	var temp float64 = Powf(a, n/2)
	if n%2 == 0 {
		return temp * temp
	} else {
		if n > 0 {
			return a * temp * temp
		} else {
			return (temp * temp) / a
		}
	}
}

//return sin of n
func Sin(n float64) float64 {
	return math.Sin(n)
}

//return cos of n
func Cos(n float64) float64 {
	return math.Cos(n)
}

//return tan of n
func Tan(n float64) float64 {
	return math.Tan(n)
}

//return asin of n
func Asin(n float64) float64 {
	return math.Asin(n)
}

//return acos of n
func Acos(n float64) float64 {
	return math.Acos(n)
}

//return atan of n
func Atan(n float64) float64 {
	return math.Atan(n)
}

//return abs of n, m
func Abs(n float64) float64 {
	return math.Abs(n)
}

//return distance between two points in two-dimensional coordinate system
func Distance(x1, y1, x2, y2 float64) float64 {
	return (((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)))
}

/** OPERATORS **/
// a + b
func Add(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Add(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) + b.(int8)
	case int16:
		*a = (*a).(int16) + b.(int16)
	case int32:
		*a = (*a).(int32) + b.(int32)
	case int64:
		*a = (*a).(int64) + b.(int64)
	case uint8:
		*a = (*a).(uint8) + b.(uint8)
	case uint16:
		*a = (*a).(uint16) + b.(uint16)
	case uint32:
		*a = (*a).(uint32) + b.(uint32)
	case uint64:
		*a = (*a).(uint64) + b.(uint64)
	case int:
		*a = (*a).(int) + b.(int)
	case uint:
		*a = (*a).(uint) + b.(uint)
	case string:
		*a = (*a).(string) + b.(string)
	case float32:
		*a = (*a).(float32) + b.(float32)
	case float64:
		*a = (*a).(float64) + b.(float64)
	default:
		panic("Add(): can't get a + b")
	}
}

// a - b
func Sub(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Sub(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) - b.(int8)
	case int16:
		*a = (*a).(int16) - b.(int16)
	case int32:
		*a = (*a).(int32) - b.(int32)
	case int64:
		*a = (*a).(int64) - b.(int64)
	case uint8:
		*a = (*a).(uint8) - b.(uint8)
	case uint16:
		*a = (*a).(uint16) - b.(uint16)
	case uint32:
		*a = (*a).(uint32) - b.(uint32)
	case uint64:
		*a = (*a).(uint64) - b.(uint64)
	case int:
		*a = (*a).(int) - b.(int)
	case uint:
		*a = (*a).(uint) - b.(uint)
	case float32:
		*a = (*a).(float32) - b.(float32)
	case float64:
		*a = (*a).(float64) - b.(float64)
	default:
		panic("Sub(): can't get a - b")
	}

}

// a * b
func Mult(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Mult(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) * b.(int8)
	case int16:
		*a = (*a).(int16) * b.(int16)
	case int32:
		*a = (*a).(int32) * b.(int32)
	case int64:
		*a = (*a).(int64) * b.(int64)
	case uint8:
		*a = (*a).(uint8) * b.(uint8)
	case uint16:
		*a = (*a).(uint16) * b.(uint16)
	case uint32:
		*a = (*a).(uint32) * b.(uint32)
	case uint64:
		*a = (*a).(uint64) * b.(uint64)
	case int:
		*a = (*a).(int) * b.(int)
	case uint:
		*a = (*a).(uint) * b.(uint)
	case float32:
		*a = (*a).(float32) * b.(float32)
	case float64:
		*a = (*a).(float64) * b.(float64)
	default:
		panic("Mult(): can't get a * b")
	}
}

// a / b
func Del(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Del(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) / b.(int8)
	case int16:
		*a = (*a).(int16) / b.(int16)
	case int32:
		*a = (*a).(int32) / b.(int32)
	case int64:
		*a = (*a).(int64) / b.(int64)
	case uint8:
		*a = (*a).(uint8) / b.(uint8)
	case uint16:
		*a = (*a).(uint16) / b.(uint16)
	case uint32:
		*a = (*a).(uint32) / b.(uint32)
	case uint64:
		*a = (*a).(uint64) / b.(uint64)
	case int:
		*a = (*a).(int) / b.(int)
	case uint:
		*a = (*a).(uint) / b.(uint)
	case float32:
		*a = (*a).(float32) / b.(float32)
	case float64:
		*a = (*a).(float64) / b.(float64)
	default:
		panic("Del(): can't get a / b")
	}
}

// a % b
func Mod(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Mod(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) % b.(int8)
	case int16:
		*a = (*a).(int16) % b.(int16)
	case int32:
		*a = (*a).(int32) % b.(int32)
	case int64:
		*a = (*a).(int64) % b.(int64)
	case uint8:
		*a = (*a).(uint8) % b.(uint8)
	case uint16:
		*a = (*a).(uint16) % b.(uint16)
	case uint32:
		*a = (*a).(uint32) % b.(uint32)
	case uint64:
		*a = (*a).(uint64) % b.(uint64)
	case int:
		*a = (*a).(int) % b.(int)
	case uint:
		*a = (*a).(uint) % b.(uint)
	default:
		panic("Mod(): can't get a % b")
	}
}

// a ^ b
func Xor(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Xor(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) ^ b.(int8)
	case int16:
		*a = (*a).(int16) ^ b.(int16)
	case int32:
		*a = (*a).(int32) ^ b.(int32)
	case int64:
		*a = (*a).(int64) ^ b.(int64)
	case uint8:
		*a = (*a).(uint8) ^ b.(uint8)
	case uint16:
		*a = (*a).(uint16) ^ b.(uint16)
	case uint32:
		*a = (*a).(uint32) ^ b.(uint32)
	case uint64:
		*a = (*a).(uint64) ^ b.(uint64)
	case int:
		*a = (*a).(int) ^ b.(int)
	case uint:
		*a = (*a).(uint) ^ b.(uint)
	default:
		panic("Xor(): can't get a ^ b")
	}
}

// a | b
func Or(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Or(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) | b.(int8)
	case int16:
		*a = (*a).(int16) | b.(int16)
	case int32:
		*a = (*a).(int32) | b.(int32)
	case int64:
		*a = (*a).(int64) | b.(int64)
	case uint8:
		*a = (*a).(uint8) | b.(uint8)
	case uint16:
		*a = (*a).(uint16) | b.(uint16)
	case uint32:
		*a = (*a).(uint32) | b.(uint32)
	case uint64:
		*a = (*a).(uint64) | b.(uint64)
	case int:
		*a = (*a).(int) | b.(int)
	case uint:
		*a = (*a).(uint) | b.(uint)
	default:
		panic("Or(): can't get a | b")
	}
}

// a & b
func And(a *interface{}, b interface{}) {
	typeOfA := fmt.Sprintf("%T", *a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("Mod(): types of a and b must be same")
	}
	switch (*a).(type) {
	case int8:
		*a = (*a).(int8) & b.(int8)
	case int16:
		*a = (*a).(int16) & b.(int16)
	case int32:
		*a = (*a).(int32) & b.(int32)
	case int64:
		*a = (*a).(int64) & b.(int64)
	case uint8:
		*a = (*a).(uint8) & b.(uint8)
	case uint16:
		*a = (*a).(uint16) & b.(uint16)
	case uint32:
		*a = (*a).(uint32) & b.(uint32)
	case uint64:
		*a = (*a).(uint64) & b.(uint64)
	case int:
		*a = (*a).(int) & b.(int)
	case uint:
		*a = (*a).(uint) & b.(uint)
	default:
		panic("And(): can't get a & b")
	}
}

// compare a < b
func IsLess(a interface{}, b interface{}) bool {
	typeOfA := fmt.Sprintf("%T", a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("IsLess(): types of a and b must be same")
	}
	switch (a).(type) {
	case int8:
		return (a.(int8) < b.(int8))
	case int16:
		return (a.(int16) < b.(int16))
	case int32:
		return (a.(int32) < b.(int32))
	case int64:
		return (a.(int64) < b.(int64))
	case uint8:
		return (a.(uint8) < b.(uint8))
	case uint16:
		return (a.(uint16) < b.(uint16))
	case uint32:
		return (a.(uint32) < b.(uint32))
	case uint64:
		return (a.(uint64) < b.(uint64))
	case int:
		return (a.(int) < b.(int))
	case uint:
		return (a.(uint) < b.(uint))
	case string:
		return (a.(string) < b.(string))
	case float32:
		return (a.(float32) < b.(float32))
	case float64:
		return (a.(float64) < b.(float64))
	default:
		panic("IsLess(): can't compare a and b")
	}
}

// compare a == b
func IsEqual(a interface{}, b interface{}) bool {
	typeOfA := fmt.Sprintf("%T", a)
	typeOfB := fmt.Sprintf("%T", b)
	if typeOfA != typeOfB {
		panic("IsEqual(): types of a and b must be same")
	}
	switch (a).(type) {
	case int8:
		return (a.(int8) == b.(int8))
	case int16:
		return (a.(int16) == b.(int16))
	case int32:
		return (a.(int32) == b.(int32))
	case int64:
		return (a.(int64) == b.(int64))
	case uint8:
		return (a.(uint8) == b.(uint8))
	case uint16:
		return (a.(uint16) == b.(uint16))
	case uint32:
		return (a.(uint32) == b.(uint32))
	case uint64:
		return (a.(uint64) == b.(uint64))
	case int:
		return (a.(int) == b.(int))
	case uint:
		return (a.(uint) == b.(uint))
	case string:
		return (a.(string) == b.(string))
	case float32:
		return (a.(float32) == b.(float32))
	case float64:
		return (a.(float64) == b.(float64))
	default:
		panic("IsEqual(): can't compare a and b")
	}
}

// compare a <= b
func IsLessE(a interface{}, b interface{}) bool {
	return (IsLess(a, b) || IsEqual(a, b))
}

// compare a > b
func IsBig(a interface{}, b interface{}) bool {
	return (IsLessE(a, b) == false)
}

// compare a >= b
func IsBigE(a interface{}, b interface{}) bool {
	return (IsLess(a, b) == false)
}

/** EXTRA **/
//return max elemen from given array
func Max(array ...int) int {
	if len(array) == 0 {
		panic("cppf/pkg/func/c.Max(): Error: lenght of array must be greater zero")
	}
	var mxElement int = array[0]
	for _, x := range array {
		if mxElement < x {
			mxElement = x
		}
	}
	return mxElement
}

//return min elemen from given array
func Min(array ...int) int {
	if len(array) == 0 {
		panic("cppf/pkg/func/c.Min(): Error: lenght of array must be greater zero")
	}
	var mnElement int = array[0]
	for _, x := range array {
		if mnElement > x {
			mnElement = x
		}
	}
	return mnElement
}
func Sum(array ...int) int {
	if len(array) == 0 {
		panic("cppf/pkg/func/c.Sum(): Error: lenght of array must be greater zero")
	}
	var sum int = 0
	for _, x := range array {
		sum += x
	}
	return sum
}

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
		panic("went out of bounds of the array")
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
