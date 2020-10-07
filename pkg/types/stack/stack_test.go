package stack

import(
	"testing"
)

func BenchmarkStack(b *testing.B){
	for i := 0; i < b.N; i++{
		stack();
	}
}
func stack(){
	var st Stack;
	const it int = 1e6;
	for i := 0; i < it; i ++{
		st.Push(i);
	}
}