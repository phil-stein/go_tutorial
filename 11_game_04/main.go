package main

import "fmt"
import "math/rand"
import rl "github.com/gen2brain/raylib-go/raylib"


type game_state_t struct{
  current_score   int
  font_size_main  int
  font_size_score int
  window_width    int
  window_height   int
}

type entity_t struct{
  pos_x  int
  pos_y  int
  speed  int
}

type sphere_t struct{
  entity_t  // imports all members of entity_t into sphere_t
  radius float32
}
type cube_t struct{
  entity_t  // imports all members of entity_t into cube_t
  size int
}


func main() {

  // --- game state ---
  var gs game_state_t
  gs.font_size_main  = 30 
  gs.font_size_score = 60
  gs.window_width    = 1200
  gs.window_height   = 800

  rl.InitWindow(int32(gs.window_width), int32(gs.window_height), "falling cubes game window title")
  // defer puts function call at the end of a function
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)

  // --- sphere ---
  var sphere sphere_t
  sphere.speed  = 400
  sphere.radius = 30.0
  sphere.pos_x  = int(sphere.radius)
  sphere.pos_y  = gs.window_height

  // --- cube ---
  var cube cube_t
  cube.size  = 150
  cube.speed = 400
  cube.pos_y = gs.window_height - cube.size
  cube.pos_x = gs.window_width / 2 - ( cube.size / 2 )

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()

    rl.ClearBackground(rl.RayWhite)

    // --- sphere logic ---

    // move sphere down at sphere_speed * delta_t, framerate independant
    sphere.pos_y += int( float32(sphere.speed) * rl.GetFrameTime() )

    // check if sphere reached the bottom and reset its position
    if sphere.pos_y >= gs.window_height - int(sphere.radius) {

      // check if sphere was in cube
      if sphere.pos_x >= cube.pos_x &&
         sphere.pos_x <= cube.pos_x + cube.size {
        gs.current_score += 1
      }

      // reset y position back to top
      sphere.pos_y = 0
      // rng the sphere pos x
      sphere.pos_x = int( rand.Float32() * float32(gs.window_width) )

      // clamp between sphere_radius and window_width - sphere_radius
      if sphere.pos_x < int(sphere.radius){
        sphere.pos_x = int(sphere.radius)
      } else if sphere.pos_x > int(float32(gs.window_width) - sphere.radius){
        sphere.pos_x = int(float32(gs.window_width) - sphere.radius)
      }
    }

    // --- cube logic ---

    // check keyboard input
    if rl.IsKeyDown(rl.KeyRight) {
      cube.pos_x += int( float32(cube.speed) * rl.GetFrameTime() )
    }else if rl.IsKeyDown(rl.KeyLeft) {
      cube.pos_x -= int( float32(cube.speed) * rl.GetFrameTime() )
    }
    // clamp to 0 - window_width
    if cube.pos_x < 0 {
      cube.pos_x = 0
    } else if cube.pos_x > gs.window_width - cube.size {
      cube.pos_x = gs.window_width - cube.size
    }

    // --- rendering --- 

    // draw main rectangle
    rl.DrawRectangle( int32(cube.pos_x), int32(cube.pos_y), int32(cube.size), int32(cube.size), rl.Blue )
    // draw sphere
    rl.DrawCircle( int32(sphere.pos_x), int32(sphere.pos_y), sphere.radius, rl.Red )
    // draw another rectangle slightly below the main rectangle to the sphere disapears  
    rl.DrawRectangle( int32(cube.pos_x), int32(cube.pos_y + (cube.size / 2) ), int32(cube.size), int32(cube.size), rl.DarkBlue )

    rl.DrawText("catch the falling sphere", int32(gs.window_width) - 400, 50, int32(gs.font_size_main), rl.LightGray)
    
    var score_str string = fmt.Sprintf( "score: %v", gs.current_score )
    // unsigned int TextLength(const char *text);
    rl.DrawText(score_str, int32(gs.window_width / 2 - ( len(score_str) * (gs.font_size_score / 2) ) ), 25, int32(gs.font_size_score), rl.Orange)

    rl.EndDrawing()
  }
}
