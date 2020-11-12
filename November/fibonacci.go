package main
import (
  "fmt"
 )
 
 func main() {
 
  var a int = 0
  var b int = 1
  var c int = 0
  
    for c < 11 {
      c = a + b
      fmt.Println(c)
      a = b
      b = c
    }
 
 }
