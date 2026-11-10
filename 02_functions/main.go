package main

// imports import a set of x number of functions accessible via name_of_import.name_of_function
import "fmt"
import "errors"

// when in package main you need a main function
// this function gets called automatically when your 
// program is started, so all logic starts in func main()
// and from there it can go to many places using function calls
func main(){

  // functions are called by adding () to the end
  // optionally with arguments inside

  fmt.Println( "--- print_simple() ---" )
  print_simple()

  fmt.Println( "--- arguments ---" )
  // --- arguments ---

  print_arg( "calling func arg_func()" )

  // u can use variables to pass info to functions
  var str_00 string = "cock"
  print_arg( str_00 )

  fmt.Println( "--- return values ---" )
  // --- return values ---

  var result_00 int = int_divide( 12, 6 )
  fmt.Println( "result_00:", result_00 )

  var result_01, result_02 int = int_divide_complex( 11, 2 )
  fmt.Printf( "result_01: %v, result_02: %v\n", result_01, result_02 )
  
  fmt.Println( "-- return values: ignore --" )
  // -- return values: ignore --
  // you can use _ instead of assigning to a variable
  var result_03, _ = int_divide_complex( 11, 2 )
  fmt.Println( "result_03:", result_03 )

  fmt.Println( "-- return values: errors --" )
  // -- return values: errors --

  var result_03, result_04, err = int_divide_complex_error( 123, 0 )
  if err != nil {
    fmt.Println( "[ERROR] divided by zero, err:", err.Error() )
  } else {
    fmt.Printf( "result_03: %v, result_04: %v, err: %v\n", result_03, result_04, err )
  }
}

// function without args or return value
func print_simple(){
  fmt.Println( "hello world" )
}

// function with single argument and no return value
func print_arg(str string){
  fmt.Println( str )
}

// function with two args and one int return value
func int_divide( numerator int, denomanator int ) int {

  var res int = numerator / denomanator

  return res
}

// function with two args and two int return value
func int_divide_complex( numerator int, denomanator int ) (int, int) {

  var res_div int = numerator / denomanator
  var res_mod int = numerator % denomanator

  return res_div, res_mod
}

// function with two args and two int return value and error return
func int_divide_complex_error( numerator int, denominator int ) (int, int, error) {
  
  var err error
  if denominator == 0 {
    err = errors.New( "Cannot divide by zero !!!" )

    return 0, 0, err
  }

  var res_div int = numerator / denominator
  var res_mod int = numerator % denominator

  return res_div, res_mod, nil
}

