package game

import (
	"math/rand"

	"github.com/hamao0820/goids-ebiten/vector"
)

type ImageType int

const (
	Front ImageType = iota
	Side
	Pink
)

const GopherSize = 32

type Goid struct {
	position     vector.Vector
	velocity     vector.Vector
	acceleration vector.Vector
	maxSpeed     float64
	maxForce     float64
	sight        float64
	imageType    ImageType
}

func NewGoid(position vector.Vector, maxSpeed, maxForce float64, sight float64) Goid {
	velocity := vector.New(rand.Float64()*2-1, rand.Float64()*2-1)
	velocity.Scale(rand.Float64()*4 - rand.Float64()*2)

	var t ImageType

	r := rand.Float64()

	if r < 0.001 { // 0.1%
		t = Pink
	} else if r < 0.011 { // 1%
		t = Side
	} else {
		t = Front
	}

	return Goid{position: position, velocity: velocity, maxSpeed: float64(maxSpeed), maxForce: float64(maxForce), sight: sight, imageType: t}
}

func (g *Goid) Seek(t vector.Vector) {
	tv := vector.Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := vector.Sub(tv, g.velocity)
	g.acceleration.Add(force)
}

func (g *Goid) Flee(t vector.Vector) {
	tv := vector.Sub(t, g.position)
	tv.Limit(g.maxSpeed)
	force := vector.Sub(tv, g.velocity)
	g.acceleration.Sub(force)
}

func (g Goid) IsInsight(g2 Goid) bool {
	d := vector.Sub(g.position, g2.position).Len()
	return d < g.sight
}

func (g *Goid) Align(neighbors []Goid) {
	var avgVel vector.Vector
	n := len(neighbors)
	for _, other := range neighbors {
		avgVel.Add(other.velocity)
	}
	if n > 0 {
		avgVel.ScalarMul(1 / float64(n))
		avgVel.Limit(g.maxSpeed)
		g.acceleration.Add(vector.Sub(avgVel, g.velocity))
	}
}

func (g *Goid) Separate(neighbors []Goid) {
	for _, other := range neighbors {
		d := vector.Sub(g.position, other.position).Len()
		if d < 50 {
			g.Flee(other.position)
		}
	}
}

func (g *Goid) Cohesive(neighbors []Goid) {
	var avgPos vector.Vector
	n := len(neighbors)
	for _, other := range neighbors {
		avgPos.Add(other.position)
	}
	if n > 0 {
		avgPos.ScalarMul(1 / float64(n))
		g.Seek(avgPos)
	}
}

func (g *Goid) AvoidMouse(mouse vector.Vector) {
	if mouse.X < 0 || mouse.Y < 0 {
		return
	}
	d := vector.Sub(g.position, mouse).Len()
	if d < 100 {
		g.acceleration.ScalarMul(0)
		g.Flee(mouse)
	}
}

func (g *Goid) neighbors(goids []Goid) []Goid {
	neighbors := make([]Goid, 0)
	for _, other := range goids {
		if g == &other || !g.IsInsight(other) {
			continue
		}
		neighbors = append(neighbors, other)
	}
	return neighbors
}

func (g *Goid) Flock(goids []Goid, mouse vector.Vector) {
	neighbors := g.neighbors(goids)
	g.Align(neighbors)
	g.Separate(neighbors)
	g.Cohesive(neighbors)
	g.AvoidMouse(mouse)
}

func (g *Goid) AdjustEdge(width, height float64) {
	if g.position.X < -float64(GopherSize) {
		g.position.X = width
	} else if g.position.X > width+float64(GopherSize) {
		g.position.X = -float64(GopherSize)
	}
	if g.position.Y < -float64(GopherSize) {
		g.position.Y = height
	} else if g.position.Y > height+float64(GopherSize) {
		g.position.Y = -float64(GopherSize)
	}
}

func (g *Goid) Update(width, height float64) {
	g.acceleration.Limit(g.maxForce)
	g.velocity.Add(g.acceleration)
	g.velocity.Limit(g.maxSpeed)
	g.position.Add(g.velocity)
	g.acceleration.ScalarMul(0)

	g.AdjustEdge(width, height)
}

func (g Goid) Position() vector.Vector {
	return g.position
}

func (g Goid) ImageType() ImageType {
	return g.imageType
}
