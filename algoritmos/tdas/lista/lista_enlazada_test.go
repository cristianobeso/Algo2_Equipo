package lista_test

import (
	"testing"

	"tdas/lista"

	"github.com/stretchr/testify/require"
)

func TestListaVacia(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	require.True(t, l.EstaVacia())
	require.Equal(t, 0, l.Largo())
}

func TestInsertarPrimero(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(10)
	require.False(t, l.EstaVacia())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, 10, l.VerPrimero())
	require.Equal(t, 10, l.VerUltimo())
}

func TestInsertarUltimo(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(20)
	require.False(t, l.EstaVacia())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, 20, l.VerPrimero())
	require.Equal(t, 20, l.VerUltimo())
}

func TestBorrarPrimero(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarPrimero(30)
	l.InsertarPrimero(40)
	require.Equal(t, 40, l.BorrarPrimero())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, 30, l.VerPrimero())
}

func TestBorrarListaVaciaPanico(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	require.PanicsWithValue(t, "La lista esta vacia", func() { l.BorrarPrimero() })
}

func TestIterar(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)

	var suma int
	l.Iterar(func(dato int) bool {
		suma += dato
		return true
	})
	require.Equal(t, 6, suma)
}

func TestIteradorInsertarAlInicio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	iter := l.Iterador()

	iter.Insertar(1)
	require.Equal(t, 1, l.VerPrimero())
	require.Equal(t, 3, l.Largo())
}

func TestIteradorInsertarAlFinal(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	iter := l.Iterador()

	for iter.HaySiguiente() {
		iter.Siguiente()
	}
	iter.Insertar(3)

	require.Equal(t, 3, l.VerUltimo())
	require.Equal(t, 3, l.Largo())
}

func TestIteradorInsertarEnMedio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(3)
	iter := l.Iterador()

	iter.Siguiente() // Apunta al segundo elemento
	iter.Insertar(2)

	require.Equal(t, 3, l.Largo())

	iter = l.Iterador()
	require.Equal(t, 1, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 2, iter.VerActual())
	iter.Siguiente()
	require.Equal(t, 3, iter.VerActual())
}

func TestIteradorBorrarPrimero(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	iter := l.Iterador()

	require.Equal(t, 1, iter.Borrar())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, 2, l.VerPrimero())
}

func TestIteradorBorrarUltimo(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	iter := l.Iterador()

	iter.Siguiente() // Apunta al segundo elemento
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 1, l.Largo())
	require.Equal(t, 1, l.VerUltimo())
}

func TestIteradorBorrarEnMedio(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	l.InsertarUltimo(2)
	l.InsertarUltimo(3)
	iter := l.Iterador()

	iter.Siguiente() // Apunta al segundo elemento
	require.Equal(t, 2, iter.Borrar())
	require.Equal(t, 2, l.Largo())

	iter = l.Iterador()
	iter.Siguiente()
	require.Equal(t, 3, iter.VerActual())
}

func TestIteradorVerActualPanico(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	iter := l.Iterador()

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.VerActual() })
}

func TestIteradorSiguientePanico(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	iter := l.Iterador()

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Siguiente() })
}

func TestIteradorBorrarPanico(t *testing.T) {
	l := lista.CrearListaEnlazada[int]()
	l.InsertarUltimo(1)
	iter := l.Iterador()

	iter.Siguiente()
	require.PanicsWithValue(t, "El iterador termino de iterar", func() { iter.Borrar() })
}
