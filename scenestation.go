package main

import "fmt"

type SceneStation struct {
}

func (s SceneStation) Print(g *Game) {
	f("현재 위치: %s %s\n", g.location.Name, g.location.Direction)
}

func (s SceneStation) HandleInput() {
	fmt.Scanf()
}
