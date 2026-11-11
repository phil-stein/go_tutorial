package main

import "fmt"

type enum_type int

const (
  Enum_One = iota // -> 0
  Enum_Two        // -> 1
  Enum_Three      // -> 2
  Enum_Four       // -> 3
  Enum_Max        // -> 4
)

func main(){

  fmt.Println( "Enum_One:  ", Enum_One )
  fmt.Println( "Enum_Two:  ", Enum_Two )
  fmt.Println( "Enum_Three:", Enum_Three )
  fmt.Println( "Enum_Four: ", Enum_Four )
  fmt.Println( "Enum_Max:  ", Enum_Max )

  var enum enum_type = Enum_One
  fmt.Println( "enum:", enum )
  enum = 2
  fmt.Println( "enum:", enum )

  if enum == Enum_Three {
    fmt.Println( "enum == Enum_Three" )
  }

  switch enum{
    case Enum_One:   fmt.Println( "case Enum_One" )
    case Enum_Two:   fmt.Println( "case Enum_Two" )
    case Enum_Three: fmt.Println( "case Enum_Three" )
    case Enum_Four:  fmt.Println( "case Enum_Four" )
    case Enum_Max:   fmt.Println( "case Enum_Max" )
    default:         fmt.Println( "default -> enum is out of bounds" )
  }

}
