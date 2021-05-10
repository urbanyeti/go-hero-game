package graphics

import (
	"fmt"
	"image"
	"io/ioutil"
	"os"

	"github.com/hajimehoshi/ebiten"
	"github.com/hajimehoshi/ebiten/ebitenutil"
	log "github.com/sirupsen/logrus"
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
	X           int
	Y           int
	vx          int
	vy          int
	Flipped     bool
}

type Animations map[string]*Animation

type Animation struct {
	Sprite
	current int
	images  []*ebiten.Image
	isDone  bool
	isLoop  bool
}

func (a *Animation) Clone(isLoop bool) *Animation {
	n := Animation{}
	n.images = a.images
	n.isLoop = isLoop
	n.Sprite = a.Sprite

	return &n
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

type Graphics struct {
	count      int
	op         ebiten.DrawImageOptions
	inited     bool
	Animations *Animations
}

func NewGraphics() *Graphics {
	g := &Graphics{}
	g.Animations = &Animations{}
	return g
}

func NewAnimation(folder string, isEnemy bool) *Animation {
	imgs, err := loadImageFolder(folder)
	if err != nil {
		log.Fatal(err)
	}
	w, h := imgs[0].Size()
	w, h = w/4, h/4
	x, y := (SCREENWIDTH-w)/3, (SCREENHEIGHT)/2
	if isEnemy {
		x += 200
	}
	return &Animation{
		images: imgs,
		Sprite: Sprite{
			imageWidth:  w,
			imageHeight: h,
			X:           x,
			Y:           y,
			Flipped:     isEnemy,
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
		if err != nil {
			log.Warn(err)
			continue
		}
		ebitenImage, _ := ebiten.NewImageFromImage(img, ebiten.FilterDefault)
		imgs = append(imgs, ebitenImage)
	}
	return imgs, nil
}

func (g *Graphics) Update(screen *ebiten.Image) error {
	g.count++

	return nil
}

func (g *Graphics) Draw(screen *ebiten.Image) {
	for _, a := range *g.Animations {
		g.op.GeoM.Reset()
		if a.Flipped {
			g.op.GeoM.Scale(-.25, .25)
			g.op.GeoM.Translate(float64(a.imageWidth), 0)
		} else {
			g.op.GeoM.Scale(.25, .25)
		}
		g.op.GeoM.Translate(float64(a.X), float64(a.Y))
		screen.DrawImage(a.Play(), &g.op)
	}

	msg := fmt.Sprintf(`TPS: %0.2f
FPS: %0.2f`, ebiten.CurrentTPS(), ebiten.CurrentFPS())
	ebitenutil.DebugPrint(screen, msg)
}

func (g *Graphics) Layout(outsideWidth, outsideHeight int) (int, int) {
	return SCREENWIDTH, SCREENHEIGHT
}

func (g *Graphics) Run() {
	ebiten.SetWindowSize(SCREENWIDTH*2, SCREENHEIGHT*2)
	ebiten.SetWindowTitle("go-hero-game")
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
