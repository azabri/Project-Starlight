package main
import(
  "fmt"
  "time"
 )
 
 func main() {
  fmt.Println("When is Saturday?")
  today := time.Now().Weekday()
  switch time.Saturday {
  case today + 0:
    fmt.Println("It's today!")
   case today + 1:
    fmt.Println("Tomorrow!")
   case today + 2:
    fmt.Println("In two days")
   default:
    fmt.Println("Too far...!")
    
    t := time.Now()
    switch {
    case t.Hour() < 12:
      fmt.Println("Good morning!")
    case t.Hour() < 17:
      fmt.Println("Good afternoon!")
    default:
      fmt.Println("Good evening!")
    }
  }
 }
