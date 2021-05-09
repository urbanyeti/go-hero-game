package main

import (
	"fmt"
	"image"
	"io/ioutil"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
)

const (
	SCREENWIDTH      = 500
	SCREENHEIGHT     = 370
	HOLD             = 6
	ATTACKPATH       = `Knight_02\03-Attack\`
	WALKPATH         = `Knight_02\02-Walk\`
	HERO_IDLE_PATH   = `Knight_02\01-Idle\`
	GOBLIN_IDLE_PATH = `Goblin_02\01-Idle\`
	GOBLIN_DIE_PATH  = `Goblin_02\07-Die\`
)

type Sprite struct {
	imageWidth  int
	imageHeight int
	x           int
	y           int
	vx          int
	vy          int
	hold        int
	flipped     bool
}

type Animation struct {
	Sprite
	current int
	images  []*ebiten.Image
	isDone  bool
	isLoop  bool
}

func (a *Animation) Play() *ebiten.Image {
	if a.isDone {
		return a.images[len(a.images)-1]
	}

	a.current++
	if (a.current / HOLD) >= len(a.images) {
		if a.isLoop {
			a.current = 0
		} else {
			a.isDone = true
			return a.images[len(a.images)-1]
		}
	}
	return a.images[(a.current / HOLD)]
}

func (s *Sprite) Update() {

}

type Game struct {
	count  int
	op     ebiten.DrawImageOptions
	inited bool
	a      *Animation
	a2     *Animation
	g      *Animation
	g2     *Animation
}

func (g *Game) LoadContent() {
	g.a = g.NewAnimation(ATTACKPATH, true)
	g.a2 = g.NewAnimation(HERO_IDLE_PATH, true)
	g.g = g.NewAnimation(GOBLIN_IDLE_PATH, true)
	g.g.x += 150
	g.g.flipped = true
	g.g2 = g.NewAnimation(GOBLIN_DIE_PATH, false)
	g.g2.x += 150
	g.g2.flipped = true
}

func (g *Game) NewAnimation(folder string, isLoop bool) *Animation {
	imgs, err := loadImageFolder(folder)
	if err != nil {
		log.Fatal(err)
	}
	w, h := imgs[0].Size()
	w, h = w/4, h/4
	x, y := (SCREENWIDTH-w)/3, (SCREENHEIGHT)/2
	return &Animation{
		images: imgs,
		isLoop: isLoop,
		Sprite: Sprite{
			imageWidth:  w,
			imageHeight: h,
			x:           x,
			y:           y,
		},
	}

}

func loadImageFolder(folderName string) ([]*ebiten.Image, error) {
	var imgs []*ebiten.Image
	files, err := ioutil.ReadDir(folderName)
	if err != nil {
		return nil, err
	}
	for _, f := range files {
		f, err := os.Open(folderName + f.Name())
		if err != nil {
			return nil, err
		}
		img, _, err := image.Decode(f)
		ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		imgs = append(imgs, ebitenImage)
	}
	return imgs, nil
}

func (g *Game) Update(screen *ebiten.Image) error {
	g.count++
	if g.count == 100 {
		g.g = g.g2
		g.a = g.a2
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.op.GeoM.Reset()
	g.op.GeoM.Scale(.25, .25)
	g.op.GeoM.Translate(float64(g.a.x), float64(g.a.y))
	screen.DrawImage(g.a.Play(), &g.op)

	g.op.GeoM.Reset()
	g.op.GeoM.Scale(-.25, .25)
	g.op.GeoM.Translate(float64(g.g.imageWidth), 0)
	g.op.GeoM.Translate(float64(g.g.x), float64(g.g.y))
	screen.DrawImage(g.g.Play(), &g.op)

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f
X: %v Y: %v`, ebiten.CurrentTPS(), ebiten.CurrentFPS(), g.a.x, g.a.y)
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func main() {
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	ebiten.SetWindowTitle("Game")
	game := Game{}
	game.LoadContent()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
