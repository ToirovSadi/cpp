package assert

import(
	"os"
	"fmt"
)

//if condition is false the program will print the error and exit
func Assert(cond bool, err string){
	if(cond == false){
		fmt.Println("\n\nError:", err);
		os.Exit(0);
	}
}