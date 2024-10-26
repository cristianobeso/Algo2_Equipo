package diccionario_test

import (
	"fmt"
	TDADiccionario "tdas/diccionario"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDiccionarioAbbVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})
	require.Equal(t, 0, dic.Cantidad())
	iter := dic.Iterador()
	require.False(t, iter.HaySiguiente())
}

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

	require.False(t, iter.HaySiguiente())

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
}

func TestIterarRango(t *testing.T) {
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
	dic.Guardar(33, 481)

	claveInicio := 13
	claveFin := 36
	slice := []int{}
	dic.IterarRango(&claveInicio, &claveFin, func(clave, dato int) bool {
		slice = append(slice, clave)
		return true
	})
	require.Equal(t, slice[0], 15)
	require.Equal(t, slice[1], 20)
	require.Equal(t, slice[2], 22)
	require.Equal(t, slice[3], 30)
	require.Equal(t, slice[4], 33)
	require.Equal(t, slice[5], 35)
}

func TestIterarRangoSinRango(t *testing.T) {
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
	dic.Guardar(33, 481)

	slice := []int{}
	dic.IterarRango(nil, nil, func(clave, dato int) bool {
		slice = append(slice, clave)
		return true
	})
	require.Equal(t, slice[0], 3)
	require.Equal(t, slice[1], 5)
	require.Equal(t, slice[2], 6)
	require.Equal(t, slice[3], 10)
	require.Equal(t, slice[4], 15)
	require.Equal(t, slice[5], 20)
	require.Equal(t, slice[6], 22)
	require.Equal(t, slice[7], 30)
	require.Equal(t, slice[8], 33)
	require.Equal(t, slice[9], 35)
	require.Equal(t, slice[10], 40)
}

func TestIteradorRangoSinRango(t *testing.T) {
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
	dic.Guardar(33, 481)

	iter := dic.IteradorRango(nil, nil)
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
	require.Equal(t, k, 33)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 35)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 40)

	require.Equal(t, 423, v)

}

func TestIteradorRango(t *testing.T) {
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
	dic.Guardar(33, 481)

	claveInicio := 13
	claveFin := 36

	iter := dic.IteradorRango(&claveInicio, &claveFin)
	k, v := iter.VerActual()
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
	require.Equal(t, k, 33)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 35)

	require.Equal(t, v, 126)

	require.False(t, iter.HaySiguiente())
}

func TestEliminar(t *testing.T) {
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
	dic.Guardar(30, 789)

	require.Equal(t, dic.Cantidad(), 3)
	dic.Borrar(10)
	require.Equal(t, dic.Cantidad(), 2)
	require.False(t, dic.Pertenece(10))

	dic.Borrar(20)
	require.Equal(t, dic.Cantidad(), 1)
	require.False(t, dic.Pertenece(20))
}

func TestEliminarConDosHijos(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	dic.Guardar(20, 1)
	dic.Guardar(10, 2)
	dic.Guardar(30, 3)
	dic.Guardar(25, 4)
	dic.Guardar(22, 5)
	dic.Guardar(23, 6)
	dic.Guardar(35, 7)

	require.Equal(t, dic.Cantidad(), 7)
	dic.Borrar(30)
	require.Equal(t, dic.Cantidad(), 6)
	require.True(t, dic.Pertenece(25))
	require.False(t, dic.Pertenece(30))

	iter := dic.Iterador()
	k, v := iter.VerActual()
	require.Equal(t, k, 10)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 20)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 22)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 23)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 25)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 35)
	require.Equal(t, v, 7)

	require.False(t, iter.HaySiguiente())

	dic.Guardar(30, 3)
	require.True(t, dic.Pertenece(30))

	iterador := dic.Iterador()
	k, v = iterador.VerActual()
	require.Equal(t, k, 10)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 20)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 22)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 23)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 25)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 30)
	iter.Siguiente()

	k, v = iter.VerActual()
	require.Equal(t, k, 35)
	require.Equal(t, v, 7)

}

func TestBuscarEnArbolGrande(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	for i := 1; i <= 1000; i++ {
		dic.Guardar(i, i*100)
	}

	for i := 1; i <= 1000; i++ {
		require.True(t, dic.Pertenece(i))
		require.Equal(t, i*100, dic.Obtener(i))
	}

	require.Equal(t, 1000, dic.Cantidad())
}

func TestIterarGrandesCantidad(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	for i := 0; i < 1000; i++ {
		dic.Guardar(i, i*10)
	}

	count := 0
	dic.Iterar(func(clave, dato int) bool {
		count++
		return true
	})

	require.Equal(t, count, 1000)
}
