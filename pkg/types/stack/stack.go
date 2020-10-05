package stack

type node struct{
	next *node
	prev *node
	data interface{}
}
type Stack struct{
	head *node
	tail *node
	size int
}
func (s *Stack)Push(v interface{}){
	var newVal = new(node);
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
		panic("cpp/pkg/types/queue.go.Pop(): you are poping empty queue");
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
		panic("cpp/pkg/types/queue.go.Front():you getting first element from empty queue");
	}
	return s.tail.data;
}
func (s Stack)Size()int{
	return (s.size);
}