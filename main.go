package main

import (
	"fmt"
	"math"
)

func main() {
	camera := Camera{500, 250, math.Pi / 3, ViewTransform(Point(0, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))}
	sphereMaterial := DefaultMaterial()
	sphereMaterial.color = Color{0.2, 1, 1}
	pattern1 := MakePerturbPattern(MakeRingPattern(Color{0.2, 1, 1}, Color{.8, 1, .8}, Identity().Scale(.05, .05, .05).RotateX(-math.Pi/3)))
	pattern2 := MakeGradientPattern(Color{1, .5, .3}, Color{.3, .5, 1}, Identity().RotateZ(-math.Pi/3))
	sphereMaterial.pattern = MakeBlendPattern(pattern1, pattern2)

	sphere := MakeSphere(Identity().Translate(-0.5, 1, 0.5), sphereMaterial)

	sphere2Material := DefaultMaterial()
	sphere2Material.color = Color{0.5, 1, 0.1}
	sphere2Material.diffuse = 0.7
	sphere2Material.specular = 0.3
	sphere2Material.reflective = 0.2

	gradientPattern := MakeGradientPattern(Color{.3, .5, 1}, Color{.5, 1, .3}, Identity())
	noisePattern := MakePerturbPattern(gradientPattern)
	sphere2Material.pattern = noisePattern
	sphere2 := MakeSphere(Identity().Translate(1.5, 0.7, -0.5).Scale(0.7, 0.7, 0.7), sphere2Material)

	sphere3Material := DefaultMaterial()

	sphere3Material.color = Color{1, 0.8, 0.1}
	sphere3Material.diffuse = 0.7
	sphere3Material.specular = 0.3
	sphere3Material.pattern = MakePerturbPattern(MakeStripePattern(Color{.7, .9, .5}, Color{.7, .7, .3}, Identity().Scale(.1, .1, .1).RotateZ(math.Pi/2).RotateY(math.Pi/4)))
	sphere3 := MakeSphere(Identity().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33), sphere3Material)

	floorMaterial := DefaultMaterial()
	floorMaterial.color = Color{1, 0.9, 0.9}
	floorMaterial.specular = 0
	floorMaterial.pattern = MakeCheckersPattern(Color{1, 1, 1}, Color{0, 0, 0}, Identity())
	floor := MakePlane(Identity(), floorMaterial)

	lightPosition := Point(-10, 10, -10)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	world := World{light, []Object{sphere, sphere2, sphere3, floor}}
	c := camera.Render(world)
	fmt.Printf(c.PPM())
}
