package main

type Camera struct {
	aspectRatio     float64
	viewportHeight  float64
	viewportWidth   float64
	focalLength     float64
	origin          Vec3
	horizontal      Vec3
	vertical        Vec3
	lowerLeftCorner Vec3
}

func NewCamera() Camera {
	cam := Camera{}
	cam.aspectRatio = 16.0 / 9.0
	cam.viewportHeight = 2.0
	cam.viewportWidth = cam.aspectRatio * cam.viewportHeight
	cam.focalLength = 1.0
	cam.origin = Vec3{0, 0, 0}
	cam.horizontal = Vec3{cam.viewportWidth, 0.0, 0.0}
	cam.vertical = Vec3{0.0, cam.viewportHeight, 0.0}
	cam.lowerLeftCorner = cam.origin.Sub(cam.horizontal.Div(2)).Sub(cam.vertical.Div(2)).Sub(Vec3{0, 0, cam.focalLength})
	return cam
}

func (c Camera) getRay(u float64, v float64) Ray {
	hori := c.horizontal.Scale(u)
	vert := c.vertical.Scale(v)
	return Ray{
		c.origin,
		c.lowerLeftCorner.Add(hori).Add(vert).Sub(c.origin),
	}
}
