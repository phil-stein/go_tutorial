package main


import "fmt"
import "math/rand" // rand.Float32() -> random number between 0-1 


func main(){

  const MAX_QUESTIONS int = 3
  var q float32 = rand.Float32() * float32(MAX_QUESTIONS)

  if int(q) == 0 {
    fmt.Println( "[0] whats the capital of germany ?" )
  } else if int(q) == 1{
    fmt.Println( "[1] whats 2 * 4 + 8 ?" )
  } else if int(q) == 2{
    fmt.Println( "[2] whats H on the periodic table ?" )
  } else {
    fmt.Println( "[ERROR] q:", int(q), "is bigger than the amount of questions !!!" ) 
  }

  fmt.Print( "Enter Answer:" )
  var answer string
  _, err := fmt.Scanln( &answer ) 
  if err != nil {
    fmt.Println( "[ERROR] failed reading input !!!" )
    return
  }
  fmt.Println( "answer:", answer )

  var answer_correct bool = false
  switch int(q) {
    case 0: answer_correct = answer == "berlin" || answer == "Berlin" || answer  == "BERLIN"
    case 1: answer_correct = answer == "16"
    case 2: answer_correct = answer == "hydrogen" || answer == "Hydrogen" || answer  == "HYDROGEN"
    default: fmt.Println( "[ERROR] q:", int(q), "is bigger than the amount of questions !!!" ) 
  }

  fmt.Println( "answer_correct:", answer_correct )

}
