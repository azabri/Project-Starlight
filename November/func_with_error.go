package main
import (
  "fmt"
  "math"
  "errors"
 )
 
 func main() {
  result, err := sqrt(16)
  
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Println(result)
  }
 }
 
 func sqrt(x float64) (float64, error) {
  if x < 0 {
    return 0, errors.New("Cannot be negativ!")
  }
    return math.Sqrt(x), nil
 }
