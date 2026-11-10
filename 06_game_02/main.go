package main


import "fmt"
import "math/rand" // rand.Float32() -> random number between 0-1 


func main(){

  var questions_arr = [...]string{ "whats the capital of germany ?", "whats 2 * 4 + 8 ?", "whats H on the periodic table ?", "whats the name of this programming language ?", "whats the maximum value of uint8 ?", "whats the maximum value of int8 ?", "whats the minimum value of int8 ?" }
  var answers_arr   = [...]string{ "berlin", "16", "hydrogen", "go", "255", "127", "-128" }
  if len(questions_arr) != len(answers_arr) {
    fmt.Println( "[ERROR] questions and answers arrays are different length !!!" )
    return
  }
  var q_idx int = int( rand.Float32() * float32(len(questions_arr)) )

  fmt.Println( "answers all lower case !!!" )

  fmt.Printf( "[%v] %v\n", q_idx, questions_arr[q_idx] )

  fmt.Print( "Enter Answer:" )
  var answer string
  _, err := fmt.Scanln( &answer )   // & symbol will be explained in later section, _ means ignore result
  if err != nil {
    fmt.Println( "[ERROR] failed reading input !!!" )
    return
  }
  // fmt.Println( "answer:", answer )

  var answer_correct bool = answer == answers_arr[q_idx] 
  fmt.Println( "answer_correct:", answer_correct )

  if !answer_correct {
    fmt.Println( "answer is:", answers_arr[q_idx] )
  }

}
