package main

import "fmt"
import "time"
import "sync"

// ------------------------------------------
// goroutines are function calls that may be 
// executed at the same time as one another
// ------------------------------------------


func main(){
  fmt.Println( "--- goroutines ---" )
  
  var wg sync.WaitGroup
  time_goroutines := time.Now()
 
  wg.Add(3)
  go load_shaders_goroutine( &wg )
  go load_assets_goroutine( &wg )
  go generate_level_goroutine( &wg )
  
  // wait on all wait groups to wg.Done()
  wg.Wait()

  elapsed := time.Since( time_goroutines )
  fmt.Println( "elapsed:", elapsed )

  fmt.Println( "-- normal function calls --" )
  // -- normal function calls --

  time_normal := time.Now()

  load_shaders()
  load_assets()
  generate_level()

  elapsed = time.Since( time_normal )
  fmt.Println( "elapsed:", elapsed )

}

func load_shaders_goroutine( wg *sync.WaitGroup ){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
func load_assets_goroutine( wg *sync.WaitGroup ){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
func generate_level_goroutine( wg *sync.WaitGroup ){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
func load_shaders(){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
}
func load_assets(){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
}
func generate_level(){
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
}
