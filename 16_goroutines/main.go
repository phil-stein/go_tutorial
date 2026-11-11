package main

import "fmt"
import "time"
import "sync"



func main(){
  fmt.Println( "--- goroutines ---" )
  
  var wg sync.WaitGroup
  time_goroutines := time.Now()
 
  go load_shaders( &wg )
  go load_assets( &wg )
  go generate_level( &wg )
  
  // wait on all wait groups to wg.Done()
  wg.Wait()

  elapsed := time.Since( time_goroutines )
  fmt.Println( "elapsed:", elapsed )

  fmt.Println( "-- normal function calls --" )
  // -- normal function calls --

  time_normal := time.Now()

  load_shaders( &wg )
  load_assets( &wg )
  generate_level( &wg )

  elapsed = time.Since( time_normal )
  fmt.Println( "elapsed:", elapsed )

}

func load_shaders( wg *sync.WaitGroup ){
  (*wg).Add(1)
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
func load_assets( wg *sync.WaitGroup ){
  (*wg).Add(1)
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
func generate_level( wg *sync.WaitGroup ){
  (*wg).Add(1)
  time.Sleep( time.Duration( 2000 ) * time.Millisecond )
  (*wg).Done()
}
