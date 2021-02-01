package main

type Vec3 struct {
	x, y, z float64
}
type Color Vec3

func (v *Vec3) Add(v0, v1 *Vec3) {
	v.x = v0.x + v1.x
	v.y = v0.y + v1.y
	v.z = v0.z + v1.z
}

func (v *Vec3) Mul(v0, v1 *Vec3) {
	v.x = v0.x * v1.x
	v.y = v0.y * v1.y
	v.z = v0.z * v1.z
}

func (v *Vec3) Dot(v1 *Vec3) float64 {
	return v.x*v1.x + v.y*v1.y + v.z*v1.z
}
