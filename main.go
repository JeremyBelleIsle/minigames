package main

import (
	"MiniGames/interfaceSettings"
	"MiniGames/menu"
	"MiniGames/tic_tac_toe"
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct{}

var mplusSource *text.GoTextFaceSource

func (g *Game) Update() error {
	switch menu.CurrentGame {
	case menu.Menu:
		menu.ClickOnMiniGame()
	case menu.TicTacToe:
		tic_tac_toe.Run()
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	if menu.CurrentGame == menu.Menu {
		menu.Draw(screen, mplusSource)
	} else {
		switch menu.CurrentGame {
		case menu.TicTacToe:
			tic_tac_toe.Draw(screen, mplusSource)
		}
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return interfaceSettings.ScreenWidth, interfaceSettings.ScreenHeight
}

func main() {
	ebiten.SetWindowSize(interfaceSettings.ScreenWidth, interfaceSettings.ScreenHeight)
	ebiten.SetWindowTitle("")

	s, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatal(err)
	}

	mplusSource = s

	menu.Init()
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
