package main

import "fmt"


func main(){
  
  fmt.Println( "--- if statements ---" ) 

  // first look at the if statements and the accompanying? logic operators
  if_statement()

  fmt.Println() 
  fmt.Println( "--- switch statements ---" ) 

  // switch statements are realy useful, just not when youre a beginner, so you can skip these if you want
  switch_statement()

}

func if_statement(){
  var x int = 10
  var y int = 10

  // --- if statement ---
  
  // -- logic operators --

  if x == y {
    fmt.Println( "x is equal to y" )
  }
  if x != y {
    fmt.Println( "x is NOT equal to y" )
  }
  if x < y {
    fmt.Println( "x is less than y" )
  }
  if x > y {
    fmt.Println( "x is more than y" )
  }
  if x <= y {
    fmt.Println( "x is less than or equals to y" )
  }
  if x >= y {
    fmt.Println( "x is more than or equals to y" )
  }

  // -- and or operators --

  var b_00 bool = true
  var b_01 bool = false
  
  if !b_00 {
    fmt.Println( "b_00 is false, ! operator inverts boolean values" )
  }
  if b_00 && b_01 {
    fmt.Println( "b_00 and b_01 are true" )
  }
  if b_00 || b_01 {
    fmt.Println( "b_00 or b_01 is true" )
  }
  if x == 10 && y < 100 {
    fmt.Println( "x is equal to 10 AND y is less than 100" )
  }
  if x != 12 || y > 100 {
    fmt.Println( "x is not equal to 12 OR y is more than 100" )
  }

  // -- else statement --

  if b_00 {
    fmt.Println( "b_00 is true" )
  } else {
    fmt.Println( "b_00 is false" )
  }

  if x == y {
    fmt.Println( "x is equal to y" )
  } else if x == 12 {
    fmt.Println( "x is equal to 12" )
  } else {
    fmt.Println( "x is neither equal to y or 12" )
  }

  if x == y {
    fmt.Println( "x is equal to y" )
  } else if x == 12 {
    fmt.Println( "x is equal to 12" )
  } else if x == 24 {
    fmt.Println( "x is equal to 24" )
  } else if x == 69 {
    fmt.Println( "x is equal to 69" )
  } else {
    fmt.Println( "x is neither equal to y or 12 or 24 or 69" )
  }
}

func switch_statement(){

  var i_00 int = 123

  // switch statement, just basically another way of making if-else statements, 
  // but can be optimized better by compiler, is basically uninteresting for you 
  switch{
    case i_00 == 0: 
      fmt.Println( "i_00 is equal to 0" )
    case i_00 != 0:
      fmt.Println( "i_00 is NOT equal to 0" )
    case i_00 < 1234:
      fmt.Println( "i_00 is less than 1234" )
    default:
      fmt.Println( "default case, executes when no other cases execute" )
  }

  // switch on the value of the passed variable
  switch i_00{
    case 0:
      fmt.Println( "i_00 is equal to 0" )
    case 24:
      fmt.Println( "i_00 is equal to 24" )
    case 28, 123:
      fmt.Println( "i_00 is equal to 28 or 123" )
    default:
      fmt.Println( "default case, executes when no other cases execute" )
  }

}
