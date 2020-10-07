package c

import(
	"os"
	"fmt"
	"math"
)
		/** ASSERT **/
//if condition is false the program will print the error and exit
func Assert(cond bool, err string){
	if(cond == false){
		var res string = ("*  Error: " + err + "  *");
		for i := 0; i < len(res); i ++{
			fmt.Print("*");
		}
		fmt.Print("\n\n" + res + "\n\n");
		for i := 0; i < len(res); i ++{
			fmt.Print("*");
		}
		fmt.Print("\n");
		os.Exit(1);
	}
}


		/** MATH **/
//print sqrt of number
func Sqrt(n float64) float64{
	return math.Sqrt(n);
}
//round a number up
func Ceil(n float64) float64{
	return math.Ceil(n);
}
//round a number down to an integer
func Floor(n float64) float64{
	return math.Floor(n);
}
//round a number based on arifmetic laws
func Round(n float64) float64{
	return math.Round(n);
}
//get (a ^ n) when a and n is int
func Pow(a int64, n int64) int64{
	var res int64 = 1;
	for n > 0{
		if(n % 2 == 1){
			res = res * a;
		}
		a = a * a;
		n >>= 1;
	}
	return res;
}
//get (a ^ n) % mod when a and n is int
func Powmod(a int64, n int64, mod int64)int64{
	var res int64 = 1;
	a %= mod;
	for n > 0{
		if(n % 2 == 1){
			res = (res * a) % mod;
		}
		a = (a * a) % mod;
		n >>= 1;
	}
	return res;
}
//get (a ^ n) % mod when a is float and n is int
func Powf(a float64, n int) float64{
	if(n == 0){
		return 1;
	}
	var temp float64 = Powf(a, n / 2);
	if(n % 2 == 0){
		return temp * temp;
	}else{
		if(n > 0){
			return a * temp * temp;
		}else{
			return (temp * temp) / a;
		}
	}
}
//return sin of n
func Sin(n float64) float64{
	return math.Sin(n);
}
//return cos of n
func Cos(n float64) float64{
	return math.Cos(n);
}
//return tan of n
func Tan(n float64) float64{
	return math.Tan(n);
}
//return asin of n
func Asin(n float64) float64{
	return math.Asin(n);
}
//return acos of n
func Acos(n float64) float64{
	return math.Acos(n);
}
//return atan of n
func Atan(n float64) float64{
	return math.Atan(n);
}
//return abs of n, m
func Abs(n float64)float64{
	return math.Abs(n);
}
//return distance between two points in two-dimensional coordinate system
func Distance(x1, y1, x2, y2 float64) float64{
	return (((x2 - x1) * (x2 - x1)) + ((y2 - y1) * (y2 - y1)));
}

		/** EXTRA **/
//return max elemen from given array
func Max(array ...int) int{
	if(len(array) == 0){
		panic("cppf/pkg/func/c.Max(): Error: lenght of array must be greater zero");
	}
	var mxElement int = array[0];
	for _, x := range(array){
		if(mxElement < x){
			mxElement = x;
		}
	}
	return mxElement;
}
//return min elemen from given array
func Min(array ...int) int{
	if(len(array) == 0){
		panic("cppf/pkg/func/c.Min(): Error: lenght of array must be greater zero");
	}
	var mnElement int = array[0];
	for _, x := range(array){
		if(mnElement > x){
			mnElement = x;
		}
	}
	return mnElement;
}
func Sum(array ...int) int{
	if(len(array) == 0){
		panic("cppf/pkg/func/c.Sum(): Error: lenght of array must be greater zero");
	}
	var sum int = 0;
	for _, x := range(array){
		sum += x;
	}
	return sum;
}

		/** TYPES **/
var(
	INT_MAX int = 2147483647
	INT_MIN int = (-INT_MAX - 1)
	LLONG_MAX int64 = 9223372036854775807
	LLONG_MIN int64 = (-LLONG_MAX - 1)
)

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