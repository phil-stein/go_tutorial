package main

import "fmt"


func main(){

  fmt.Println( "--- arrays ---" )  
  // --- arrays ---
  // arrays are are contigous? values of the same type with a fixed length
  // indexing is 0 based meaning an array of length 10 can be indexed with 0...9 not with 10
  // f.e. 
  //  var arr [4]int
  //  arr[0] = 10
  //  arr[3] = 12
  //  arr -> [ 10, 0, 0, 12 ]

  var i_arr_00 [12]int
  i_arr_00[0] = 123
  i_arr_00[11] = 321
  fmt.Println( "i_arr_00:", i_arr_00 )

  var s_arr_00 [3]string = [3]string{ "hello", "cock", "there" }
  fmt.Println( "a_arr_00:", s_arr_00 )

  // the compiler can also infer the length
  var i_arr_01 = [...]int{ 1, 2, 3, 4, 5, 6 }
  fmt.Println( "i_arr_01:", i_arr_01 )

  fmt.Println( "--- slices ---" )  
  // --- slices ---
  // slices are whats also called dynamic arrays
  // they can be resized and have elements added to them
  // this works by having a len() value and a cap() value
  // cap() is the actual physical capacity and len is the amount of elements

  var i_slice_00 []int = []int{ 1, 2, 3, 4 }
  fmt.Println( "i_slice_00:", i_slice_00, "len(i_slice_00):", len(i_slice_00), "cap(i_slice_00):", cap(i_slice_00) )
  i_slice_00 = append( i_slice_00, 5 )
  fmt.Println( "i_slice_00:", i_slice_00, "len(i_slice_00):", len(i_slice_00), "cap(i_slice_00):", cap(i_slice_00) )

  var i_slice_01 []int = make( []int, 5 )     // len(): 5, cap(): 5
  var i_slice_02 []int = make( []int, 5, 8 )  // len(): 5, cap(): 8
  fmt.Println( "i_slice_01:", i_slice_01, "len(i_slice_01):", len(i_slice_01), "cap(i_slice_01):", cap(i_slice_01) )
  fmt.Println( "i_slice_02:", i_slice_02, "len(i_slice_02):", len(i_slice_02), "cap(i_slice_02):", cap(i_slice_02) )

  i_slice_01[0] = 123
  i_slice_01[1] = 123
  i_slice_01[2] = 123
  i_slice_01[3] = 123
  i_slice_01[4] = 123

  // append slice to a slice
  i_slice_00 = append( i_slice_00, i_slice_01... )
  fmt.Println( "i_slice_00:", i_slice_00, "len(i_slice_00):", len(i_slice_00), "cap(i_slice_00):", cap(i_slice_00) )


  fmt.Println( "--- maps ---" )  
  // --- maps ---

  var map_00 map[string]int = make( map[string]int )
  fmt.Println( "map_00:", map_00 )
  
  var map_01 = map[string]int{ "key":123, "cock":69, "hello":23 }
  fmt.Println( "map_01:", map_01 )


  // getting the value via the key
  // dont do this
  var value_00 = map_01["key"]  // value is 123
  fmt.Println( "value_00:", value_00 )
  // do this instead, ok value returns whether key was found in map
  var value_01, ok = map_01["unknown"]  // value is 0 bc. key is invalid, ok is false
  if !ok{
    value_01 = 123
  }
  fmt.Println( "value_01:", value_01, "ok:", ok )

  delete( map_01, "cock" )
  fmt.Println( "map_01:", map_01 )

  // @TODO: appending to maps ???

}
