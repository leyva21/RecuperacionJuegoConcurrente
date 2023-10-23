package main

import (
	"juego/principal"
	"juego/procesos"
	"log"
	"sync"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

var lastUpdate time.Time
var mutex sync.Mutex

const (
	ScreenHeight = 480
	ScreenWidth  = 600
)

func main() {
	lastUpdate = time.Now()
	ebiten.SetWindowSize(ScreenWidth, ScreenHeight)
	ebiten.SetWindowTitle("Porteo")

	principal.Imagenes()
	go procesos.MoverPortero()
	go procesos.MoverBalon()
	go procesos.MoverDefensas()

	game := &principal.Gameplay{}
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
