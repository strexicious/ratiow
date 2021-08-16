package vec

type Ray struct {
	orig Point3
	dir  Vec3
}

func ZeroRay() Ray {
	return Ray{orig: ZeroPoint3(), dir: ZeroVec3()}
}

func NewRay(origin Point3, direction Vec3) Ray {
	return Ray{orig: origin, dir: direction}
}

func (r Ray) Origin() Point3 {
	return r.orig
}

func (r Ray) Direction() Vec3 {
	return r.dir
}

func (r Ray) At(t float64) Point3 {
	return r.orig.Add(r.dir.Scale(t))
}
