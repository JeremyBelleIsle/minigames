package backToHomeButton

import (
	"MiniGames/interfaceSettings"
	"image/color"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type HomeButton struct {
	x, y, w, h                      float64
	outlineClr, inlineClr, wordsClr color.RGBA
}

func New() HomeButton {
	w, h := 250.0, 70.0

	return HomeButton{
		x:          (interfaceSettings.ScreenWidth / 2) - (w / 2),
		y:          interfaceSettings.ScreenHeight/2 + 100,
		w:          w,
		h:          h,
		outlineClr: color.RGBA{255, 16, 240, 255},
		inlineClr:  color.RGBA{0, 255, 255, 255},
		wordsClr:   color.RGBA{255, 94, 0, 255},
	}
}

func (homeButton HomeButton) Clic() bool {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false
	}
	cx, cy := ebiten.CursorPosition()

	return gameutil.ClickInARect(float64(cx), float64(cy), homeButton.x, homeButton.y, homeButton.w, homeButton.h)
}

func (homeButton HomeButton) Draw(screen *ebiten.Image, mplusSource *text.GoTextFaceSource) {
	vector.DrawFilledRect(screen, float32(homeButton.x), float32(homeButton.y), float32(homeButton.w), float32(homeButton.h), homeButton.inlineClr, false)

	vector.StrokeRect(screen, float32(homeButton.x), float32(homeButton.y), float32(homeButton.w), float32(homeButton.h), 20, homeButton.outlineClr, false)

	gameutil.DrawText("HOME", 50, int(homeButton.x)+int(homeButton.w), homeButton.x+30, homeButton.y+15, 0, screen, homeButton.wordsClr, mplusSource)
}
