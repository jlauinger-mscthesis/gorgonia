package gorgonia

import (
	"testing"

	"gonum.org/v1/gonum/blas"
	"gonum.org/v1/gonum/blas/gonum"
	"gorgonia.org/tensor"
)

var gonumImpl = gonum.Implementation{}

// testBLASImplementation of the interface
type testBLASImplementation struct {
	used bool
}

// Sdsdot computes the dot product of the two vectors plus a constant
//  alpha + \sum_i x[i]*y[i]
//
// Float32 implementations are autogenerated and not directly tested.
// Sdsdot ...
func (*testBLASImplementation) Sdsdot(n int, alpha float32, x []float32, incX int, y []float32, incY int) float32 {
	return gonumImpl.Sdsdot(n, alpha, x, incX, y, incY)
}

// Dsdot computes the dot product of the two vectors
//  \sum_i x[i]*y[i]
//
// Float32 implementations are autogenerated and not directly tested.
// Dsdot ...
func (*testBLASImplementation) Dsdot(n int, x []float32, incX int, y []float32, incY int) float64 {
	return gonumImpl.Dsdot(n, x, incX, y, incY)
}

// Sdot ...
func (*testBLASImplementation) Sdot(n int, x []float32, incX int, y []float32, incY int) float32 {
	return gonumImpl.Sdot(n, x, incX, y, incY)
}

// Snrm2 ...
func (*testBLASImplementation) Snrm2(n int, x []float32, incX int) float32 {
	return gonumImpl.Snrm2(n, x, incX)
}

// Sasum ...
func (*testBLASImplementation) Sasum(n int, x []float32, incX int) float32 {
	return gonumImpl.Sasum(n, x, incX)
}

// Isamax ...
func (*testBLASImplementation) Isamax(n int, x []float32, incX int) int {
	return gonumImpl.Isamax(n, x, incX)
}

// Sswap ...
func (*testBLASImplementation) Sswap(n int, x []float32, incX int, y []float32, incY int) {
	gonumImpl.Sswap(n, x, incX, y, incY)
}

// Scopy ...
func (*testBLASImplementation) Scopy(n int, x []float32, incX int, y []float32, incY int) {
	gonumImpl.Scopy(n, x, incX, y, incY)
}

// Saxpy ...
func (*testBLASImplementation) Saxpy(n int, alpha float32, x []float32, incX int, y []float32, incY int) {
	gonumImpl.Saxpy(n, alpha, x, incX, y, incY)
}

// Srotg ...
func (*testBLASImplementation) Srotg(a float32, b float32) (c float32, s float32, r float32, z float32) {
	return gonumImpl.Srotg(a, b)
}

// Srotmg ...
func (*testBLASImplementation) Srotmg(d1 float32, d2 float32, b1 float32, b2 float32) (p blas.SrotmParams, rd1 float32, rd2 float32, rb1 float32) {
	return gonumImpl.Srotmg(d1, d2, b1, b2)
}

// Srot ...
func (*testBLASImplementation) Srot(n int, x []float32, incX int, y []float32, incY int, c float32, s float32) {
	gonumImpl.Srot(n, x, incX, y, incY, c, s)
}

// Srotm ...
func (*testBLASImplementation) Srotm(n int, x []float32, incX int, y []float32, incY int, p blas.SrotmParams) {
	gonumImpl.Srotm(n, x, incX, y, incY, p)
}

// Sscal ...
func (*testBLASImplementation) Sscal(n int, alpha float32, x []float32, incX int) {
	gonumImpl.Sscal(n, alpha, x, incX)
}

// Sgemv ...
func (*testBLASImplementation) Sgemv(tA blas.Transpose, m int, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	gonumImpl.Sgemv(tA, m, n, alpha, a, lda, x, incX, beta, y, incY)
}

// Sgbmv ...
func (*testBLASImplementation) Sgbmv(tA blas.Transpose, m int, n int, kL int, kU int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	gonumImpl.Sgbmv(tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)
}

// Strmv ...
func (*testBLASImplementation) Strmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	gonumImpl.Strmv(ul, tA, d, n, a, lda, x, incX)
}

// Stbmv ...
func (*testBLASImplementation) Stbmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	gonumImpl.Stbmv(ul, tA, d, n, k, a, lda, x, incX)
}

// Stpmv ...
func (*testBLASImplementation) Stpmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	gonumImpl.Stpmv(ul, tA, d, n, ap, x, incX)
}

// Strsv ...
func (*testBLASImplementation) Strsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float32, lda int, x []float32, incX int) {
	gonumImpl.Strsv(ul, tA, d, n, a, lda, x, incX)
}

// Stbsv ...
func (*testBLASImplementation) Stbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float32, lda int, x []float32, incX int) {
	gonumImpl.Stbsv(ul, tA, d, n, k, a, lda, x, incX)
}

// Stpsv ...
func (*testBLASImplementation) Stpsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float32, x []float32, incX int) {
	gonumImpl.Stpsv(ul, tA, d, n, ap, x, incX)

}

// Ssymv ...
func (*testBLASImplementation) Ssymv(ul blas.Uplo, n int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	gonumImpl.Ssymv(ul, n, alpha, a, lda, x, incX, beta, y, incY)

}

// Ssbmv ...
func (*testBLASImplementation) Ssbmv(ul blas.Uplo, n int, k int, alpha float32, a []float32, lda int, x []float32, incX int, beta float32, y []float32, incY int) {
	gonumImpl.Ssbmv(ul, n, k, alpha, a, lda, x, incX, beta, y, incY)

}

// Sspmv ...
func (*testBLASImplementation) Sspmv(ul blas.Uplo, n int, alpha float32, ap []float32, x []float32, incX int, beta float32, y []float32, incY int) {
	gonumImpl.Sspmv(ul, n, alpha, ap, x, incX, beta, y, incY)

}

// Sger ...
func (*testBLASImplementation) Sger(m int, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	gonumImpl.Sger(m, n, alpha, x, incX, y, incY, a, lda)

}

// Ssyr ...
func (*testBLASImplementation) Ssyr(ul blas.Uplo, n int, alpha float32, x []float32, incX int, a []float32, lda int) {
	gonumImpl.Ssyr(ul, n, alpha, x, incX, a, lda)

}

// Sspr ...
func (*testBLASImplementation) Sspr(ul blas.Uplo, n int, alpha float32, x []float32, incX int, ap []float32) {
	gonumImpl.Sspr(ul, n, alpha, x, incX, ap)

}

// Ssyr2 ...
func (*testBLASImplementation) Ssyr2(ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32, lda int) {
	gonumImpl.Ssyr2(ul, n, alpha, x, incX, y, incY, a, lda)

}

// Sspr2 ...
func (*testBLASImplementation) Sspr2(ul blas.Uplo, n int, alpha float32, x []float32, incX int, y []float32, incY int, a []float32) {
	gonumImpl.Sspr2(ul, n, alpha, x, incX, y, incY, a)

}

// Ssymm ...
func (*testBLASImplementation) Ssymm(s blas.Side, ul blas.Uplo, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	gonumImpl.Ssymm(s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)

}

// Ssyrk ...
func (*testBLASImplementation) Ssyrk(ul blas.Uplo, t blas.Transpose, n int, k int, alpha float32, a []float32, lda int, beta float32, c []float32, ldc int) {
	gonumImpl.Ssyrk(ul, t, n, k, alpha, a, lda, beta, c, ldc)

}

// Ssyr2k ...
func (*testBLASImplementation) Ssyr2k(ul blas.Uplo, t blas.Transpose, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	gonumImpl.Ssyr2k(ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)

}

// Strmm ...
func (*testBLASImplementation) Strmm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	gonumImpl.Strmm(s, ul, tA, d, m, n, alpha, a, lda, b, ldb)

}

// Strsm ...
func (*testBLASImplementation) Strsm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float32, a []float32, lda int, b []float32, ldb int) {
	gonumImpl.Strsm(s, ul, tA, d, m, n, alpha, a, lda, b, ldb)

}

// Ddot ...
func (*testBLASImplementation) Ddot(n int, x []float64, incX int, y []float64, incY int) float64 {
	return gonumImpl.Ddot(n, x, incX, y, incY)

}

// Dnrm2 ...
func (*testBLASImplementation) Dnrm2(n int, x []float64, incX int) float64 {
	return gonumImpl.Dnrm2(n, x, incX)

}

// Dasum ...
func (*testBLASImplementation) Dasum(n int, x []float64, incX int) float64 {
	return gonumImpl.Dasum(n, x, incX)

}

// Idamax ...
func (*testBLASImplementation) Idamax(n int, x []float64, incX int) int {
	return gonumImpl.Idamax(n, x, incX)

}

// Dswap ...
func (*testBLASImplementation) Dswap(n int, x []float64, incX int, y []float64, incY int) {
	gonumImpl.Dswap(n, x, incX, y, incY)

}

// Dcopy ...
func (*testBLASImplementation) Dcopy(n int, x []float64, incX int, y []float64, incY int) {
	gonumImpl.Dcopy(n, x, incX, y, incY)

}

// Daxpy ...
func (*testBLASImplementation) Daxpy(n int, alpha float64, x []float64, incX int, y []float64, incY int) {
	gonumImpl.Daxpy(n, alpha, x, incX, y, incY)

}

// Drotg ...
func (*testBLASImplementation) Drotg(a float64, b float64) (c float64, s float64, r float64, z float64) {
	return gonumImpl.Drotg(a, b)

}

// Drotmg ...
func (*testBLASImplementation) Drotmg(d1 float64, d2 float64, b1 float64, b2 float64) (p blas.DrotmParams, rd1 float64, rd2 float64, rb1 float64) {
	return gonumImpl.Drotmg(d1, d2, b1, b2)

}

// Drot ...
func (*testBLASImplementation) Drot(n int, x []float64, incX int, y []float64, incY int, c float64, s float64) {
	gonumImpl.Drot(n, x, incX, y, incY, c, s)

}

// Drotm ...
func (*testBLASImplementation) Drotm(n int, x []float64, incX int, y []float64, incY int, p blas.DrotmParams) {
	gonumImpl.Drotm(n, x, incX, y, incY, p)

}

// Dscal ...
func (*testBLASImplementation) Dscal(n int, alpha float64, x []float64, incX int) {
	gonumImpl.Dscal(n, alpha, x, incX)

}

// Dgemv ...
func (*testBLASImplementation) Dgemv(tA blas.Transpose, m int, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	gonumImpl.Dgemv(tA, m, n, alpha, a, lda, x, incX, beta, y, incY)

}

// Dgbmv ...
func (*testBLASImplementation) Dgbmv(tA blas.Transpose, m int, n int, kL int, kU int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	gonumImpl.Dgbmv(tA, m, n, kL, kU, alpha, a, lda, x, incX, beta, y, incY)

}

// Dtrmv ...
func (*testBLASImplementation) Dtrmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	gonumImpl.Dtrmv(ul, tA, d, n, a, lda, x, incX)

}

// Dtbmv ...
func (*testBLASImplementation) Dtbmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	gonumImpl.Dtbmv(ul, tA, d, n, k, a, lda, x, incX)

}

// Dtpmv ...
func (*testBLASImplementation) Dtpmv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	gonumImpl.Dtpmv(ul, tA, d, n, ap, x, incX)

}

// Dtrsv ...
func (*testBLASImplementation) Dtrsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, a []float64, lda int, x []float64, incX int) {
	gonumImpl.Dtrsv(ul, tA, d, n, a, lda, x, incX)

}

// Dtbsv ...
func (*testBLASImplementation) Dtbsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, k int, a []float64, lda int, x []float64, incX int) {
	gonumImpl.Dtbsv(ul, tA, d, n, k, a, lda, x, incX)

}

// Dtpsv ...
func (*testBLASImplementation) Dtpsv(ul blas.Uplo, tA blas.Transpose, d blas.Diag, n int, ap []float64, x []float64, incX int) {
	gonumImpl.Dtpsv(ul, tA, d, n, ap, x, incX)

}

// Dsymv ...
func (*testBLASImplementation) Dsymv(ul blas.Uplo, n int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	gonumImpl.Dsymv(ul, n, alpha, a, lda, x, incX, beta, y, incY)

}

// Dsbmv ...
func (*testBLASImplementation) Dsbmv(ul blas.Uplo, n int, k int, alpha float64, a []float64, lda int, x []float64, incX int, beta float64, y []float64, incY int) {
	gonumImpl.Dsbmv(ul, n, k, alpha, a, lda, x, incX, beta, y, incY)

}

// Dspmv ...
func (*testBLASImplementation) Dspmv(ul blas.Uplo, n int, alpha float64, ap []float64, x []float64, incX int, beta float64, y []float64, incY int) {
	gonumImpl.Dspmv(ul, n, alpha, ap, x, incX, beta, y, incY)

}

// Dger ...
func (*testBLASImplementation) Dger(m int, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	gonumImpl.Dger(m, n, alpha, x, incX, y, incY, a, lda)

}

// Dsyr ...
func (*testBLASImplementation) Dsyr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, a []float64, lda int) {
	gonumImpl.Dsyr(ul, n, alpha, x, incX, a, lda)

}

// Dspr ...
func (*testBLASImplementation) Dspr(ul blas.Uplo, n int, alpha float64, x []float64, incX int, ap []float64) {
	gonumImpl.Dspr(ul, n, alpha, x, incX, ap)

}

// Dsyr2 ...
func (*testBLASImplementation) Dsyr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64, lda int) {
	gonumImpl.Dsyr2(ul, n, alpha, x, incX, y, incY, a, lda)

}

// Dspr2 ...
func (*testBLASImplementation) Dspr2(ul blas.Uplo, n int, alpha float64, x []float64, incX int, y []float64, incY int, a []float64) {
	gonumImpl.Dspr2(ul, n, alpha, x, incX, y, incY, a)

}

// Dsymm ...
func (*testBLASImplementation) Dsymm(s blas.Side, ul blas.Uplo, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	gonumImpl.Dsymm(s, ul, m, n, alpha, a, lda, b, ldb, beta, c, ldc)

}

// Dsyrk ...
func (*testBLASImplementation) Dsyrk(ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, beta float64, c []float64, ldc int) {
	gonumImpl.Dsyrk(ul, t, n, k, alpha, a, lda, beta, c, ldc)

}

// Dsyr2k ...
func (*testBLASImplementation) Dsyr2k(ul blas.Uplo, t blas.Transpose, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	gonumImpl.Dsyr2k(ul, t, n, k, alpha, a, lda, b, ldb, beta, c, ldc)

}

// Dtrmm ...
func (*testBLASImplementation) Dtrmm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	gonumImpl.Dtrmm(s, ul, tA, d, m, n, alpha, a, lda, b, ldb)

}

// Dtrsm ...
func (*testBLASImplementation) Dtrsm(s blas.Side, ul blas.Uplo, tA blas.Transpose, d blas.Diag, m int, n int, alpha float64, a []float64, lda int, b []float64, ldb int) {
	gonumImpl.Dtrsm(s, ul, tA, d, m, n, alpha, a, lda, b, ldb)

}

// Sgemm ...
func (t *testBLASImplementation) Sgemm(tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float32, a []float32, lda int, b []float32, ldb int, beta float32, c []float32, ldc int) {
	t.used = true
	gonumImpl.Sgemm(tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)

}

// Dgemm ...
func (*testBLASImplementation) Dgemm(tA blas.Transpose, tB blas.Transpose, m int, n int, k int, alpha float64, a []float64, lda int, b []float64, ldb int, beta float64, c []float64, ldc int) {
	gonumImpl.Dgemm(tA, tB, m, n, k, alpha, a, lda, b, ldb, beta, c, ldc)

}
func TestUse(t *testing.T) {
	blasI := &testBLASImplementation{}
	Use(blasI)
	g := NewGraph()
	x := NodeFromAny(g, tensor.New(
		tensor.WithShape(1, 1, 7, 5),
		tensor.WithBacking([]float32{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34})))
	filter := NodeFromAny(g, tensor.New(
		tensor.WithShape(1, 1, 3, 3),
		tensor.WithBacking([]float32{1, 1, 1, 1, 1, 1, 1, 1, 1})))
	y := Must(Conv2d(x, filter, []int{3, 3}, []int{0, 0}, []int{2, 2}, []int{1, 1}))
	m := NewTapeMachine(g)
	if err := m.RunAll(); err != nil {
		t.Fatal(err)
	}
	//54 72 144 162 234 252
	output := y.Value().Data().([]float32)
	if output[0] != 54 ||
		output[1] != 72 ||
		output[2] != 144 ||
		output[3] != 162 ||
		output[4] != 234 ||
		output[5] != 252 {
		t.Fatal("wrong computation value")
	}

	if !blasI.used {
		t.Fail()
	}
	if WhichBLAS() != blasI {
		t.Fail()
	}
}
