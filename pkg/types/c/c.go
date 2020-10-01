package c

		/** TYPES **/
var(
	INT_MAX int = 2147483647
	INT_MIN int = (-INT_MAX - 1)
	LLONG_MAX int64 = 9223372036854775807
	LLONG_MIN int64 = (-LLONG_MAX - 1)
)

		/** QUEUE **/

type Node struct{
	next *Node
	data interface{}
}
type Queue struct{
	head *Node
	tail *Node
	size int
}
func (q *Queue)Push(v interface{}){
	var newVal = new(Node);
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