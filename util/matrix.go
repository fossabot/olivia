package util

import "gonum.org/v1/gonum/mat"

// Returns the value of the dot product between *m* and *n*.
// Row number of *m* must be equal to columns number of *n*.
func DotProduct(m, n mat.Matrix) mat.Matrix {
	rows, _ := m.Dims()
	_, columns := n.Dims()

	result := mat.NewDense(rows, columns, nil)
	result.Product(m, n)
	return result
}

// Return the value of a *fn* function (which takes as parameters the position row,
// column and his value) applied to a matrix *m*.
func ApplyFunction(fn func(i, j int, v float64) float64, m mat.Matrix) mat.Matrix {
	rows, columns := m.Dims()

	result := mat.NewDense(rows, columns, nil)
	result.Apply(fn, m)
	return result
}

// Returns the value of *m* mutliplied by *s* which is a float64.
func Scale(s float64, m mat.Matrix) mat.Matrix {
	rows, columns := m.Dims()

	result := mat.NewDense(rows, columns, nil)
	result.Scale(s, m)
	return result
}

// Returns the value of *m* multiplied by *n*.
// *m* and *n* must have the same dimension.
func Multiply(m, n mat.Matrix) mat.Matrix {
	rows, columns := m.Dims()

	result := mat.NewDense(rows, columns, nil)
	result.MulElem(m, n)
	return result
}