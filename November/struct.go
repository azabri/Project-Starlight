package main
import (
  "fmt"
 )
 
 type person struct {
  name string
  age int
 }
 
 func main() {
  p := person{name: "Geza", age: 34}
  fmt.Println(p)
  fmt.Println(p.age)
 }
