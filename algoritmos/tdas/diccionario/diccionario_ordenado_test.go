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
			return -1
		} else if elemento1 > elemento2 {
			return 1
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

	iter := dic.Iterador()

	k, v := iter.VerActual()
	require.Equal(t, k, 5)
	require.Equal(t, v, 500)

	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 10)
	require.Equal(t, v, 1000)

	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 15)
	require.Equal(t, v, 1500)

	dic.Iterar(func(clave, dato int) bool {
		fmt.Println(clave, dato)
		return true
	})
}

func TestDiccionarioOrden(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	dic.Guardar(20, 123)

	dic.Guardar(10, 456)
	dic.Guardar(15, 789)
	dic.Guardar(5, 159)
	dic.Guardar(3, 753)
	dic.Guardar(6, 486)

	dic.Guardar(30, 426)
	dic.Guardar(40, 423)
	dic.Guardar(35, 126)
	dic.Guardar(22, 781)

	iter := dic.Iterador()
	k, v := iter.VerActual()
	require.Equal(t, k, 3)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 5)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 6)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 10)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 15)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 20)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 22)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 30)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 35)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 40)

	require.Equal(t, 423, v)
	// La verdad estoy sorprendido de que hay pasado la prueba
	// Sigo sintiendo que hay algo mal
	// Hare mas pruebas
}
