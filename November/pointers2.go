package main
import (
	"fmt"
)

func main() {
	i, j := 42, 2701
	p := &i // point to i
	fmt.Println(*p) // read i through the pointer
	
	*p = 21 // set i through pointer
	fmt.Println(i) // prints 21
	
	p = &j // point to j
	fmt.Println(p)
	fmt.Println(*p)
	*p = *p / 37 // divide j though the pointer
	fmt.Println(j)
}
