package main

import "fmt"

// --------------------------------------------------------
// pointers are super useful and way easier to understand 
// than people make them out to be
// hope this is a good explanaition lol
// --------------------------------------------------------

// a pointer is just a variable whose value 
// is the memory adress of another variable
// f.e. 
//    var value    int = 10
//    var pointer *int = &value
//    *pointer += 1
// 
// *type      -> its a pointer type
// &variable  -> makes a pointer to variable
// *pointer   -> dereferences the pointer, explained below

type example_t struct{
  value int
}

func (e example_t) increment_value(){
  e.value += 1
}

func main(){
  fmt.Println( "--- pointers ---" )
  // --- pointers ---
  // pointers are all initially nil until you set them to be the adress of another variable
  // while there nil you cant dereference them or you will get a runtime exception
  var ptr_00 *int 
  fmt.Println( "ptr_00:", ptr_00, "- pointers are all initialized to be nil" )

  var val_00 int = 10

  // ptr_00 now points to val_00
  ptr_00 = &val_00
  fmt.Println( "ptr_00:", ptr_00, "- pointers now has memory adress of val_00 as value" )

  fmt.Println( "-- dereferencing --" )
  // -- dereferencing --
  // dereferencing means getting the value stored at the pointers location
  fmt.Println( "*ptr_00:", *ptr_00 )  // this prints 10, the value of val_00

  // this adds to the variable the pointer points to
  *ptr_00 += 1
  fmt.Println( "*ptr_00:", *ptr_00 )  // this prints 11, the value of val_00
  fmt.Println( "val_00:", val_00 )  

  fmt.Println( "- dereferencing method calls -" )
  // - dereferencing method calls -
  var e example_t 
  e.value = 10
  var e_ptr *example_t = &e
  (*e_ptr).increment_value()  // use (*pointer).method_call() syntax
  fmt.Println( "e:", e )

  // this becomes especiallly useful with functions

  fmt.Println( "- function without pointers -" )
  var val_01 int = 10
  fmt.Println( "val_01:", val_01, "<- pre function call" )  
  int_increment( val_01 )
  fmt.Println( "val_01:", val_01, "<- post function call" )  
  
  // val_01 stays at 10 even though its incremented in the function
  // bc. only a copy was passed to the function not the exact variable 
  // if we instead pass a pointer we can manipulate val_01 all we want
   
  fmt.Println( "- function with pointers -" )
  fmt.Println( "val_01:", val_01, "<- pre function call" )  
  int_increment_ptr( &val_01 )
  fmt.Println( "val_01:", val_01, "<- post function call" )  

  fmt.Println( "-- slices --" )
  // -- slices --

  var slice_00 = []int{ 1, 2, 3, 4, 5 }
  var slice_01 = slice_00
  fmt.Println( "slice_00:", slice_00 )
  fmt.Println( "slice_01:", slice_01 )
  slice_01[0] = 69
  fmt.Println( "after -> slice_01[0] = 69" )
  fmt.Println( "slice_00:", slice_00 )
  fmt.Println( "slice_01:", slice_01 )

  // both slices now have [0] set to 69
  // this is bc. slices are just pointers to
  // underlying arrays

}

func int_increment( i int ){
  i += 1
}
func int_increment_ptr( i *int ){
  *i += 1
}
