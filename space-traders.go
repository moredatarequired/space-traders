package main

import (
	"github.com/moredatarequired/space-traders/lib"
	"image"
	"image/png"
	"os"
	"syscall"
	"image/color"
	"fmt"
)

func max(x, y uint8) uint8 {
	if x > y { return x }
	return y
}

func abs(x int) int {
	if x >= 0 { return x }
	return -x
}

func main() {
	size, scale := 1024, 100
	plot := image.NewGray(image.Rect(-size, -size, size, size))
	fixed := &lib.Ship{}
	gnat := &lib.Ship{}
	gnat.Position.X = 1
	gnat.Velocity.Y = 1
	color := color.Gray{128}
	x, y := int(gnat.Position.X), int(gnat.Position.Y)
	for abs(x) < size*scale && abs(y) < size*scale {
	//for i := 0; i < 10; i++ {
		plot.SetGray(x/scale, y/scale, color)
		color.Y = max(128, color.Y + 1)
		gnat.CircleTarget(fixed, 1)
		gnat.Move(0.1)
		x, y = int(gnat.Position.X), int(gnat.Position.Y)
	}
	png.Encode(os.NewFile(uintptr(syscall.Stdout), "/dev/stdout"), plot)
}
