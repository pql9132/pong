package main

import (
	"github.com/veandco/go-sdl2/sdl"
)

const paddleSensitivity int32 = 25

func movePaddle(paddle PongItem, direction string) {
	activePaddle := paddle.(*Paddle)

	switch direction {
	case "up":
		if activePaddle.y-paddleSensitivity-(paddleHeight/2) >= 0 {
			activePaddle.y -= paddleSensitivity
		}
	case "down":
		if activePaddle.y+paddleSensitivity+(paddleHeight/2) <= Height {
			activePaddle.y += paddleSensitivity
		}
	}
}

func handleKeyboardEvent(event *sdl.KeyboardEvent) {
	if event.Type == sdl.KEYDOWN {
        gameItemMutex.Lock()
    	defer gameItemMutex.Unlock()
        
		switch event.Keysym.Sym {
		case sdl.K_w:
			movePaddle(gameItems[0], "up")
		case sdl.K_s:
			movePaddle(gameItems[0], "down")
		case sdl.K_UP:
			movePaddle(gameItems[1], "up")
		case sdl.K_DOWN:
			movePaddle(gameItems[1], "down")
		}
	}
}
