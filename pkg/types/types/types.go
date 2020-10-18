package types

var(
	INT_MAX int = 2147483647
	INT_MIN int = (-INT_MAX - 1)
	LLONG_MAX int64 = 9223372036854775807
	LLONG_MIN int64 = (-LLONG_MAX - 1)
)
type Pair struct{
	First interface{}
	Second interface{}
}

type Tuple struct{
	sliceTuple []interface{}
}
func (t *Tuple)Init(array ...interface{}){
	t.sliceTuple = array;
}
func (t Tuple)Get(index int)interface{}{
	if(index >= len(t.sliceTuple)){
		panic("went out of bounds of the array");
	}
	return t.sliceTuple[index];
}