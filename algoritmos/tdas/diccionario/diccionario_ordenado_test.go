package diccionario_test

import (
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

func TestIterarRangoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	iter := dic.IteradorRango(nil, nil)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.False(t, iter.HaySiguiente())
}

func TestIteradorVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})
	iter := dic.Iterador()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorRangoVacio(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})
	iter := dic.IteradorRango(nil, nil)
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
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

	// Definir el rango de claves
	claveInicio := 14
	claveFin := 39

	// Crear el iterador de rango
	iter := dic.IteradorRango(&claveInicio, &claveFin)

	// Definir el recorrido esperado
	recorridoEsperado := []struct {
		clave int
		valor int
	}{
		{15, 789},
		{20, 123},
		{22, 781},
		{30, 426},
		{33, 481},
		{35, 126},
	}

	for _, esperado := range recorridoEsperado {
		require.True(t, iter.HaySiguiente()) // Verificar que hay siguiente
		k, v := iter.VerActual()             // Verificar clave y valor actuales
		require.Equal(t, esperado.clave, k)
		require.Equal(t, esperado.valor, v)
		iter.Siguiente()
	}

	require.False(t, iter.HaySiguiente())
}

func TestIteradorOrdenado(t *testing.T) {
	dic := TDADiccionario.CrearABB[int, int](func(elemento1, elemento2 int) int {
		if elemento1 < elemento2 {
			return -1
		} else if elemento1 > elemento2 {
			return 1
		} else {
			return 0
		}
	})

	dic.Guardar(1, 10)
	dic.Guardar(2, 20)
	dic.Guardar(3, 30)
	dic.Guardar(4, 40)
	dic.Guardar(5, 50)
	dic.Guardar(6, 60)
	dic.Guardar(7, 70)
	iter := dic.Iterador()

	// Definir el recorrido esperado
	recorridoEsperado := []struct {
		clave int
		valor int
	}{
		{1, 10},
		{2, 20},
		{3, 30},
		{4, 40},
		{5, 50},
		{6, 60},
		{7, 70},
	}

	for _, esperado := range recorridoEsperado {
		require.True(t, iter.HaySiguiente()) // Verificar que hay siguiente
		k, v := iter.VerActual()             // Verificar clave y valor actuales
		require.Equal(t, esperado.clave, k)
		require.Equal(t, esperado.valor, v)
		iter.Siguiente()
	}

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
	require.Equal(t, 456, dic.Borrar(10))

	require.Equal(t, dic.Cantidad(), 2)
	require.False(t, dic.Pertenece(10))
	require.True(t, dic.Pertenece(20))
	dic.Borrar(20)
	require.Equal(t, dic.Cantidad(), 1)
	require.False(t, dic.Pertenece(20))
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
