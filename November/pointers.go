package main
import (
  "fmt"
)

func main() {
  i := 7
  fmt.Println(i)  // prints 7
  inc(&i)
  fmt.Println(i)  // prints 8
}

func inc(x *int){
  *x++
}
