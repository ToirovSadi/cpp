package queue

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
		panic("you are poping empty queue");
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
		panic("you getting first element from empty queue");
	}
	return q.head.data;
}
func (q Queue)Size()int{
	return (q.size);
}