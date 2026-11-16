package tic_tac_toe

import (
	"MiniGames/interfaceSettings"
	"fmt"
	"image/color"

	"github.com/JeremyBelleIsle/gameutil"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

const (
	Xturn = iota
	Oturn
)

type state int

const (
	X state = iota
	O
	tie
	unknown
)

func (s state) String() string {
	switch s {
	case X:
		return "X"
	case O:
		return "O"
	case tie:
		return "tie"
	}
	return "unknown"
}

type line struct {
	x1, y1, x2, y2 float64
	thickness      float64
	clr            color.RGBA
}

type Case struct {
	x, y, w, h float64
	statut     state
	clr        color.RGBA
}

var lines = []line{}
var cases = make([]Case, 9)
var win = unknown
var winCases = []int{}
var turn = Xturn

func Init() {
	neon := color.RGBA{57, 255, 20, 255}
	thickness := 30.0

	cellW := float64(interfaceSettings.ScreenWidth) / 3
	cellH := float64(interfaceSettings.ScreenHeight) / 3

	// Initialisation des 9 cases avec leur position et taille
	for i := 0; i < 9; i++ {
		col := float64(i % 3)
		row := float64(i / 3)

		cases[i] = Case{
			x:      col * cellW,
			y:      row * cellH,
			w:      cellW,
			h:      cellH,
			statut: unknown,
		}
	}

	for i := 1; i <= 2; i++ {
		fi := float64(i)

		// verticale
		x := float64(interfaceSettings.ScreenWidth) / 3 * fi
		lines = append(lines, line{
			x1:        x,
			y1:        0,
			x2:        x,
			y2:        interfaceSettings.ScreenHeight,
			thickness: thickness,
			clr:       neon,
		})

		// horizontales
		y := float64(interfaceSettings.ScreenHeight) / 3 * fi
		lines = append(lines, line{
			x1:        0,
			y1:        y,
			x2:        interfaceSettings.ScreenWidth,
			y2:        y,
			thickness: thickness,
			clr:       neon,
		})
	}
}

func DetectClickInCase() (bool, int) {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return false, -1
	}

	clicX, clicY := ebiten.CursorPosition()
	float64clicX, float64clicY := float64(clicX), float64(clicY)

	for i, c := range cases {
		if float64clicX >= c.x && float64clicX < c.x+c.w && float64clicY >= c.y && float64clicY < c.y+c.h {
			return true, i
		}
	}

	return false, -1
}

func DetectWinOrTieGame() (state, []int) {
	for i := 0; i < 3; i++ {
		if cases[i*3].statut != unknown && cases[i*3].statut == cases[i*3+1].statut && cases[i*3+1].statut == cases[i*3+2].statut {
			return cases[i*3].statut, []int{i * 3, i*3 + 1, i*3 + 2}
		}

		// Détection verticale
		if cases[i].statut != unknown && cases[i].statut == cases[i+3].statut && cases[i+3].statut == cases[i+6].statut {
			return cases[i].statut, []int{i, i + 3, i + 6}
		}

	}

	// Détection diagonal
	if cases[0].statut != unknown && cases[0].statut == cases[4].statut && cases[4].statut == cases[8].statut {
		return cases[0].statut, []int{0, 4, 8}
	} else if cases[2].statut != unknown && cases[2].statut == cases[4].statut && cases[4].statut == cases[6].statut {
		return cases[2].statut, []int{2, 4, 6}
	}

	// Détection d'égalité
	for _, c := range cases {
		if c.statut == unknown {
			return unknown, []int{}
		}
	}

	return tie, []int{}
}

func Run() {
	ok, index := DetectClickInCase()

	// On vérifie qu'il y a eu un clic ET que la case est vide
	if ok && cases[index].statut == unknown {
		if turn == Xturn {
			cases[index].statut = X
			cases[index].clr = color.RGBA{255, 16, 240, 255}
			fmt.Printf("Case %d is now a X", index)
			turn = Oturn
			print("O turn")
		} else {
			cases[index].statut = O
			cases[index].clr = color.RGBA{0, 255, 255, 255}
			fmt.Printf("Case %d is now a O", index)
			turn = Xturn
			print("X turn")
		}
	}

	win, winCases = DetectWinOrTieGame()
}

func Draw(screen *ebiten.Image, mplusSource *text.GoTextFaceSource) {
	for _, l := range lines {
		vector.StrokeLine(screen, float32(l.x1), float32(l.y1), float32(l.x2), float32(l.y2), float32(l.thickness), l.clr, false)
	}

	for _, c := range cases {
		if c.statut == unknown {
			continue
		}

		gameutil.DrawText(c.statut.String(), 100, int(c.x+93), c.x+63, c.y+35, 0, screen, c.clr, mplusSource)
	}

	type winLineStruct struct {
		x1, y1, x2, y2 float64
		thickness      float64
		clr            color.RGBA
	}

	winLine := winLineStruct{
		thickness: 25,
		clr:       color.RGBA{31, 81, 255, 255},
	}

	if win != unknown {
		vector.DrawFilledRect(screen, 0, 0, interfaceSettings.ScreenWidth, interfaceSettings.ScreenHeight, color.RGBA{0, 0, 0, 128}, false)

		for i, winCase := range winCases {
			for j, normalCase := range cases {
				if winCase == j {
					if i == 0 {
						winLine.x1 = normalCase.x + normalCase.w/2
						winLine.y1 = normalCase.y + normalCase.h/2
					}

					if i == 2 {
						winLine.x2 = normalCase.x + normalCase.w/2
						winLine.y2 = normalCase.y + normalCase.h/2
					}
				}
			}
		}

		vector.StrokeLine(screen, float32(winLine.x1), float32(winLine.y1), float32(winLine.x2), float32(winLine.y2), float32(winLine.thickness), winLine.clr, false)

		txt := "Tie!"
		txtClr := color.RGBA{255, 255, 255, 255}
		if win != tie {
			txt = fmt.Sprintf("Player %s win!", win)
			txtClr = cases[winCases[0]].clr
		}

		gameutil.DrawText(txt, 40, interfaceSettings.ScreenWidth, float64(interfaceSettings.ScreenWidth/2-((len(txt)*40)/2)), interfaceSettings.ScreenHeight/2-100, 0, screen, txtClr, mplusSource)
	}
}
