package main

import (
	"math"
)

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

type HittableList []Sphere

func (hl HittableList) Hit(r Ray, tMin float64, tMax float64, rec HitRecord) bool {
	hitAnything := false
	closestSoFar := tMax
	var tempRec HitRecord
	for i := 0; i < len(hl); i++ {
		if hl[i].Hit(r, tMin, closestSoFar, tempRec) {
			hitAnything = true
			closestSoFar = tempRec.t
			rec = tempRec
		}
	}
	return hitAnything
}

func (hl *HittableList) Add(spheres ...Sphere) {
	*hl = append(*hl, spheres...)
}

func (rec HitRecord) setFaceNormal(r Ray, outwardNormal Vec3) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	if frontFace {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.Scale(-1.0)
	}
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64, rec HitRecord) bool {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.length_squared()
	halfB := oc.Dot(r.Direction)
	c := oc.length_squared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c
	if discriminant < 0 {
		return false
	}
	sqrtd := math.Sqrt(discriminant)

	root := (-halfB - sqrtd) / a
	if root < tMin || tMax < root {
		root = (-halfB + sqrtd) / a
		if root < tMin || tMax < root {
			return false
		}
	}
	rec.t = root
	rec.p = r.At(rec.t)
	outwardNormal := (rec.p.Sub(s.Center)).Div(s.Radius)
	rec.setFaceNormal(r, outwardNormal)
	return true
}
