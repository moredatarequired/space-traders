package main

import (
	"github.com/moredatarequired/space-traders/lib"
	"image"
	"image/png"
	"os"
	"image/color"
	"fmt"
)

func min(x, y uint8) uint8 {
	if x < y { return x }
	return y
}

func abs(x int) int {
	if x >= 0 { return x }
	return -x
}

func main() {
	size, scale := 1024, 1
	plot := image.NewGray(image.Rect(-size, -size, size, size))
	fixed := &lib.Ship{}
	gnat := &lib.Ship{}
	gnat.Position.X = 400 //float64(size * scale)
	gnat.Velocity.Y = 0.1
	x, y := int(gnat.Position.X), int(gnat.Position.Y)
	steps := 100000
	//for abs(x) < size*scale && abs(y) < size*scale {
	for i := 0; i < steps; i++ {
		color := color.Gray{uint8(256 * float64(i) / float64(steps))}
		plot.SetGray(x/scale, y/scale, color)
		gnat.MaintainDistance(fixed, 40, 1000)
		if false {
		p, v, a := gnat.Position, gnat.Velocity, gnat.Acceleration
		fmt.Printf("At (%.3f, %.3f)->(%.3f, %.3f)=>(%.3f, %.3f)\n",
			p.X, p.Y, v.X, v.Y, a.X, a.Y)
		}
		if i % 100 == 100 {
			p, v, a := gnat.Position, gnat.Velocity, gnat.Acceleration
			fmt.Printf("At (%.3f, %.3f)->(%.3f, %.3f)=>(%.3f, %.3f)\n",
				p.X, p.Y, v.X, v.Y, a.X, a.Y)
			fmt.Println(gnat.Velocity.Length())
		}
		gnat.Move(0.01)
		x, y = int(gnat.Position.X), int(gnat.Position.Y)
	}
	if fo, err := os.Create("trajectory.png"); err == nil {
		png.Encode(fo, plot)
		if err := fo.Close(); err != nil { panic(err) }
	} else { panic(err) }
}
