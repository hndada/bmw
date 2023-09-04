package main

import (
	"fmt"
	"time"

	"github.com/gdamore/tcell"
)

type Game struct {
	scene  Scene
	screen tcell.Screen

	time  time.Time
	place Place
}

func NewGame() *Game {
	var err error

	g := Game{}
	g.screen, err = tcell.NewScreen()
	if err != nil {
		panic(err)
	}

	err = g.screen.Init()
	if err != nil {
		panic(err)
	}

	return &g
}

// printToCenter: screen.Size(), screen.SetContent(x, y, c, nil, tcell.StyleDefault)
func (g *Game) Render() {
	g.screen.Clear()
	fmt.Printf("시간: %d시 %d분\n", g.time.Hour(), g.time.Minute())
	fmt.Printf("장소: %s\n", g.place.Name)
	fmt.Println()
	g.scene.Render()
	g.screen.Show()
}

func (g *Game) HandleInput(ev *tcell.EventKey) any {
	return g.scene.HandleInput(ev)
}

func (g *Game) Finish() {
	g.screen.Fini()
}

func main() {
	g := NewGame()
	parseBusData()
	parseTrainData()
	for {
		g.Render()
		ev := g.screen.PollEvent().(*tcell.EventKey)
		switch g.HandleInput(ev).(type) {
		case SceneEnd:
			g.Finish()
			break
		}
	}
}
