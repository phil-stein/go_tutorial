package main

import "fmt"
import "unicode/utf8"

// --------------------------------------------------------------
// variables are basically a box with a certain item inside
// the type of item is always the same but the value can change
// f.e. the number type `int` can be any number between 
//      -9,223,372,036,854,775,808 to 9,223,372,036,854,775,807
//      but it cant be a floating point number 
// --------------------------------------------------------------


func main(){

  fmt.Println( "--- declarations ---" )
  // --- declaration --- 

  var inum int = 123456789
  fmt.Println( "inum:", inum )
  var inum_small int16 = 255
  fmt.Println( "inum_small:", inum_small )

  var fnum float32 = 0.123456
  fmt.Println( "fnum:", fnum )
  var fnum_big float64 = 0.1234567891234567
  fmt.Println( "fnum_big:", fnum_big )
  
  fmt.Println( "--- maths ---" )
  // --- maths ---
  // all signed integer types are: int, int8, int16, int32, int64
  
  var x int = 123
  var y int = 2

  var res_00 int = x + y  // plus
  var res_01 int = x - y  // minus
  var res_02 int = x * y  // mal
  var res_03 int = x / y  // geteilt
  var res_04 int = x % y  // geteilt - rest (modulo)
  fmt.Println( "res_00:", res_00 )
  fmt.Println( "res_01:", res_01 )
  fmt.Println( "res_02:", res_02 )
  fmt.Println( "res_03:", res_03 )
  fmt.Println( "res_04:", res_04 )

  var num int = 123

  fmt.Println( "--- shorthands ---" )
  // -- shorthands --

  num++     // equivalent to num = num + 1 
  num--     // equivalent to num = num - 1 
  num *= 2  // equivalent to num = num * 2
  num += 2  // equivalent to num = num + 2
  num -= 2  // equivalent to num = num - 2
  num /= 2  // equivalent to num = num / 2

  fmt.Println( "--- type conversion ---" )
  // --- type conversion ---

  var i int     = 123
  var f float32 = 0.123

  var res_05 float32 = float32(i) + f
  fmt.Println( "res_05:", res_05 )

  fmt.Println( "--- strings ---" )
  // --- strings, aka text ---

  var str_00 string = "hello there "
  var str_01 string = "hello \nthere"
  var str_02 string = `hello
there`
  var str_03 string = "hello" + " " + "there"
  fmt.Println( "str_00:", str_00 )
  fmt.Println( "str_01:", str_01 )
  fmt.Println( "str_02:", str_02 )
  fmt.Println( "str_03:", str_03 )

  fmt.Println( "len(str_00):", len(str_00) )  // length in bytes, some characters may be more than 1 byte
  fmt.Println( "utf8.RuneCountInString(str_00):", utf8.RuneCountInString(str_00) )  // length in character/runes

  var rune_00 rune = 'A'
  fmt.Println( "rune_00", rune_00 )

  fmt.Println( "--- booleans ---" )
  // --- booleans ---

  var b_00 bool = true
  var b_01 bool = false
  fmt.Println( "b_00:", b_00 )
  fmt.Println( "b_01:", b_01 )

  fmt.Println( "--- default value ---" )
  // --- default value ---

  var default_i int
  var default_f float32
  var default_b bool
  var default_s string 
  fmt.Println( "default_i:", default_i )
  fmt.Println( "default_f:", default_f )
  fmt.Println( "default_b:", default_b )
  fmt.Println( "default_s:", default_s )

  fmt.Println( "--- unsigned ints ---" )
  // --- unsigned ints ---
  // unsigned integers can only have positive values no negative ones

  var u_00 uint = 123456789   // cannot be f.e. -123
  fmt.Println( "u_00:", u_00 )
  // all unsigned integer types are: uint, uint8, uint16, uint32, uint64
  
  fmt.Println( "--- type inferance ---" )
  // --- type inferance ---
  // declaring variables without specifing the type
  // prefer doing it with type declaration, bc. its easier to read

  var infered_00 = 123    // int
  var infered_01 = 0.123  // float32 or float64 not sure ???
  var infered_02 = "cock" // string
  var infered_03 = 'a'    // rune
  var infered_04 = true   // bool
  infered_05 := 123       // int
  infered_06 := 0.123     // float32 or float64 not sure ???
  infered_07 := "cock"    // string
  infered_08 := 'a'       // rune
  infered_09 := true      // bool
  fmt.Println( "infered_00:", infered_00 )
  fmt.Println( "infered_01:", infered_01 )
  fmt.Println( "infered_02:", infered_02 )
  fmt.Println( "infered_03:", infered_03 )
  fmt.Println( "infered_04:", infered_04 )
  fmt.Println( "infered_05:", infered_05 )
  fmt.Println( "infered_06:", infered_06 )
  fmt.Println( "infered_07:", infered_07 )
  fmt.Println( "infered_08:", infered_08 )
  fmt.Println( "infered_09:", infered_09 )

  fmt.Println( "--- multiple declarations ---" )
  // --- multiple declarations ---
  // !!! i recommend not doing this as its hard to read
  //     unless receiveing multiple return values from func

  var md_00, md_01, md_02, md_03 = 123, "string", true, 'a'
  md_04, md_05 := 0.123, -200
  fmt.Println( "md_00:", md_00 )
  fmt.Println( "md_01:", md_01 )
  fmt.Println( "md_02:", md_02 )
  fmt.Println( "md_03:", md_03 )
  fmt.Println( "md_04:", md_04 )
  fmt.Println( "md_05:", md_05 )


  fmt.Println( "--- constants ---" )
  // --- constants ---
  // cant change the value later in the code
  // useful when u know a value is fixed, bc. the compiler will complain when u try u change its value

  const pi float32 = 3.1415 // pi is a mathematical constant no need to change it ever
  fmt.Println( "pi:", pi )

}
