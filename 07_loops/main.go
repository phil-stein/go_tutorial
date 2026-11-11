package main


import "fmt"


func main(){

  fmt.Println( "--- for loops ---" )
  // --- for loops ---

  // -- "while" loops --
  // runs as long as condition is met
  // condition could be a boolean or idk like x <= 10

  var is_game_running bool = true

  for is_game_running{
    if 0 == 0 {
      is_game_running = false
    }
  }

  // -- break statement --
  for {   // always true same as for true {}
    if 0 == 0 {
      break   // break breaks out of the loop completely
    }
  }

  // -- continue statement --

  var arr_00 = [...]int{ 0, 1, 2, 3, 4, 5, 6, 7, 8, 9 }

  for _, value := range arr_00 {

    if value <= 5 {
      continue  // continue skips to the next loop iteration
    }
    fmt.Println( "value:", value)
  }

  // -- traditional for loops --
  // first block declares a variable
  // second block defines end condition
  // last block defines what happens after every iteration

  for i := 0; i < 10; i++{
    fmt.Println( "i:", i )
  }

  // -- loop over array / slice --
  // @NOTE: is orderded from 0..<len(arr)

  var slice_00 = []int{ 3, 4, 23, 4, 65, 98 }
  for idx, value := range slice_00{
    fmt.Println( "idx:", idx, "value:", value )
  }

  // -- loop over hashmap --
  // @NOTE: ordering is random 

  var map_00 = map[string]int{ "bby":0, "alice":23, "bob":53, "jenga":36 }
  for name := range map_00{
    fmt.Println( "name:", name )
  }
  for name, value := range map_00{
    fmt.Println( "name:", name, "value:", value )
  }


}
