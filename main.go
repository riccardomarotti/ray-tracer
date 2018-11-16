package main

import (
	"fmt"
	"math"
)

func main() {
	camera := Camera{500, 250, math.Pi / 2, ViewTransform(Point(-2, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))}
	sphereMaterial := DefaultMaterial()
	sphereMaterial.color = Color{0.2, 1, 1}
	pattern1 := MakePerturbPattern(MakeRingPattern(Color{0.2, 1, 1}, Color{.8, 1, .8}, Identity().Scale(.05, .05, .05).RotateX(-math.Pi/3)))
	pattern2 := MakeGradientPattern(Color{1, .5, .3}, Color{.3, .5, 1}, Identity().RotateZ(-math.Pi/3))
	sphereMaterial.pattern = MakeBlendPattern(pattern1, pattern2)
	sphereMaterial.reflective = 0
	sphere := MakeSphere(Identity().Translate(-0.5, 1, 0.5), sphereMaterial)

	sphere2Material := DefaultMaterial()
	sphere2Material.color = Color{0.5, 1, 0.1}
	sphere2Material.diffuse = 0.7
	sphere2Material.specular = 0.1
	sphere2Material.reflective = 0

	gradientPattern := MakeGradientPattern(Color{.3, .5, 1}, Color{.5, 1, .3}, Identity())
	noisePattern := MakePerturbPattern(gradientPattern)
	sphere2Material.pattern = noisePattern
	sphere2 := MakeSphere(Identity().Translate(1.5, 0.7, -0.5).Scale(0.7, 0.7, 0.7), sphere2Material)

	sphere3Material := DefaultMaterial()

	sphere3Material.color = Color{1, 0.8, 0.1}
	sphere3Material.diffuse = 0
	sphere3Material.specular = 1
	// sphere3Material.transparency = 1
	sphere3Material.pattern = MakePerturbPattern(MakeStripePattern(Color{.7, .9, .5}, Color{.7, .7, .3}, Identity().Scale(.1, .1, .1).RotateZ(math.Pi/2).RotateY(math.Pi/4)))
	sphere3 := MakeSphere(Identity().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33), sphere3Material)

	floorMaterial := DefaultMaterial()
	floorMaterial.color = Color{1, 0.9, 0.9}
	floorMaterial.specular = 0.3
	floorMaterial.pattern = MakeCheckersPattern(Color{.5, .5, .5}, Color{.7, .7, .7}, Identity())
	floorMaterial.reflective = .1
	floor := MakePlane(Identity(), floorMaterial)

	wall1Material := DefaultMaterial()
	wall1Material.color = Color{.1, .1, .1}
	// wall1Material.pattern = MakeStripePattern(Color{0, 0, 0}, Color{.3, .3, .3}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	wall1Material.specular = 1
	wall1Material.reflective = 1
	wall1 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(-math.Pi/24).Translate(0, -3, 0), wall1Material)

	// wall2Material := DefaultMaterial()
	// wall2Material.color = Color{.1, .1, .1}
	// wall2Material.pattern = MakeStripePattern(Color{.7, .7, .7}, Color{.5, .5, .5}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	// wall2Material.specular = 1
	// wall2Material.reflective = .5
	// wall2 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(math.Pi/3).Translate(0, -4, 0), wall2Material)

	wall3Material := DefaultMaterial()
	wall3Material.color = Color{1, 1, 1}
	// wall3Material.pattern = MakeStripePattern(Color{.7, .7, .7}, Color{.3, .3, .3}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	wall3Material.specular = 1
	wall3Material.reflective = 1
	wall3 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(-math.Pi/24).Translate(0, 5, 5), wall3Material)

	// wall4Material := DefaultMaterial()
	// wall4Material.color = Color{.1, .1, .1}
	// wall4Material.pattern = MakeStripePattern(Color{.7, .7, .7}, Color{.5, .5, .5}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	// wall4Material.specular = 1
	// wall4Material.reflective = .5
	// wall4 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(math.Pi/3).Translate(0, 15, 15), wall4Material)

	lightPosition := Point(-4, 5, -4)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	world := World{light, []Object{sphere, sphere2, sphere3, floor, wall1, wall3}}
	c := camera.Render(world)
	fmt.Printf(c.PPM())
}
