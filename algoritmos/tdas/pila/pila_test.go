package pila_test

import (
	TDAPila "tdas/pila"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestPilaVacia(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.True(t, pila.EstaVacia())

	pila2 := TDAPila.CrearPilaDinamica[bool]()
	require.True(t, pila2.EstaVacia())

	pila.Apilar(30)
	require.EqualValues(t, pila.VerTope(), 30)
	pila.Apilar(50)
	require.EqualValues(t, pila.VerTope(), 50)

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.Desapilar() })
}

func TestVerTope(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	pila.Apilar(10)
	pila.Apilar(30)
	pila.Apilar(20)
	pila.Apilar(5)
	require.EqualValues(t, pila.VerTope(), 5)
	pila.Apilar(100)
	pila.Apilar(150)
	require.EqualValues(t, pila.VerTope(), 150)
	removed := pila.Desapilar()
	require.EqualValues(t, removed, 150)
	require.EqualValues(t, pila.VerTope(), 100)

	pila2 := TDAPila.CrearPilaDinamica[bool]()
	pila2.Apilar(true)
	require.EqualValues(t, pila2.VerTope(), true)
	pila2.Apilar(false)
	require.EqualValues(t, pila2.VerTope(), false)
	pila2.Apilar(true)
	pila2.Apilar(true)
	pila2.Apilar(false)
	require.EqualValues(t, pila2.VerTope(), false)
	require.EqualValues(t, pila2.Desapilar(), false)
	require.EqualValues(t, pila2.VerTope(), true)
}

func TestApilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	for i := 0; i < 1000; i++ {
		pila.Apilar(i * 2)
		require.EqualValues(t, pila.VerTope(), i*2)
	}
	for i := 999; i >= 0; i-- {
		require.EqualValues(t, pila.VerTope(), i*2)
		removed := pila.Desapilar()
		require.EqualValues(t, removed, i*2)
	}
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila2 := TDAPila.CrearPilaDinamica[float64]()
	for i := 0; i < 100; i++ {
		pila2.Apilar(0.5 + float64(i))
		require.EqualValues(t, pila2.VerTope(), 0.5+float64(i))
	}
	for i := 99; i >= 0; i-- {
		require.EqualValues(t, pila2.Desapilar(), 0.5+float64(i))
	}

}

func TestDesapilar(t *testing.T) {
	pila := TDAPila.CrearPilaDinamica[int]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila.Apilar(20)
	require.EqualValues(t, pila.Desapilar(), 20)
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila.Desapilar() })

	pila2 := TDAPila.CrearPilaDinamica[float64]()
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.Desapilar() })

	pila2.Apilar(10.09)
	pila2.Apilar(2.35)
	require.EqualValues(t, pila2.Desapilar(), 2.35)
	require.EqualValues(t, pila2.Desapilar(), 10.09)

	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.VerTope() })
	require.PanicsWithValue(t, "La pila esta vacia", func() { pila2.Desapilar() })
}
