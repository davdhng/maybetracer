package main

import "math"

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
	return v.x*v.x + v.y*v.y + v.z*v.z
}

func (v Vec3) length() float64 {
	return math.Sqrt(v.length_squared())
}

func unit_vector(v Vec3) Vec3 {
	return v.Div(v.length())
}
