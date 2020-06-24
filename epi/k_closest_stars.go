package epi

import (
	"fmt"
	"math"
)

type Star struct {
	X float64
	Y float64
	Z float64
}

func (s Star) Distance() float64 {
	return math.Sqrt(s.X*s.X + s.Y*s.Y + s.Z*s.Z)
}

func (s Star) String() string {
	return fmt.Sprintf("%f", s.Distance())
}

func FindClosestKStars(stars []Star, k int) []Star {
	// TODO - Add your code here
	return nil
}
