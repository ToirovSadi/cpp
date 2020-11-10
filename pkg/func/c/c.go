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
		fmt.Print("\n*\n" + res + "\n*\n");
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
		/** OPERATORS **/
// a + b
func Add(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Add(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) + b.(int8);
		case int16:
			*a = (*a).(int16) + b.(int16);
		case int32:
			*a = (*a).(int32) + b.(int32);
		case int64:
			*a = (*a).(int64) + b.(int64);
		case uint8:
			*a = (*a).(uint8) + b.(uint8);
		case uint16:
			*a = (*a).(uint16) + b.(uint16);
		case uint32:
			*a = (*a).(uint32) + b.(uint32);
		case uint64:
			*a = (*a).(uint64) + b.(uint64);
		case int:
			*a = (*a).(int) + b.(int);
		case uint:
			*a = (*a).(uint) + b.(uint);
		case string:
			*a = (*a).(string) + b.(string);
		case float32:
			*a = (*a).(float32) + b.(float32);
		case float64:
			*a = (*a).(float64) + b.(float64);
		default:
			panic("Add(): can't get a + b");
	}
}
// a - b
func Sub(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Sub(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) - b.(int8);
		case int16:
			*a = (*a).(int16) - b.(int16);
		case int32:
			*a = (*a).(int32) - b.(int32);
		case int64:
			*a = (*a).(int64) - b.(int64);
		case uint8:
			*a = (*a).(uint8) - b.(uint8);
		case uint16:
			*a = (*a).(uint16) - b.(uint16);
		case uint32:
			*a = (*a).(uint32) - b.(uint32);
		case uint64:
			*a = (*a).(uint64) - b.(uint64);
		case int:
			*a = (*a).(int) - b.(int);
		case uint:
			*a = (*a).(uint) - b.(uint);
		case float32:
			*a = (*a).(float32) - b.(float32);
		case float64:
			*a = (*a).(float64) - b.(float64);
		default:
			panic("Sub(): can't get a - b");
	}

}
// a * b
func Mult(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Mult(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) * b.(int8);
		case int16:
			*a = (*a).(int16) * b.(int16);
		case int32:
			*a = (*a).(int32) * b.(int32);
		case int64:
			*a = (*a).(int64) * b.(int64);
		case uint8:
			*a = (*a).(uint8) * b.(uint8);
		case uint16:
			*a = (*a).(uint16) * b.(uint16);
		case uint32:
			*a = (*a).(uint32) * b.(uint32);
		case uint64:
			*a = (*a).(uint64) * b.(uint64);
		case int:
			*a = (*a).(int) * b.(int);
		case uint:
			*a = (*a).(uint) * b.(uint);
		case float32:
			*a = (*a).(float32) * b.(float32);
		case float64:
			*a = (*a).(float64) * b.(float64);
		default:
			panic("Mult(): can't get a * b");
	}
}
// a / b
func Del(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Del(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) / b.(int8);
		case int16:
			*a = (*a).(int16) / b.(int16);
		case int32:
			*a = (*a).(int32) / b.(int32);
		case int64:
			*a = (*a).(int64) / b.(int64);
		case uint8:
			*a = (*a).(uint8) / b.(uint8);
		case uint16:
			*a = (*a).(uint16) / b.(uint16);
		case uint32:
			*a = (*a).(uint32) / b.(uint32);
		case uint64:
			*a = (*a).(uint64) / b.(uint64);
		case int:
			*a = (*a).(int) / b.(int);
		case uint:
			*a = (*a).(uint) / b.(uint);
		case float32:
			*a = (*a).(float32) / b.(float32);
		case float64:
			*a = (*a).(float64) / b.(float64);
		default:
			panic("Del(): can't get a / b");
	}
}
// a % b
func Mod(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Mod(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) % b.(int8);
		case int16:
			*a = (*a).(int16) % b.(int16);
		case int32:
			*a = (*a).(int32) % b.(int32);
		case int64:
			*a = (*a).(int64) % b.(int64);
		case uint8:
			*a = (*a).(uint8) % b.(uint8);
		case uint16:
			*a = (*a).(uint16) % b.(uint16);
		case uint32:
			*a = (*a).(uint32) % b.(uint32);
		case uint64:
			*a = (*a).(uint64) % b.(uint64);
		case int:
			*a = (*a).(int) % b.(int);
		case uint:
			*a = (*a).(uint) % b.(uint);
		default:
			panic("Mod(): can't get a % b");
	}
}
// a ^ b
func Xor(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Xor(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) ^ b.(int8);
		case int16:
			*a = (*a).(int16) ^ b.(int16);
		case int32:
			*a = (*a).(int32) ^ b.(int32);
		case int64:
			*a = (*a).(int64) ^ b.(int64);
		case uint8:
			*a = (*a).(uint8) ^ b.(uint8);
		case uint16:
			*a = (*a).(uint16) ^ b.(uint16);
		case uint32:
			*a = (*a).(uint32) ^ b.(uint32);
		case uint64:
			*a = (*a).(uint64) ^ b.(uint64);
		case int:
			*a = (*a).(int) ^ b.(int);
		case uint:
			*a = (*a).(uint) ^ b.(uint);
		default:
			panic("Xor(): can't get a ^ b");
	}
}
// a | b
func Or(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Or(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) | b.(int8);
		case int16:
			*a = (*a).(int16) | b.(int16);
		case int32:
			*a = (*a).(int32) | b.(int32);
		case int64:
			*a = (*a).(int64) | b.(int64);
		case uint8:
			*a = (*a).(uint8) | b.(uint8);
		case uint16:
			*a = (*a).(uint16) | b.(uint16);
		case uint32:
			*a = (*a).(uint32) | b.(uint32);
		case uint64:
			*a = (*a).(uint64) | b.(uint64);
		case int:
			*a = (*a).(int) | b.(int);
		case uint:
			*a = (*a).(uint) | b.(uint);
		default:
			panic("Or(): can't get a | b");
	}
}
// a & b
func And(a *interface{}, b interface{}){
	typeOfA := fmt.Sprintf("%T", *a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("Mod(): types of a and b must be same");
	}
	switch (*a).(type){
		case int8:
			*a = (*a).(int8) & b.(int8);
		case int16:
			*a = (*a).(int16) & b.(int16);
		case int32:
			*a = (*a).(int32) & b.(int32);
		case int64:
			*a = (*a).(int64) & b.(int64);
		case uint8:
			*a = (*a).(uint8) & b.(uint8);
		case uint16:
			*a = (*a).(uint16) & b.(uint16);
		case uint32:
			*a = (*a).(uint32) & b.(uint32);
		case uint64:
			*a = (*a).(uint64) & b.(uint64);
		case int:
			*a = (*a).(int) & b.(int);
		case uint:
			*a = (*a).(uint) & b.(uint);
		default:
			panic("And(): can't get a & b");
	}
}
// compare a < b
func IsLess(a interface{}, b interface{})bool{
	typeOfA := fmt.Sprintf("%T", a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("IsLess(): types of a and b must be same");
	}
	switch (a).(type){
		case int8:
			return (a.(int8) < b.(int8));
		case int16:
			return (a.(int16) < b.(int16));
		case int32:
			return (a.(int32) < b.(int32));
		case int64:
			return (a.(int64) < b.(int64));
		case uint8:
			return (a.(uint8) < b.(uint8));
		case uint16:
			return (a.(uint16) < b.(uint16));
		case uint32:
			return (a.(uint32) < b.(uint32));
		case uint64:
			return (a.(uint64) < b.(uint64));
		case int:
			return (a.(int) < b.(int));
		case uint:
			return (a.(uint) < b.(uint));
		case string:
			return (a.(string) < b.(string));
		case float32:
			return (a.(float32) < b.(float32));
		case float64:
			return (a.(float64) < b.(float64));
		default:
			panic("IsLess(): can't compare a and b");
	}
}
// compare a == b
func IsEqual(a interface{}, b interface{})bool{
	typeOfA := fmt.Sprintf("%T", a);
	typeOfB := fmt.Sprintf("%T", b);
	if(typeOfA != typeOfB){
		panic("IsEqual(): types of a and b must be same");
	}
	switch (a).(type){
		case int8:
			return (a.(int8) == b.(int8));
		case int16:
			return (a.(int16) == b.(int16));
		case int32:
			return (a.(int32) == b.(int32));
		case int64:
			return (a.(int64) == b.(int64));
		case uint8:
			return (a.(uint8) == b.(uint8));
		case uint16:
			return (a.(uint16) == b.(uint16));
		case uint32:
			return (a.(uint32) == b.(uint32));
		case uint64:
			return (a.(uint64) == b.(uint64));
		case int:
			return (a.(int) == b.(int));
		case uint:
			return (a.(uint) == b.(uint));
		case string:
			return (a.(string) == b.(string));
		case float32:
			return (a.(float32) == b.(float32));
		case float64:
			return (a.(float64) == b.(float64));
		default:
			panic("IsEqual(): can't compare a and b");
	}
}
// compare a <= b
func IsLessE(a interface{}, b interface{})bool{
	return (IsLess(a, b) || IsEqual(a, b));
}
// compare a > b
func IsBig(a interface{}, b interface{})bool{
	return (IsLessE(a, b) == false);
}
// compare a >= b
func IsBigE(a interface{}, b interface{})bool{
	return (IsLess(a, b) == false);
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