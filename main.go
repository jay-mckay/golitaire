package main

import (
  "fmt"
  "os"
  "os/exec"
)

// TODO: look into tcell

var up byte = 65
var down byte = 66
var right byte = 67
var left byte = 68

func main() {
  // no input buffering
  exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
  // don't display characters entered
  exec.Command("stty", "-F", "/dev/tty", "-echo").Run()
  
  var b []byte = make([]byte, 1)
  for {
    os.Stdin.Read(b)
    if b[0] == up {
    } else if b[0] == down {
    } else if b[0] == right {
    } else if b[0] == left {
    }
  }
}
