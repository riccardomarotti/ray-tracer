package main

import (
	"fmt"
	"math"
)

func main() {
	camera := Camera{1000, 500, math.Pi / 3, ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))}
	sphereMaterial := DefaultMaterial()
	sphereMaterial.color = Color{0.2, 1, 1}
	sphere := MakeSphere(Identity().Translate(-0.5, 1, 0.5), sphereMaterial)

	sphere2Material := DefaultMaterial()
	sphere2Material.color = Color{0.5, 1, 0.1}
	sphere2Material.diffuse = 0.7
	sphere2Material.specular = 0.3
	sphere2 := MakeSphere(Identity().Translate(1.5, 0.5, -0.5).Scale(0.5, 0.5, 0.5), sphere2Material)

	sphere3Material := DefaultMaterial()
	sphere3Material.color = Color{1, 0.8, 0.1}
	sphere3Material.diffuse = 0.7
	sphere3Material.specular = 0.3
	sphere3 := MakeSphere(Identity().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33), sphere3Material)

	floorMaterial := DefaultMaterial()
	floorMaterial.color = Color{1, 0.9, 0.9}
	floorMaterial.specular = 0
	floor := MakeSphere(Identity().Scale(10, 0.01, 10), floorMaterial)

	leftWall := MakeSphere(Identity().Translate(0, 0, 5).RotateY(-math.Pi/4).RotateX(math.Pi/2).Scale(10, 0.01, 10), floorMaterial)
	rightWall := MakeSphere(Identity().Translate(0, 0, 5).RotateY(math.Pi/4).RotateX(math.Pi/2).Scale(10, 0.01, 10), floorMaterial)

	lightPosition := Point(-10, 10, -10)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	world := World{light, []Object{sphere, sphere2, sphere3, floor, leftWall, rightWall}}
	c := camera.Render(world)
	fmt.Printf(c.PPM())
}
