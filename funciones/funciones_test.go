package funciones_test

import (
	"testing"

	"github.com/aldogayaladh/go-test-1598/funciones"

	"github.com/stretchr/testify/assert"
)

// Se suben los tres casos que vimos
// Ustedes eligen cual es mas comodo para trabajar

// Test unitario simple
func TestSimpleSuma(t *testing.T) {
	// Arrange.
	assert := assert.New(t)
	var a, b int
	a = 2
	b = 3
	expectedResult := 5

	// Act.
	result := funciones.Suma(a, b)

	// Assert.
	assert.Equal(expectedResult, result)

}

// Uso de escenarios
func TestEscenariosSuna(t *testing.T) {

	t.Run("Esperamos que el resultado sea 5", func(t *testing.T) {
		// Arrange.
		assert := assert.New(t)
		var a, b int
		a = 2
		b = 3
		expectedResult := 5

		// Act.
		result := funciones.Suma(a, b)

		// Assert.
		assert.Equal(expectedResult, result)
	})

	t.Run("Esperamos que el resultado sea 8", func(t *testing.T) {
		// Arrange.
		assert := assert.New(t)
		var a, b int
		a = 5
		b = 3
		expectedResult := 8

		// Act.
		result := funciones.Suma(a, b)

		// Assert.
		assert.Equal(expectedResult, result)
	})

}

// Uso de Test Table
func TestTableSuma(t *testing.T) {

	// Arrange.
	assert := assert.New(t)

	type args struct {
		a int
		b int
	}

	// Creamos el test table.
	tests := []struct {
		name           string
		args           args
		expectedResult int
	}{
		// Casos a evaluar
		{
			name:           "Evaluamos cuando el resultado es 8",
			args:           args{a: 4, b: 4},
			expectedResult: 8,
		},
		{
			name:           "Evaluamos cuando el resultado es 5",
			args:           args{a: 2, b: 3},
			expectedResult: 5,
		},
		{
			name:           "Evaluamos cuando el resultado es 10",
			args:           args{a: 5, b: 5},
			expectedResult: 10,
		},
		{
			name:           "Evaluamos cuando el resultado es 10",
			args:           args{a: 100, b: 5},
			expectedResult: 105,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Act.
			result := funciones.Suma(test.args.a, test.args.b)

			// Assert.
			assert.Equal(test.expectedResult, result)
		})
	}

}
