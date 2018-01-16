package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"sync"
)

var (
	gameItemMutex = &sync.Mutex{}
	gameItems     = make([]PongItem, 3)
)

func initializeGameItems() {
	gameItemMutex.Lock()
	defer gameItemMutex.Unlock()
	gameItems[0] = &Paddle{
		side: "left",
		y:    Height / 2,
	}
	gameItems[1] = &Paddle{
		side: "right",
		y:    Height / 2,
	}
	gameItems[2] = &Ball{
		x:  Width / 2,
		y:  Height / 2,
		vx: 3,
		vy: 1,
	}
}

func renderGameScene(renderer *sdl.Renderer, scene []PongItem) error {
	renderer.Clear()
	renderer.SetDrawColor(225, 225, 225, 255)
	for _, item := range scene {
		if err := item.Draw(renderer); err != nil {
			return err
		}
	}
	renderer.SetDrawColor(0, 0, 0, 255) //Makes window background black
	renderer.Present()
	return nil
}

func moveBall(b *Ball) {
	b.x += b.vx
	b.y += b.vy
	if b.y < 0 || b.y > Height {
		b.vy *= -1
	}
}

func paddleContact(b *Ball) bool {
	leftPosY := gameItems[0].(*Paddle).y
	rightPosY := gameItems[1].(*Paddle).y

	return !(b.x-(ballSize/2) >= paddleWidth && b.x+(ballSize/2) <= Width-paddleWidth) &&
		((b.y-(ballSize/2) >= leftPosY-(paddleHeight/2) && b.y+(ballSize/2) <= leftPosY+(paddleHeight/2)) ||
			(b.y-(ballSize/2) >= rightPosY-(paddleHeight/2) && b.y+(ballSize/2) <= rightPosY+(paddleHeight/2)))
}

func ballOutOfBounds(b *Ball) bool {
	if b.x > Width || b.x < 0 {
		return true
	}
	return false
}

func gameloop(renderer *sdl.Renderer) error {
	for runLoop := true; runLoop; {
		gameItemMutex.Lock()
		if err := renderGameScene(renderer, gameItems); err != nil {
			return err
		}
		ball := gameItems[2].(*Ball)
		moveBall(ball)
		if paddleContact(ball) {
			ball.vx *= -1
		}
		if ballOutOfBounds(ball) {
			runLoop = false
			fmt.Println("Game Over!")
		}
		gameItemMutex.Unlock()
		sdl.Delay(12)
	}
	return nil
}

func runGame(renderer *sdl.Renderer, errorChannel chan error) {
	initializeGameItems()
	if err := gameloop(renderer); err != nil {
		errorChannel <- err
	}
}
