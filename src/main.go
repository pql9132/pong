package main

import (
	"github.com/veandco/go-sdl2/sdl"
	"log"
	"sync"
)

const (
	Width  int32 = 500
	Height int32 = 500
)

func checkEvent(event sdl.Event) (windowRunning bool) {
	switch e := event.(type) {
	case *sdl.QuitEvent:
		return false
	case *sdl.KeyboardEvent:
		handleKeyboardEvent(e)
	}
	return true
}

func run() error {
	//Initialize program
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		return err
	}
	defer sdl.Quit()
	window, err := sdl.CreateWindow("Pong", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED, Width, Height, sdl.WINDOW_SHOWN)
	if err != nil {
		return err
	}
	defer window.Destroy()
	renderer, err := sdl.CreateRenderer(window, -1, 0)
	if err != nil {
		return err
	}
	defer renderer.Destroy()

	//Run program
	running := true
	runMutex := &sync.Mutex{}
	errorChannel := make(chan error)

	go runGame(renderer, errorChannel)

	//Goroutine to check for errors from runGame()
	go func(running *bool) {
		err := <-errorChannel
		if err != nil {
			runMutex.Lock()
			log.Println(err)
			*running = false
			runMutex.Unlock()
		}
	}(&running)

	//Window loop
	for {
		runMutex.Lock()
		if !running {
			break
		}
		running = checkEvent(sdl.PollEvent())
		runMutex.Unlock()
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		log.Fatal(err)
	}
}
