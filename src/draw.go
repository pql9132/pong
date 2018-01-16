package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
)

const (
	paddleWidth  int32 = 10
	paddleHeight int32 = 100
	ballSize     int32 = 10
)

type PongItem interface {
	Draw(r *sdl.Renderer) error
}

type Paddle struct {
	side string
	y    int32
}

type Ball struct {
	x  int32
	y  int32
	vx int32
	vy int32
}

func (p *Paddle) Draw(r *sdl.Renderer) error {
	var paddleY int32 = p.y - (paddleHeight / 2)
	var paddleX int32
	switch p.side {
	case "left":
		paddleX = 0
	case "right":
		paddleX = Width - paddleWidth
	default:
		return fmt.Errorf("Error: cannot draw paddle, invalid side \"%v\"", p.side)
	}
	if err := r.FillRect(&sdl.Rect{X: paddleX, Y: paddleY, W: paddleWidth, H: paddleHeight}); err != nil {
		return err
	}
	return nil
}

func (b *Ball) Draw(r *sdl.Renderer) error {
	var ballX int32 = b.x - (ballSize / 2)
	var ballY int32 = b.y - (ballSize / 2)
	if err := r.FillRect(&sdl.Rect{X: ballX, Y: ballY, W: ballSize, H: ballSize}); err != nil {
		return err
	}
	return nil
}
