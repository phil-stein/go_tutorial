package main

import "fmt"
import "strings"
import "unicode/utf8"

// ------------------------------------------------------------------------------------------
// not hyper critical to understand right away, just more useful as a resource in the future
// ------------------------------------------------------------------------------------------


func main(){
  var str_00 string = "résumé"
  
  fmt.Println( "--- indexing ---" )
  // --- indexing ---

  fmt.Println( "str_00:", str_00 )
  fmt.Printf( "str_00[0]: %v, type of str_00[0]: %T\n", str_00[0], str_00[0] )

  // bc. of utf8 encoding characters 1 & 5 (é) are 2 bytes instead of one
  // watch 27:10 - 29:00 of https://www.youtube.com/watch?v=8uiZC0l4Ajw for better explanation
  // going over str_00 character by character, characters in go are called runes
  for idx, value := range str_00 {
    fmt.Println( "idx:", idx, "value:", value )
  }

  fmt.Println( "--- length ---" )
  // --- length ---

  fmt.Println( "len(str_00):", len(str_00), "<- length in bytes, 2 longer because both é have an extra byte" )
  fmt.Println( "utf8.RuneCountInString(str_00):", utf8.RuneCountInString(str_00), "<- length in runes" )  // length in character/runes

  fmt.Println( "--- runes ---" )
  // --- runes ---
  // runes are basically int32's which store the utf8 encoding of characters

  var rune_00 rune = 'a'
  var rune_01 rune = 'á'
  fmt.Println( "rune_00:", rune_00, "rune_01:", rune_01 )

  // slice of runes
  var rune_slice_00 = []rune("le cöck")
  fmt.Println( "rune_slice_00:", rune_slice_00 )

  fmt.Println( "--- concatenation ---" )
  // --- concatenation ---

  var str_01 string = "cock"
  var str_02 string = "hello"
  var str_03 string = str_02 + " " + str_01
  fmt.Println( "str_03:", str_03 )

  fmt.Println( "--- string builder ---" )
  // --- string builder ---
  // string builders allow you to compose strings on the fly 
  // and they are super useful

  var str_builder strings.Builder
  str_builder.WriteString( "hello" )
  str_builder.WriteString( " " )
  str_builder.WriteString( str_01 )
  var str_04 string = str_builder.String()
  fmt.Println( "str_04:", str_04 )

}
