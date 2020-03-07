package geo

import (
	"math"
)

// Point represents a geographic 2D coordinate
type Point struct {
	Lat float64
	Lon float64
}

// Vincenty calculates the great circle distance between Points x and y
// The distance is returned in meters, x and y must be specified in degrees
func Vincenty(x, y Point) float64 {

	sq := func(x float64) float64 { return x * x }
	degToRad := func(x float64) float64 { return x * math.Pi / 180 }

	var lambda, tmp, q, p float64 = 0, 0, 0, 0
	var sigma, sinSigma, cosSigma float64 = 0, 0, 0
	var sinAlpha, cos2Alpha, cos2Sigma float64 = 0, 0, 0
	var c float64 = 0

	A := 6378137.0
	F := 1 / 298.257223563
	B := (1 - F) * A
	C := (sq(A) - sq(B)) / sq(B)

	uX := math.Atan((1 - F) * math.Tan(degToRad(x.Lat)))
	sinUX := math.Sin(uX)
	cosUX := math.Cos(uX)

	uY := math.Atan((1 - F) * math.Tan(degToRad(y.Lat)))
	sinUY := math.Sin(uY)
	cosUY := math.Cos(uY)

	l := degToRad(y.Lon) - degToRad(x.Lon)

	lambda = l

	for i := 0; i < 10; i++ {

		tmp = math.Cos(lambda)
		q = cosUY * math.Sin(lambda)
		p = cosUX*sinUY - sinUX*cosUY*tmp

		sinSigma = math.Sqrt(q*q + p*p)
		cosSigma = sinUX*sinUY + cosUX*cosUY*tmp
		sigma = math.Atan2(sinSigma, cosSigma)

		sinAlpha = (cosUX * cosUY * math.Sin(lambda)) / sinSigma
		cos2Alpha = 1 - sq(sinAlpha)
		cos2Sigma = cosSigma - (2*sinUX*sinUY)/cos2Alpha

		c = F / 16.0 * cos2Alpha * (4 + F*(4-3*cos2Alpha))
		tmp = lambda
		lambda = (l + (1-c)*F*sinAlpha*(sigma+c*sinSigma*(cos2Sigma+c*cosSigma*(-1+2*cos2Sigma*cos2Sigma))))

		if math.Abs(lambda-tmp) < 0.00000001 {
			break
		}
	}

	uu := cos2Alpha * C
	a := 1 + uu/16384*(4096+uu*(-768+uu*(320-175*uu)))
	b := uu / 1024 * (256 + uu*(-128+uu*(74-47*uu)))

	deltaSigma := (b * sinSigma * (cos2Sigma + 1.0/4.0*b*(cosSigma*(-1+2*sq(cos2Sigma))*(-3+4*sq(sinSigma))*(-3+4*sq(cos2Sigma)))))

	return B * a * (sigma - deltaSigma)
}
