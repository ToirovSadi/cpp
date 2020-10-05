package queue

type node struct{
	next *node
	data interface{}
}
type Queue struct{
	head *node
	tail *node
	size int
}
func (q *Queue)Push(v interface{}){
	var newVal = new(node);
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
		panic("cpp/pkg/types/queue.go.Pop(): you are poping empty queue");
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
		panic("cpp/pkg/types/queue.go.Front():you getting first element from empty queue");
	}
	return q.head.data;
}
func (q Queue)Size()int{
	return (q.size);
}