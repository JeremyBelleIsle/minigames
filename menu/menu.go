package menu

import (
	"MiniGames/interfaceSettings"
	"MiniGames/tic_tac_toe"
	"image/color"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game int

const (
	Menu Game = iota
	TicTacToe
	Snake
	RockPaperScissors
	Pong
	FlappyBird
)

type miniGame struct {
	name       Game
	x, y, w, h float64
	clr        color.RGBA
	img        *ebiten.Image
}

type line struct {
	x1, y1, x2, y2 float64
	thickness      float64
	clr            color.RGBA
}

var CurrentGame Game
var MiniGames = []miniGame{}
var lines []line

func (g Game) String() string {
	switch g {
	case TicTacToe:
		return "tic tac toe"
	case RockPaperScissors:
		return "rock paper scissors"
	case Snake:
		return "snake"
	case Pong:
		return "pong"
	case FlappyBird:
		return "flappy bird"
	default:
		return "unknown"
	}
}

func Init() {
	names := []Game{
		TicTacToe,
		Snake,
		RockPaperScissors,
		Pong,
		FlappyBird,
	}
	x, y := 10.0, 5.0

	for i := 0; i < 5; i++ {
		MiniGames = append(MiniGames, miniGame{
			name: names[i],
			x:    x,
			y:    y,
			w:    interfaceSettings.ScreenWidth/3 - 30,
			h:    interfaceSettings.ScreenHeight/3 - 30,
			clr:  color.RGBA{0, 255, 255, 255},
		})

		x += interfaceSettings.ScreenWidth / 3
		if x >= interfaceSettings.ScreenWidth {
			x = 10
			y += interfaceSettings.ScreenHeight / 3
		}
	}

	neon := color.RGBA{57, 255, 20, 255}
	w, h := float64(interfaceSettings.ScreenWidth), float64(interfaceSettings.ScreenHeight)

	for i := 1; i <= 2; i++ {
		fi := float64(i)
		// verticales
		lines = append(lines, line{x1: w / 3 * fi, y1: 0, x2: w / 3 * fi, y2: h, thickness: 4, clr: neon})
		// horizontales
		lines = append(lines, line{x1: 0, y1: h / 3 * fi, x2: w, y2: h / 3 * fi, thickness: 4, clr: neon})
	}
}

func ClickOnMiniGame() {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return
	}

	cx, cy := ebiten.CursorPosition()
	for _, MiniGame := range MiniGames {
		if gameutil.ClickInARect(float64(cx), float64(cy), MiniGame.x, MiniGame.y, MiniGame.w, MiniGame.h) {
			CurrentGame = MiniGame.name

			switch CurrentGame {
			case TicTacToe:
				tic_tac_toe.Init()
			}
		}
	}
}

func Draw(screen *ebiten.Image, mplusSource *text.GoTextFaceSource) {
	for _, miniGame := range MiniGames {
		gameutil.DrawText(miniGame.name.String(), 10, int(miniGame.x+(interfaceSettings.ScreenWidth/3)-20), miniGame.x, miniGame.y, 0, screen, miniGame.clr, mplusSource)
	}

	for _, l := range lines {
		vector.StrokeLine(screen, float32(l.x1), float32(l.y1), float32(l.x2), float32(l.y2), float32(l.thickness), l.clr, false)
	}
}
