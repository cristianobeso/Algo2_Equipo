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
	nuevoNodo := crearNodo(elemento, lista.primero)
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
		panic("La lista esta vacia")
	}
	elemento := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.primero == nil {
		lista.ultimo = nil
	}
	lista.largo--
	return elemento
}

func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

func (lista *listaEnlazada[T]) Iterar(visitar func(T) bool) {
	var iterarRecursivo func(*nodoLista[T])
	iterarRecursivo = func(actual *nodoLista[T]) {
		if actual == nil {
			return
		}
		if !visitar(actual.dato) {
			return
		}
		iterarRecursivo(actual.siguiente)
	}
	iterarRecursivo(lista.primero)
}

/**************  Fin lista  **************/

/**************  Primitivas del iterador  **************/

func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, lista: lista}
}

func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

func (iterador *iterListaEnlazada[T]) Insertar(elemento T) {
	nuevoNodo := crearNodo(elemento, iterador.actual)
	if iterador.anterior == nil {
		iterador.lista.primero = nuevoNodo
	} else {
		iterador.anterior.siguiente = nuevoNodo
	}
	if !iterador.HaySiguiente() {
		iterador.lista.ultimo = nuevoNodo
	}
	iterador.actual = nuevoNodo
	iterador.lista.largo++
}

func (iterador *iterListaEnlazada[T]) Borrar() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	elemento := iterador.actual.dato
	if iterador.anterior != nil {
		iterador.anterior.siguiente = iterador.actual.siguiente
	} else {
		iterador.lista.primero = iterador.actual.siguiente
	}
	if iterador.lista.ultimo == iterador.actual {
		iterador.lista.ultimo = iterador.anterior
	}

	siguiente := iterador.actual.siguiente
	iterador.lista.largo--
	iterador.actual = siguiente
	return elemento
}

/**************  Fin iterador  **************/
