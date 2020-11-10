package operators

import(
	"fmt"
)
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