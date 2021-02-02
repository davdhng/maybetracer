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

type Hittable interface {
	Hit(ray Ray, tMin float64, tMax float64, rec HitRecord) bool
}

type HittableList []Sphere

func (hl HittableList) Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	hitAnything := false
	closestSoFar := tMax
	// var tempRec HitRecord
	var closestRecord *HitRecord = nil

	for i := 0; i < len(hl); i++ {
		if hit, record := hl[i].Hit(r, tMin, closestSoFar); hit {
			closestRecord = record
			hitAnything = true
			// closestSoFar = tempRec.t
			closestSoFar = record.t
			// rec = tempRec
			// fmt.Println(rec.normal)
		}
	}
	return hitAnything, closestRecord
}

func (hl *HittableList) Add(spheres ...Sphere) {
	*hl = append(*hl, spheres...)
}

func (rec *HitRecord) setFaceNormal(r Ray, outwardNormal Vec3) {
	frontFace := r.Direction.Dot(outwardNormal) < 0
	if frontFace == true {
		rec.normal = outwardNormal
	} else {
		rec.normal = outwardNormal.Scale(-1.0)
	}
}

func (s Sphere) Hit(r Ray, tMin float64, tMax float64) (bool, *HitRecord) {
	oc := r.Origin.Sub(s.Center)
	a := r.Direction.length_squared()
	halfB := oc.Dot(r.Direction)
	c := oc.length_squared() - s.Radius*s.Radius

	discriminant := halfB*halfB - a*c

	if discriminant > 0 {
		root := math.Sqrt(discriminant)
		t := (-halfB - root) / a
		if t < tMax && t > tMin {
			pos := r.At(t)
			outwardNormal := (pos.Sub(s.Center)).Div(s.Radius)
			// rec.setFaceNormal(r, outwardNormal)
			// return true, &rec
			return true, &HitRecord{t: t, p: pos, normal: outwardNormal}
		}
		t = (-halfB + root) / a
		if t < tMax && t > tMin {
			pos := r.At(t)
			outwardNormal := (pos.Sub(s.Center)).Div(s.Radius)
			// rec.setFaceNormal(r, outwardNormal)
			//return true, &rec
			return true, &HitRecord{t: t, p: pos, normal: outwardNormal}
		}
	}
	return false, nil

	// if discriminant < 0 {
	// 	return false
	// }
	// sqrtd := math.Sqrt(discriminant)

	// root := (-halfB - sqrtd) / a
	// if root < tMin || tMax < root {
	// 	root = (-halfB + sqrtd) / a
	// 	if root < tMin || tMax < root {
	// 		return false
	// 	}
	// }
	// rec.t = root
	// rec.p = r.At(rec.t)
	// outwardNormal := (rec.p.Sub(s.Center)).Div(s.Radius)
	// rec.setFaceNormal(r, outwardNormal)
	// return true
}
