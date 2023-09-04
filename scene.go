package main

import (
	"fmt"
	"strings"

	"github.com/gdamore/tcell"
)

// (입력 시, 시각과 위치, 갈 수 있는 정류장 목록 바뀜)
func (g Game) printSceneBoarding() {
	fmt.Fprintln(g.tw, "번호\t정류장\t소요 시간\t")
}

// 1. 2호선 구의역 (도보 이동 시간: 13분)
// 2. 2호선 강변역 (도보 이동 시간: 6분)
// 3. 2, 8호선 잠실역 (도보
// 4. 현대아파트앞 [광장동 현대아파트앞 방면]

// SceneLandmark
type Scene interface {
	Render()
	HandleInput(ev *tcell.EventKey) any
}

// 	SceneBusStop
// 	SceneTrainStation
// 	SceneTravel
// 	SceneWalk

// 종료 조건: 모든 목적지를 도착하였거나 24시에 도달
type SceneStart struct{}

func (s SceneStart) Render() {
	fmt.Println("아무 키나 눌러 게임을 시작하세요.")
}
func (s SceneStart) HandleInput(ev *tcell.EventKey) any {
	return SceneBusStop{}
}

type SceneBusStop struct {
	screen  tcell.Screen
	bus     *BusRoute
	page    int
	maxPage int
}

func NewSceneBusStop(screen tcell.Screen, bus *BusRoute) SceneBusStop {

	return SceneBusStop{screen: screen, bus: bus}
}

func (s SceneBusStop) Render() {
	w, h := s.screen.Size()
	fmt.Printf("현재 위치: %s\n", s.bus.Name)
	fmt.Println(strings.Repeat("=", w))
	fmt.Printf("0. 주변 정류장으로 걸어가기\n")
	for i, stop := range s.bus.BusStops[:h-5] {
		fmt.Printf("%d. %s\n", i+1, stop.Name)
	}
}

// 현재 위치: 현대아파트앞 [강변역 방면] (서울 광진구 구의3동)
// 정렬 기준: 버스 이름
// 1. 3212 (대기 시간: 0분)
// 2. 3214 (대기 시간: 5분)
// 3. 남양주1 (대기 시간: 20분)
// 4. 강동01 (대기 시간: 2분)
// 5. 구리96 (대기 시간: 25분)
func (s *SceneBusStop) HandleInput(ev *tcell.EventKey) any {
	switch {
	case ev.Rune() == '0':
		return SceneWalk{}
	case '1' <= ev.Rune() && ev.Rune() <= '9':
		return SceneTravel{}
	case ev.Rune() == 'q' || ev.Rune() == 'Q':
		s.page++

	}
	return nil
}

type SceneEnd struct{}

func (s SceneEnd) Render() {
	fmt.Println("게임을 종료합니다.")
}

func (s SceneEnd) HandleInput(ev *tcell.EventKey) any {
	return nil
}
