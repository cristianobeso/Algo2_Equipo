package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccionario(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return 1
		} else if elemento1 > elemento2 {
			return -1
		} else {
			return 0
		}
	})

	dic.Guardar(10, 1000)
	require.True(t, dic.Pertenece(10))
	require.Equal(t, 1000, dic.Obtener(10))
	require.Equal(t, 1, dic.Cantidad())

	dic.Guardar(5, 10)
	require.True(t, dic.Pertenece(5))
	require.Equal(t, 10, dic.Obtener(5))
	require.Equal(t, 2, dic.Cantidad())

	dic.Guardar(5, 500)
	require.True(t, dic.Pertenece(5))
	require.Equal(t, 500, dic.Obtener(5))
	require.Equal(t, 2, dic.Cantidad())

	dic.Guardar(15, 1500)
	require.True(t, dic.Pertenece(15))
	require.Equal(t, 1500, dic.Obtener(15))
	require.Equal(t, 3, dic.Cantidad())

	// lo recorre al reves, pero funciona
	iter := dic.Iterador()
	k, v := iter.VerActual()
	require.Equal(t, k, 15)
	require.Equal(t, v, 1500)

	iter.Siguiente()
	k, v = iter.VerActual()
	require.Equal(t, k, 10)
	require.Equal(t, v, 1000)

	iter.Siguiente()
	k, v = iter.VerActual()
	require.Equal(t, k, 5)
	require.Equal(t, v, 500)
	// o quizas lo guarda al reves ?

	dic.Iterar(func(clave, dato int) bool {
		fmt.Println(clave, dato)
		return true
	})
}
