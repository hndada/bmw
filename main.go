package main

import (
	"fmt"
	"text/tabwriter"
	"time"
)

// 현재 있는 데이터
// 버스 노선 별 정류장
// 정류장 이름 및 위치
// Todo: 정류장 방향은 다음 정류장 따오기.
var f = fmt.Printf

func init() {
	// json 파일에서 불러오기
}

// Bus:   0->(1->4->1)->5
// Train: 0->(1->2->3->4->1)->5
const (
	SceneModeIntro = iota
	SceneModeStation
	SceneModeNeighbor
	SceneModeTrainDirection
	SceneModeBoarding
	SceneModeFinish
)

type Place = Station

// 게임 목표: 10군데를 최대한 빨리 찍고 오기
type Game struct {
	tw        *tabwriter.Writer
	sceneMode int
	location  Place
	time      time.Time
	dsts      []Place
	visiteds  []bool
}

type Scene interface {
	Update() any
	Draw()
}

func main() {
	g := newGame()
	for {
		g.printState()
		g.printScene()
		g.printDestinations()
		g.readInput()
	}
}

func (g Game) printScene() {
	switch g.sceneMode {
	case SceneModeIntro:
		g.printSceneIntro()
	case SceneModeStation:
		g.printSceneStation()
		// 추가
	}
}

// table을 그리고 싶긴 한데.
// 오른쪽에 그려짐
// 해당 목적지까지 도보로 이동하시겠습니까?
func (g Game) printDestinations() {
}

// 현재 시각: 20:28
func (g Game) printState() {
	f("현재 시각 %02d:%02d\n", g.time.Hour(), g.time.Minute())
	l := g.location
	f("현재 위치: %s [%s] (%s)\n", l.Name, l.Direction, l.Address)
}

func newGame() *Game {
	const difficulty = 10
	// 오전 6시로 초기화
	t := time.Now()

	// 랜드마크 맵에서 10개 임의 선택
	dsts := make([]Place, 0, difficulty)
	visiteds := make([]bool, len(dsts))

	// 다음의 장소 중 하나로 랜덤 시작
	// 1. 국립과학수사연구소 (서부본부)
	// 2. 동서울터미널 (동부본부)
	// 3. 한국과학기술연구원 (북부본부)
	// startLocation :=
	return &Game{
		sceneMode: SceneModeIntro,
		location:  startLocation,
		time:      t,
		dsts:      dsts,
		visiteds:  visiteds,
	}
}

func (g Game) printSceneIntro() {
	for _, str := range []string{"아 집에 가고 싶다", ".", ".", "."} {
		fmt.Print(str)
		time.Sleep(600 * time.Millisecond)
	}
	fmt.Println()
}

// (버스 노선 목록 창; 왼쪽)

// 현재 location의 속성이 BusStop일 경우, 현재 location에 소속된 노선을 깐다
func (g Game) printSceneStation() {

}

// 네이버 API 기준, 버스정류장에 전철역 있으나 활용 어려울 듯
func (g Game) printSceneNeighbor() {
	fmt.Fprintln(g.tw, "번호\t정류장 및 역\t이동 시간\t")
	// map에서 iterate: 반경 800m 내 정류장 및 전철역 표시
	// 거리 비례하여 시간 소요

}

func (g Game) printSceneTrainDirection() {
	// (전철 선택 시)
	// 호선, 방향 선택
}

// (입력 시, 시각과 위치, 갈 수 있는 정류장 목록 바뀜)
func (g Game) printSceneBoarding() {
	fmt.Fprintln(g.tw, "번호\t정류장\t소요 시간\t")
}

func (g Game) printSceneFinish() {
}
