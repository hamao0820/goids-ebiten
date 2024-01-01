package goids

import (
	"image"
	"image/draw"
	_ "image/png"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	xdraw "golang.org/x/image/draw"
)

const (
	Width  = 640
	Height = 480
)

var (
	gopher *ebiten.Image
)

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
}

func NewGame() *Game {
	return &Game{}
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(gopher, nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return Width, Height
}
