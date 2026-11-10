package main

import "fmt"
import "math/rand"
import rl "github.com/gen2brain/raylib-go/raylib"

func main() {
  // --- game state ---
  var   current_score   int = 0
  const font_size_main  int = 30
  const font_size_score int = 60
  const window_width    int = 1200
  const window_height   int = 800

  rl.InitWindow(int32(window_width), int32(window_height), "falling cubes game window title")
  // defer puts function call at the end of a function
  defer rl.CloseWindow()

  rl.SetTargetFPS(60)

  // --- sphere ---
  const sphere_speed  int = 400
  const sphere_radius float32 = 30.0
  var   sphere_pos_x  int = int(sphere_radius) 
  var   sphere_pos_y  int = window_height

  // --- cube ---
  const cube_size  int = 150
  const cube_speed int = 400
  const cube_pos_y int = window_height - cube_size
  var   cube_pos_x int = window_width / 2 - ( cube_size / 2 )

  for !rl.WindowShouldClose() {
    rl.BeginDrawing()

    rl.ClearBackground(rl.RayWhite)

    // --- sphere logic ---

    // move sphere down at sphere_speed * delta_t, framerate independant
    sphere_pos_y += int( float32(sphere_speed) * rl.GetFrameTime() )

    // check if sphere reached the bottom and reset its position
    if sphere_pos_y >= window_height - int(sphere_radius) {

      // check if sphere was in cube
      if sphere_pos_x >= cube_pos_x &&
         sphere_pos_x <= cube_pos_x + cube_size {
        current_score += 1
      }

      // reset y position back to top
      sphere_pos_y = 0
      // rng the sphere pos x
      sphere_pos_x = int( rand.Float32() * float32(window_width) )

      // clamp between sphere_radius and window_width - sphere_radius
      if sphere_pos_x < int(sphere_radius){
        sphere_pos_x = int(sphere_radius)
      } else if sphere_pos_x > int(float32(window_width) - sphere_radius){
        sphere_pos_x = int(float32(window_width) - sphere_radius)
      }
    }

    // --- cube logic ---

    // check keyboard input
    if rl.IsKeyDown(rl.KeyRight) {
      cube_pos_x += int( float32(cube_speed) * rl.GetFrameTime() )
    }else if rl.IsKeyDown(rl.KeyLeft) {
      cube_pos_x -= int( float32(cube_speed) * rl.GetFrameTime() )
    }
    // clamp to 0 - window_width
    if cube_pos_x < 0 {
      cube_pos_x = 0
    } else if cube_pos_x > window_width - cube_size {
      cube_pos_x = window_width - cube_size
    }

    // --- rendering --- 

    // draw main rectangle
    rl.DrawRectangle( int32(cube_pos_x), int32(cube_pos_y), int32(cube_size), int32(cube_size), rl.Blue )
    // draw sphere
    rl.DrawCircle( int32(sphere_pos_x), int32(sphere_pos_y), sphere_radius, rl.Red )
    // draw another rectangle slightly below the main rectangle to the sphere disapears  
    rl.DrawRectangle( int32(cube_pos_x), int32(cube_pos_y + (cube_size / 2) ), int32(cube_size), int32(cube_size), rl.DarkBlue )

    rl.DrawText("catch the falling sphere", int32(window_width) - 400, 50, int32(font_size_main), rl.LightGray)
    
    var score_str string = fmt.Sprintf( "score: %v", current_score )
    // unsigned int TextLength(const char *text);
    rl.DrawText(score_str, int32(window_width / 2 - ( len(score_str) * (font_size_score / 2) ) ), 25, int32(font_size_score), rl.Orange)

    rl.EndDrawing()
  }
}
