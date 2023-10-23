package procesos

import (
	"github.com/hajimehoshi/ebiten/v2"
	"time"
)

const (
	ScreenHeight = 480
	paddleHeight = 80
	ScreenWidth  = 640
	paddleWidth  = 15
)

var Player = ScreenHeight / 2
func MoverPortero() {
	for {
		if ebiten.IsKeyPressed(ebiten.KeyW) && Player > 0 {
			Player -= 5
		}
		if ebiten.IsKeyPressed(ebiten.KeyS) && Player < ScreenHeight-paddleHeight {
			Player += 5
		}
		time.Sleep(time.Millisecond * 16)
	}
}
