package main

import "math"

type HitRecord struct {
	p         Vec3
	normal    Vec3
	t         float64
	frontFace bool
}

type Sphere struct {
	Center Vec3
	Radius float64
}

type hittable struct{}

func (rec HitRecord) setFaceNormal(r Ray, outwardNormal Vec3) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	if frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.Scale(-1)
	}
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64, rec HitRecord) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.length_squared()
	half_b := oc.Dot(r.Direction)
	c := oc.length_squared() - s.Radius*s.Radius
	discriminant := half_b*half_b - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)
	root := (-half_b - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-half_b + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}
	rec.t = root
	rec.p = r.At(rec.t)
	// rec.normal = (rec.p.Sub(s.Center)).Scale(1 / s.Radius)
	outwardNormal := (rec.p.Sub(s.Center).Scale(1 / s.Radius))
	rec.setFaceNormal(r, outwardNormal)
	return true

}
