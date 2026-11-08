# go language tutorial
made using this [youtube video](https://www.youtube.com/watch?v=8uiZC0l4Ajw)
@NOTE: currently at 26:40

## installation
go to [go-website](https://go.dev/dl/)
<img src="assets/installation_01.png" alt="installation_01" width="600">


## start a project

  1. open terminal 
  2. `cd` to directory you wanna start a project in 
  3. `go mod init <name>` pick name as whatever you want, your application will later be called this
      1. alternatively you can `mkdir` to make directories directly in the terminal
  4. make a new file call it `main.go`
  ```go
  package main

  import "fmt"

  func main(){
    fmt.Println( "Hello World :)" )
  }
  ```
  5. make a new file call it `build.bat`
  ```batch
  @echo off

  go build .

  <name>.exe
  ```
  6. in terminal call `build` to run program

## todo
  - [ ] add simple read write terminal input in place 04
    - answer questions get points
    - enter name and age
  - [ ] 06_game_01, make simple game using raylib
    - falling boxes, move box at bottom to collect -> get highscore
    - simple platformer with no objective
    - answer given questions get points -> get highscore
  - [ ] strings, runes, bytes
  - [ ] structs and interfaces
  - [ ] pointers
  - [ ] goroutines
  - [ ] channels
  - [ ] generics

