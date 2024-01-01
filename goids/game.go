package goids

import (
	"image"
	"image/draw"
	_ "image/png"
	"math/rand"
	"os"

	"github.com/fstanis/screenresolution"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hamao0820/goids-ebiten/vector"
	xdraw "golang.org/x/image/draw"
)

const (
	goidsNum = 100
)

var (
	gopher *ebiten.Image
	Width  = 640
	Height = 480
)

func init() {
	res := screenresolution.GetPrimary()
	if res.Width == 0 || res.Height == 0 {
		return
	}
	Width = res.Width
	Height = res.Height
}

func init() {
	img, err := loadImage("assets/images/gopher.png")
	if err != nil {
		panic(err)
	}
	gopher = ebiten.NewImageFromImage(resizeByHeight(img, GopherSize))
}

func loadImage(path string) (image.Image, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	img, _, err := image.Decode(f)
	if err != nil {
		return nil, err
	}

	return img, nil
}

func resizeByHeight(img image.Image, height float64) image.Image {
	imgDst := image.NewRGBA(image.Rect(0, 0, int(float64(img.Bounds().Dx())*height/float64(img.Bounds().Dy())), int(height))) // heightを基準にリサイズ
	xdraw.CatmullRom.Scale(imgDst, imgDst.Bounds(), img, img.Bounds(), draw.Over, nil)
	return imgDst.SubImage(imgDst.Rect)
}

type Game struct {
	goids []Goid
}

func NewGame() *Game {
	goids := make([]Goid, 0, goidsNum)
	for i := 0; i < goidsNum; i++ {
		goids = append(goids, NewGoid(vector.CreateVector(float64(rand.Intn(Width)), float64(rand.Intn(Height))), 4, 0.1, 100))
	}
	return &Game{goids: goids}
}

func (g *Game) Update() error {
	x, y := ebiten.CursorPosition()
	mouse := vector.CreateVector(float64(x), float64(y))
	for i := 0; i < len(g.goids); i++ {
		goid := &g.goids[i]
		goid.Flock(g.goids, mouse)
		goid.Update(float64(Width), float64(Height))
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for _, goid := range g.goids {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate(-float64(GopherSize)/2, -float64(GopherSize)/2)
		op.GeoM.Translate(goid.position.X, goid.position.Y)
		screen.DrawImage(gopher, op)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}
