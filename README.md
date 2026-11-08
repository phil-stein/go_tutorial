# go language tutorial

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
