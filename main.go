package main

import (
	"fmt"
	"math"
)

func main() {
	camera := Camera{1000, 500, math.Pi / 3, ViewTransform(Point(-1.3, 1.5, -5), Point(0, 1, 0), Vector(0, 1, 0))}
	sphereMaterial := DefaultMaterial()
	pattern1 := MakePerturbPattern(MakeRingPattern(Color{0.2, 1, 1}, Color{.8, 1, .8}, Identity().Scale(.05, .05, .05).RotateX(-math.Pi/3)))
	pattern2 := MakeGradientPattern(Color{1, .5, .3}, Color{.3, .5, 1}, Identity().RotateZ(-math.Pi/3))
	sphereMaterial.pattern = MakeBlendPattern(pattern1, pattern2)

	sphere := MakeSphere(Identity().Translate(-0.5, 1, 0.5), sphereMaterial)

	cylinderMaterial := DefaultMaterial()
	cylinderMaterial.diffuse = 1
	cylinderMaterial.specular = 0.1
	cylinderMaterial.ambient = 0
	cylinderMaterial.reflective = 0

	gradientPattern := MakeGradientPattern(Color{0, 0.3, 0.7}, Color{0.5, 0.7, 0.3}, Identity())
	noisePattern := MakePerturbPattern(gradientPattern)
	cylinderMaterial.pattern = noisePattern

	cylinder := MakeClosedCylinder(Identity().Translate(1.8, 0, -0.5).Scale(0.7, 0.7, 0.7), cylinderMaterial, 0, 1)

	hexagonMaterial := DefaultMaterial()

	hexagonMaterial.color = Color{1, 0.8, 0.1}
	hexagonMaterial.specular = 1
	hexagonMaterial.pattern = MakePerturbPattern(MakeStripePattern(Color{.7, .9, .5}, Color{.7, .7, .3}, Identity().Scale(.1, .1, .1).RotateZ(math.Pi/2).RotateY(math.Pi/4)))
	hex := hexagon(Identity().Translate(-1.7, 0.125, 0).Scale(0.5, 0.5, 0.5), hexagonMaterial)

	cubeMaterial := DefaultMaterial()

	cubeMaterial.color = Color{.5, .3, .1}
	cubeMaterial.specular = 1
	cubeMaterial.transparency = 1
	cubeMaterial.reflective = 1
	cubeMaterial.refractiveIndex = 1
	cubeMaterial.shininess = 300
	cube := MakeCube(Identity().Translate(-2.5, 0.25, 1.5).Scale(0.25, 0.25, 0.25), cubeMaterial)

	sphere5Material := DefaultMaterial()
	sphere5Material.color = Color{.3, 0.3, 0.3}
	sphere5Material.specular = 1
	sphere5Material.transparency = 1
	sphere5Material.reflective = 1
	sphere5Material.refractiveIndex = 1.5
	sphere5Material.shininess = 300

	sphere5 := MakeSphere(Identity().Translate(0.3, 0.5, -0.9).Scale(0.5, 0.5, 0.5), sphere5Material)

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
	wall2Material.specular = 1
	wall2Material.reflective = .5
	wall2 := MakePlane(Identity().RotateX(-math.Pi/2).RotateZ(math.Pi/3).Translate(0, -4, 0), wall2Material)

	wall3Material := DefaultMaterial()
	wall3Material.color = Color{0, 0, 0}
	wall3Material.specular = 1
	wall3Material.reflective = 1
	wall3 := MakePlane(Identity().RotateX(math.Pi/2).RotateZ(-math.Pi/24).Translate(0, -10, 0), wall3Material)

	lightPosition := Point(-4, 10, -4)
	lightColor := Color{1, 1, 1}
	light := PointLight{lightPosition, lightColor}

	world := World{light, []Object{sphere, hex, sphere5, cylinder, cube, wall1, wall2, wall3, floor}}
	c := camera.Render(world)
	fmt.Printf(c.PPM())
}

func hexagonCorner(g *Group, material Material) Object {
	return MakeSphereInGroup(Identity().Translate(0, 0, -1).Scale(.25, .25, .25), material, g)
}

func hexagonEdge(g *Group, material Material) Object {
	transform := Identity().Translate(0, 0, -1).RotateY(-math.Pi/6).RotateZ(-math.Pi/2).Scale(0.25, 1, 0.25)
	return MakeCylinderInGroup(transform, material, 0, 1, false, g)
}

func hexagonSide(g *Group, transform Matrix, material Material) *Group {
	side := MakeGroupInGroup(transform, g)
	hexagonCorner(side, material)
	hexagonEdge(side, material)

	return side
}

func hexagon(transform Matrix, material Material) Object {
	hex := MakeGroup(transform)

	for i := 0; i < 6; i++ {
		hexagonSide(hex, Identity().RotateY(float64(i)*math.Pi/3), material)
	}

	return hex
}
