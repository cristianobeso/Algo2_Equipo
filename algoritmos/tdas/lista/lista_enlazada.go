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

type iterListaEnlazada[T any] struct {
	anterior *nodoLista[T]
	actual   *nodoLista[T]
	lista    *listaEnlazada[T]
	//Tengo en mis apuntes que el iterador tenia mas o menos esta estructura
	//Pero me quedan dudas porque para las primitivas no use 'lista'
}

func crearNodo[T any](dato T, ptr *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: ptr}
}

func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

/**************  Primitivas de la lista  **************/

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

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	//Este no tengo idea te lo dejo todo a vos
}

/**************  Fin lista  **************/

/**************  Primitivas del iterador  **************/
//No estoy seguro si esto funcione, pero es algo

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{anterior: nil, actual: nil, lista: lista}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual.siguiente != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	ptr := iterador.actual.siguiente
	iterador.anterior = iterador.actual
	iterador.actual = ptr
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	ptr := iterador.actual.siguiente
	nuevoNodo := crearNodo(elemento, ptr)
	iterador.actual = nuevoNodo
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	elemento := iterador.actual.dato
	ptr := iterador.actual.siguiente
	iterador.anterior.siguiente = ptr
	iterador.actual = ptr
	return elemento
}

/**************  Fin iterador  **************/
