package main

import (
	"math"
	"math/rand"
)

type Vec3 struct {
	x, y, z float64
}
type Color Vec3
type Point3 Vec3

func (v Vec3) Add(v1 Vec3) Vec3 {
	return Vec3{
		v.x + v1.x,
		v.y + v1.y,
		v.z + v1.z,
	}
}

func (v Vec3) Sub(v1 Vec3) Vec3 {
	return Vec3{
		v.x - v1.x,
		v.y - v1.y,
		v.z - v1.z,
	}
}

func (v *Vec3) Mul(v0, v1 *Vec3) {
	v.x = v0.x * v1.x
	v.y = v0.y * v1.y
	v.z = v0.z * v1.z
}

func (v Vec3) Dot(v1 Vec3) float64 {
	return v.x*v1.x + v.y*v1.y + v.z*v1.z
}

func (u Vec3) Scale(t float64) Vec3 {
	return Vec3{
		x: t * u.x,
		y: t * u.y,
		z: t * u.z,
	}
}

func (u Vec3) Div(t float64) Vec3 {
	if t == 0 {
		panic("Division by zero")
	}
	return Vec3{x: u.x / t, y: u.y / t, z: u.z / t}
}

func (v Vec3) length_squared() float64 {
	// return v.x*v.x + v.y*v.y + v.z*v.z
	return v.Dot(v)
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.length_squared())
}

func (v Vec3) unit_vector() Vec3 {
	return v.Div(v.length())
}

func (v Vec3) randomRange(min float64, max float64) Vec3 {
	return Vec3{
		rand.Float64()*(max-min) + min,
		rand.Float64()*(max-min) + min,
		rand.Float64()*(max-min) + min,
	}
}

func (v Vec3) random() Vec3 {
	return Vec3{
		rand.Float64(),
		rand.Float64(),
		rand.Float64(),
	}
}

// func (v Vec3) randSphere() Vec3 {
// 	for {
// 		p := v.randomRange(-1.0, 1.0)
// 		if p.length_squared() >= 1 {
// 			continue
// 		}
// 		return p
// 	}
// }
func RandSphere() Vec3 {

	a := 2.0 * rand.Float64() * math.Pi
	z := 2.0 * (rand.Float64() - 0.5)
	r := math.Sqrt(1 - z*z)
	return Vec3{
		r * math.Cos(a),
		r * math.Sin(a),
		z,
	}
}

func RandUnitVec() Vec3 {
	x := RandSphere()
	return x.unit_vector()
}
