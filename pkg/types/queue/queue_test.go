package queue

import(
	"testing"
)

func Benchmark(b *testing.B){
	for i := 0; i < b.N; i++ {
		queue();
	}
}
func queue(){
	var q Queue;
	const it int = 1e6;
	for i := 0; i < it; i ++{
		q.Push(i);
	}
}