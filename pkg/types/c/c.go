package c

		/** TYPES **/
var(
	INT_MAX int = 2147483647
	INT_MIN int = (-INT_MAX - 1)
	LLONG_MAX int64 = 9223372036854775807
	LLONG_MIN int64 = (-LLONG_MAX - 1)
)
type Pair struct{
	First interface{}
	Second interface{}
}

type Tuple struct{
	sliceTuple []interface{}
}
func (t *Tuple)Init(array ...interface{}){
	t.sliceTuple = array;
}
func (t Tuple)Get(index int)interface{}{
	if(index >= len(t.sliceTuple)){
		panic("went out of bounds of the array");
	}
	return t.sliceTuple[index];
}

		/** QUEUE **/

type nodeForQueue struct{
	next *nodeForQueue
	data interface{}
}
type Queue struct{
	head *nodeForQueue
	tail *nodeForQueue
	size int
}
func (q *Queue)Push(v interface{}){
	var newVal = new(nodeForQueue);
	newVal.next = nil;
	newVal.data = v;
	if(q.head == nil){
		q.head = newVal;
		q.tail = newVal;
	}else{
		q.tail.next = newVal;
		q.tail = newVal;
	}
	q.size ++;
}
func (q Queue)Empty()bool{
	return (q.size == 0);
}
func (q *Queue)Pop(){
	if(q.Empty() == true){
		panic("cpp/pkg/types/c.go.Pop():you are poping empty queue");
	}
	if(q.head.next == nil){
		q.head = nil;
		q.tail = nil;
	}else{
		q.head = q.head.next;
	}
	q.size --;
}
func (q Queue)Front()interface{}{
	if(q.Empty() == true){
		panic("cpp/pkg/types/c.go.Front():you getting first element from empty queue");
	}
	return q.head.data;
}
func (q Queue)Size()int{
	return (q.size);
}

		/** STACK **/

type nodeForStack struct{
	next *nodeForStack
	prev *nodeForStack
	data interface{}
}
type Stack struct{
	head *nodeForStack
	tail *nodeForStack
	size int
}
func (s *Stack)Push(v interface{}){
	var newVal = new(nodeForStack);
	newVal.next = nil;
	newVal.data = v;
	if(s.head == nil){
		s.head = newVal;
		s.tail = newVal;
		s.tail.prev = nil;
	}else{
		s.tail.next = newVal;
		last := s.tail;
		s.tail = newVal;
		s.tail.prev = last;
	}
	s.size ++;
}
func (s Stack)Empty()bool{
	return (s.size == 0);
}
func (s *Stack)Pop(){
	if(s.Empty() == true){
		panic("cpp/pkg/types/stack.go.Pop(): you are poping empty stack");
	}
	if(s.tail.prev == nil){
		s.head = nil;
		s.tail = nil;
	}else{
		s.tail = s.tail.prev;
	}
	s.size --;
}
func (s Stack)Top()interface{}{
	if(s.Empty() == true){
		panic("cpp/pkg/types/stack.go.Top():you getting first element from empty stack");
	}
	return s.tail.data;
}
func (s Stack)Size()int{
	return (s.size);
}
