package mat3

import (
	"github.com/ungerik/go3d/vec3"
	"testing"
)

const EPSILON = 0.0001

// Some matrices used in multiple tests.
var (
	TEST_MATRIX1 = T{
		vec3.T{0.38016528, -0.0661157, -0.008264462},
		vec3.T{-0.19834709, 0.33884296, -0.08264463},
		vec3.T{0.11570247, -0.28099173, 0.21487603},
	}

	TEST_MATRIX2 = T{
		vec3.T{23, -4, -0.5},
		vec3.T{-12, 20.5, -5},
		vec3.T{7, -17, 13},
	}
)

func BenchmarkAssignMul(b *testing.B) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	var mMult T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		mMult.AssignMul(&m1, &m2)
	}
}

// func BenchmarkMultMatrix(b *testing.B) {
// 	m1 := TEST_MATRIX1
// 	m2 := TEST_MATRIX2
// 	b.ResetTimer()
// 	for i := 0; i < b.N; i++ {
// 		m1.MultMatrix(&m2)
// 	}
// }

func BenchmarkMulVec3(b *testing.B) {
	m1 := TEST_MATRIX1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec3.T{1, 1.5, 2}
		v_1 := m1.MulVec3(&v)
		m1.MulVec3(&v_1)
	}
}

func BenchmarkTransformVec3(b *testing.B) {
	m1 := TEST_MATRIX1
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec3.T{1, 1.5, 2}
		m1.TransformVec3(&v)
		m1.TransformVec3(&v)
	}
}

func (mat *T) TransformVec3_PassByValue(v vec3.T) (r vec3.T) {
	// Use intermediate variables to not alter further computations.
	x := mat[0][0]*v[0] + mat[1][0]*v[1] + mat[2][0]*v[2]
	y := mat[0][1]*v[0] + mat[1][1]*v[1] + mat[2][1]*v[2]
	z := mat[0][2]*v[0] + mat[1][2]*v[1] + mat[2][2]*v[2]
	r[0] = x
	r[1] = y
	r[2] = z
	return r
}

func Vec3Add_PassByValue(a, b vec3.T) vec3.T {
	return vec3.T{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func BenchmarkMulAddVec3_PassByPointer(b *testing.B) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	var v1 vec3.T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec3.T{1, 1.5, 2}
		m1.TransformVec3(&v)
		m2.TransformVec3(&v)
		v = vec3.Add(&v, &v1)
		v = vec3.Add(&v, &v1)
	}
}

// Demonstrate that
func BenchmarkMulAddVec3_PassByValue(b *testing.B) {
	m1 := TEST_MATRIX1
	m2 := TEST_MATRIX2
	var v1 vec3.T
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v := vec3.T{1, 1.5, 2}
		m1.TransformVec3_PassByValue(v)
		m2.TransformVec3_PassByValue(v)
		v = Vec3Add_PassByValue(v, v1)
		v = Vec3Add_PassByValue(v, v1)
	}
}
