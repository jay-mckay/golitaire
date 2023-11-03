package main

import (
  "fmt"
  "os"
  
  "github.com/gdamore/tcell/v2"
  "github.com/gdamore/tcell/v2/encoding"
)


func check(e error) {
  if e != nil {
    fmt.Fprintf(os.Stderr, "%v\n", e)
    os.Exit(1)
  }
}

func main() {
  encoding.Register()

  screen, err := tcell.NewScreen()
  check(err)

  err = screen.Init()
  check(err)

  defaultStyle := tcell.StyleDefault.
  	Background(tcell.ColorBlack).
	Foreground(tcell.ColorWhite)

  screen.SetStyle(defaultStyle)
	
  width, height := screen.Size()
  screen.Clear()
  newStyle := tcell.StyleDefault.
  	Background(tcell.ColorWhite).
	Foreground(tcell.ColorCadetBlue.TrueColor())
  
  screen.SetContent(width/2, height/2, 'H', nil, newStyle)
  screen.Show()

  for {
    switch event := screen.PollEvent().(type) {
      case *tcell.EventKey:
        if event.Key() == tcell.KeyEscape {  
	  screen.Fini()
	  os.Exit(0)
      }
    }
  }
}
