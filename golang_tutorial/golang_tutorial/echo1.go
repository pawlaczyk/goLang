//Echo1 wyświetla swoje argumenty wiersza poleceń
package main

import(
	"fmt"
	"os"
)
func main(){
	var s, sep string
	for i :=1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = "\n"
	}
	fmt.Println(s)

}