package assert

import(
	"os"
	"fmt"
)

//Assert if condition is false the program will print the error and exit
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