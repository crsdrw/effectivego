/*
Package math provides 3D vector types and utilities.
*/
package math

// Vector is a slice of 3 floats
type Vector []float64

// Sub subtracts one vector from another. If an out
// parameter is provided of sufficient size then it
// will be used for the result to avoid allocation.
func Sub(v1, v2, out Vector) Vector {
	if len(out) < 3 {
		out = make(Vector, 3)
	}
	out[0] = v1[0] - v2[0]
	out[1] = v1[1] - v2[1]
	out[2] = v1[2] - v2[2]
	return out
}

// Add one vector to another. If an out parameter
// is provided of sufficient size then it will be
// used for the result to avoid allocation.
func Add(v1, v2, out Vector) Vector {
	if len(out) < 3 {
		out = make(Vector, 3)
	}
	out[0] = v1[0] + v2[0]
	out[1] = v1[1] + v2[1]
	out[2] = v1[2] + v2[2]
	return out
}

// Cross returns the cross-product of two vectors.
// If an out parameter is provided of sufficient
// size then it will be used for the result to
// avoid allocation.
func Cross(v1, v2, out Vector) Vector {
	if len(out) < 3 {
		out = make(Vector, 3)
	}
	out[0] = v1[1]*v2[2] - v1[2]*v2[1]
	out[1] = v1[2]*v2[0] - v1[0]*v2[2]
	out[2] = v1[0]*v2[1] - v1[1]*v2[0]
	return out
}

// Dot returns the dot-product of two vectors.
func Dot(v1, v2 Vector) float64 {
	return v1[1]*v2[1] + v1[2]*v2[2] + v1[0]*v2[0]
}

// Mul multplies a vector by a scalar. The multipication
// is performed in-place.
func Mul(v Vector, k float64) {
	for i := range v {
		v[i] *= k
	}
}
