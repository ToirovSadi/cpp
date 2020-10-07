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