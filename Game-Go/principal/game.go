package principal

import (
	"log"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"juego/procesos"
)

const (
	screenWidth  = 640
	screenHeight = 480
	paddleWidth  = 15
)

type Gameplay struct{}

func (g *Gameplay) Update() error {
	return nil
}

func (g *Gameplay) Draw(screen *ebiten.Image) {
	screen.DrawImage(procesos.BackgroundImage, nil)
	procesos.DrawImage(screen, procesos.PaddleImage, 0, procesos.Player)
	procesos.DrawImage(screen, procesos.BallImage, procesos.BallX, procesos.BallY)
	procesos.DrawImage(screen, procesos.Def1, screenWidth/2-paddleWidth/2, procesos.Paddle1Y)
	procesos.DrawImage(screen, procesos.Def2, (3*screenWidth)/3-paddleWidth/2-20, procesos.Paddle2Y)

}

func Imagenes(){
	
	var err error
	procesos.PaddleImage, _, err = ebitenutil.NewImageFromFile("img/portero.png")
	if err != nil {
		log.Fatal(err)
	}

	procesos.BallImage, _, err = ebitenutil.NewImageFromFile("img/balon.png")
	if err != nil {
		log.Fatal(err)
	}

	procesos.BackgroundImage, _, err = ebitenutil.NewImageFromFile("img/fondo.png")
	if err != nil {
		log.Fatal(err)
	}

	procesos.Def1, _, err = ebitenutil.NewImageFromFile("img/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

	procesos.Def2, _, err = ebitenutil.NewImageFromFile("img/defensa.png")
	if err != nil {
		log.Fatal(err)
	}

}

func (g *Gameplay) Layout(outsideWidth, outsideHeight int) (int, int) {
	return screenWidth, screenHeight
}
