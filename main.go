package main

import (
	"fmt"
	"math"
)

func main() {
	camera := Camera{250, 125, math.Pi / 3, ViewTransform(Point(-1.3, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))}
	sphereMaterial := DefaultMaterial()
	sphereMaterial.color = Color{0.2, 1, 1}
	pattern1 := MakePerturbPattern(MakeRingPattern(Color{0.2, 1, 1}, Color{.8, 1, .8}, Identity().Scale(.05, .05, .05).RotateX(-math.Pi/3)))
	pattern2 := MakeGradientPattern(Color{1, .5, .3}, Color{.3, .5, 1}, Identity().RotateZ(-math.Pi/3))
	sphereMaterial.pattern = MakeBlendPattern(pattern1, pattern2)
	sphereMaterial.reflective = 0
	sphere := MakeSphere(Identity().Translate(-0.5, 1, 0.5), sphereMaterial)

	sphere2Material := DefaultMaterial()
	sphere2Material.diffuse = 1
	sphere2Material.specular = 0.1
	sphere2Material.ambient = 0
	sphere2Material.reflective = 0
	sphere2Material.shininess = 1000

	gradientPattern := MakeGradientPattern(Color{0, 0.3, 0.7}, Color{0.5, 0.7, 0.3}, Identity())
	noisePattern := MakePerturbPattern(gradientPattern)
	sphere2Material.pattern = noisePattern
	sphere2 := MakeSphere(Identity().Translate(1.5, 0.7, -0.5).Scale(0.7, 0.7, 0.7), sphere2Material)

	sphere3Material := DefaultMaterial()

	sphere3Material.color = Color{1, 0.8, 0.1}
	sphere3Material.specular = 1
	sphere3Material.pattern = MakePerturbPattern(MakeStripePattern(Color{.7, .9, .5}, Color{.7, .7, .3}, Identity().Scale(.1, .1, .1).RotateZ(math.Pi/2).RotateY(math.Pi/4)))
	sphere3 := MakeSphere(Identity().Translate(-1.5, 0.33, -0.75).Scale(0.33, 0.33, 0.33), sphere3Material)

	sphere4Material := DefaultMaterial()

	sphere4Material.color = Color{1, 0.8, 0.1}
	sphere4Material.specular = 1
	sphere4 := MakeSphere(Identity().Translate(-2.5, 0.25, 1.5).Scale(0.25, 0.25, 0.25), sphere4Material)

	sphere5Material := DefaultMaterial()

	sphere5Material.color = Color{.7, 0.3, 0.3}
	sphere5Material.specular = .3
	sphere5 := MakeSphere(Identity().Translate(.5, 0.1, -0.5).Scale(0.1, 0.1, 0.1), sphere5Material)

	floorMaterial := DefaultMaterial()
	floorMaterial.specular = 1
	floorMaterial.pattern = MakeCheckersPattern(Color{.5, .5, .5}, Color{.7, .7, .7}, Identity())
	floorMaterial.reflective = .5
	floor := MakePlane(Identity(), floorMaterial)

	wall1Material := DefaultMaterial()
	wall1Material.color = Color{.1, .1, .1}
	wall1Material.pattern = MakeStripePattern(Color{.4, .4, .4}, Color{.3, .3, .3}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	wall1Material.specular = 1
	wall1Material.reflective = .5
	wall1 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(-math.Pi/24).Translate(0, -3, 0), wall1Material)

	wall2Material := DefaultMaterial()
	wall2Material.color = Color{.1, .1, .1}
	wall2Material.pattern = MakeStripePattern(Color{.3, .3, .3}, Color{.4, .4, .4}, Identity().RotateY(math.Pi/2).Scale(.2, .2, .2))
	wall2Material.specular = 0
	wall2Material.reflective = 0
	wall2 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(math.Pi/3).Translate(0, -4, 0), wall2Material)

	wall3Material := DefaultMaterial()
	wall3Material.color = Color{0, 0, 0}
	wall3Material.specular = 1
	wall3Material.reflective = 1
	wall3 := MakePlane(Identity().RotateX(math.Pi/2).RotateZ(-math.Pi/24).Translate(0, -10, 0), wall3Material)

	lightPosition := Point(-4, 10, -4)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	world := World{light, []Object{sphere, sphere2, sphere3, sphere4, sphere5, floor, wall1, wall2, wall3}}
	c := camera.Render(world)
	fmt.Printf(c.PPM())
}
