package pila

/* Definición del struct pila proporcionado por la cátedra. */

type pilaDinamica[T any] struct {
	datos    []T
	cantidad int
}

func CrearPilaDinamica[T any]() Pila[T] {
	return &pilaDinamica[T]{datos: make([]T, 1), cantidad: 0}
}

func (pila *pilaDinamica[T]) EstaVacia() bool {
	return pila.cantidad == 0
}

func (pila *pilaDinamica[T]) VerTope() T {
	if !pila.EstaVacia() {
		return pila.datos[pila.cantidad-1]
	}
	panic("La pila esta vacia")
}

func (pila *pilaDinamica[T]) Apilar(elemento T) {
	if cap(pila.datos) == pila.cantidad {
		pila.redimensionar(cap(pila.datos) * 2)
	}
	pila.datos[pila.cantidad] = elemento
	pila.cantidad++
}

func (pila *pilaDinamica[T]) Desapilar() T {
	if !pila.EstaVacia() {
		if pila.cantidad*4 <= cap(pila.datos) {
			pila.redimensionar(cap(pila.datos) / 2)
		}
		pila.cantidad--
		return pila.datos[pila.cantidad]
	}
	panic("La pila esta vacia")
}

func (pila *pilaDinamica[T]) redimensionar(tam int) {
	nuevaPila := make([]T, tam)
	copy(nuevaPila, pila.datos)
	pila.datos = nuevaPila
}
