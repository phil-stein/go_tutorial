package main

import "fmt"

// -----------------------------------------------------------------------------------
// structs are collections of variables that neednt be the same type as one another
// they can be created using var name struct_type like any other variable
// structs can have other structs as members inside them 
// -----------------------------------------------------------------------------------


type player_t struct{
  health     int
  max_health int
  mana       uint8
}

type enemy_t struct{
  health     int
  max_health int
  power      float32
}

type game_state_t struct{
  is_running bool
  player     player_t
  enemies    []enemy_t
}

func main(){

  fmt.Println( "--- structs ---" )
  var game_state game_state_t
  fmt.Println( "game_state:", game_state ) 
  // access individual member variables via . 
  game_state.is_running = true
  // append to slice in struct, and use struct literal
  game_state.enemies = append( game_state.enemies, enemy_t{ health:100, max_health:100, power:15 } )
  fmt.Println( "game_state:", game_state ) 

  // struct literal, i.e struct directly defined
  game_state.player = player_t{ health:20, max_health:100, mana:12 }
  fmt.Println( "game_state:", game_state ) 


  fmt.Println( "--- example usage ---" )
  // --- example usage ---

  for game_state.is_running {

    // for idx, enemy := range game_state.enemies {
    //    // do something with all enemies 
    // }

    if game_state.enemy_attack_valid() {
      game_state.player.health -= 10
      fmt.Println( "player took 10 damage, player-health:", game_state.player.health )

      if game_state.player.health <= 0{
        fmt.Println( "player died - ending game" )
        game_state.is_running = false
      }
    }
  }

}

// this defines this function to be a method of game_state_t
// therefore it can be called on instances of that type
// by using this syntax:
//  var game_state game_state_t
//  var attack_valid bool = game_state.enemy_attack_valid()
// it cant be called otherwise, you need an instance of game_state_t to call it
// this only makes sense when you have a function that operates on a struct and mostly nothing else
// @NOTE: gs could be any name, gs just standing for game_state 
func (gs game_state_t) enemy_attack_valid() bool {
 
  // could do actual logic in here

  if gs.player.health >= 0 {
    return true
  } else {
    return false
  }
}
