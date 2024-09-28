package lista

type nodoLista[T any] struct {
	dato      T
	siguiente *nodoLista[T]
}

type listaEnlazada[T any] struct {
	primero *nodoLista[T]
	ultimo  *nodoLista[T]
	largo   int
}

func crearNodo[T any](dato T, ptr *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: ptr}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}
func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	ptr := lista.primero
	nuevoNodo := crearNodo(elemento, ptr)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) InsertarUltimo(elemento T) {
	nuevoNodo := crearNodo(elemento, nil)
	if lista.EstaVacia() {
		lista.primero = nuevoNodo
	} else {
		lista.ultimo.siguiente = nuevoNodo
	}
	lista.ultimo = nuevoNodo
	lista.largo++
}

func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("Error")
	}
	elemento := lista.primero.dato
	lista.primero = lista.primero.siguiente
	lista.largo--
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("Error")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("Error")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	if lista.EstaVacia() {
		panic("Error")
	}
	return lista.largo
}
