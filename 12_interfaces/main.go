package main

import "fmt"

// --------------------------------------------------------------------
// disclaimer i find these interfaces unnecessary and usually annoying 
// added this mostly here so you can look up how they work 
// if you come across one on the internet or something
// --------------------------------------------------------------------


type gas_engine_t struct{
  fuel_liter   float32
  liter_per_km float32
}

type electric_engine_t struct{
  fuel_kw   float32
  kw_per_km float32
}

func (e gas_engine_t) km_abled_to_drive() int {
  return int( e.fuel_liter * e.fuel_per_km )
}

func (e electric_engine_t) km_abled_to_drive() int {
  return int( e.fuel_kw * e.kw_per_km )
}

type engine_t interface{
  km_abled_to_drive() int
}

func can_make_it( e engine_t, km int ) {
  return e.km_abled_to_drive() >= km
}

func main(){

  fmt.Println( "--- interfaces ---" )

  var e electric_engine_t
  e.fuel_kw   = 2000
  e.kw_per_km = 100
  fmt.Println( "can_make_it( e, 20 ):", can_make_it( e, 20 )  )

  var g gas_engine_t
  g.fuel_liter   = 80
  g.liter_per_km = 13.5
  fmt.Println( "can_make_it( g, 20 ):", can_make_it( g, 20 )  )

}
