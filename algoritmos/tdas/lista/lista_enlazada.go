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

/*
Precondiciones: Ninguna.
Postcondiciones: Crea y devuelve un nuevo nodo de tipo nodoLista[T] que contiene el valor dato y un puntero al siguiente nodo ptr.
El campo siguiente del nodo creado se inicializa con el valor de ptr.
*/
func crearNodo[T any](dato T, ptr *nodoLista[T]) *nodoLista[T] {
	return &nodoLista[T]{dato: dato, siguiente: ptr}
}

/*
Precondiciones: Ninguna.
Postcondiciones: Crea y devuelve una nueva lista enlazada vacía de tipo listaEnlazada[T], con el puntero primero y ultimo
inicializados en nil y el largo en 0.
*/
func CrearListaEnlazada[T any]() Lista[T] {
	return &listaEnlazada[T]{primero: nil, ultimo: nil, largo: 0}
}

/**************  Primitivas de la lista  **************/

/*
Precondiciones: Ninguna.
Postcondiciones: Devuelve true si la lista está vacía, es decir, si no contiene elementos. Devuelve false en caso contrario.
*/
func (lista *listaEnlazada[T]) EstaVacia() bool {
	return lista.primero == nil
}

/*
Precondiciones: Ninguna.
Postcondiciones: El elemento se inserta al inicio de la lista.
Si la lista estaba vacía, el elemento se convierte también en el último. El tamaño de la lista se incrementa en 1.
*/
func (lista *listaEnlazada[T]) InsertarPrimero(elemento T) {
	nuevoNodo := crearNodo(elemento, lista.primero)
	if lista.EstaVacia() {
		lista.ultimo = nuevoNodo
	}
	lista.primero = nuevoNodo
	lista.largo++
}

/*
Precondiciones: Ninguna.
Postcondiciones: El elemento se inserta al final de la lista.
Si la lista estaba vacía, el elemento se convierte también en el primero. El tamaño de la lista se incrementa en 1.
*/
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

/*
Precondiciones: La lista no debe estar vacía.
Postcondiciones: El primer elemento de la lista es eliminado y su valor es devuelto. Si después de eliminar el elemento,
la lista queda vacía, el puntero ultimo también se actualiza a nil. El tamaño de la lista se decrementa en 1.
Si la lista estaba vacía, entra en pánico con el mensaje: "La lista esta vacia".
*/
func (lista *listaEnlazada[T]) BorrarPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	elemento := lista.primero.dato
	lista.primero = lista.primero.siguiente
	if lista.EstaVacia() {
		lista.ultimo = nil
	}
	lista.largo--
	return elemento
}

/*
Precondiciones: La lista no debe estar vacía.
Postcondiciones: Devuelve el valor del primer elemento de la lista sin eliminarlo.
Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
*/
func (lista *listaEnlazada[T]) VerPrimero() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.primero.dato
}

/*
Precondiciones: La lista no debe estar vacía.
Postcondiciones: Devuelve el valor del último elemento de la lista sin eliminarlo.
Si la lista está vacía, entra en pánico con el mensaje: "La lista esta vacia".
*/
func (lista *listaEnlazada[T]) VerUltimo() T {
	if lista.EstaVacia() {
		panic("La lista esta vacia")
	}
	return lista.ultimo.dato
}

/*
Precondiciones: Ninguna.
Postcondiciones: Devuelve el número de elementos en la lista.
*/
func (lista *listaEnlazada[T]) Largo() int {
	return lista.largo
}

/*
Precondiciones: Ninguna.
Postcondiciones: Aplica la función visitar a cada elemento de la lista en orden. Si la función visitar devuelve false,
la iteración se detiene. De lo contrario, la iteración continúa hasta el final de la lista.
*/
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

/*
Precondiciones: Ninguna.
Postcondiciones: Devuelve un nuevo iterador que comienza en el primer elemento de la lista.
*/
func (lista *listaEnlazada[T]) Iterador() IteradorLista[T] {
	return &iterListaEnlazada[T]{actual: lista.primero, lista: lista}
}

/*
Precondiciones: El iterador debe tener un elemento actual disponible.
Postcondiciones: Devuelve el valor del elemento actual del iterador.
Si no hay más elementos, entra en pánico con el mensaje: "El iterador termino de iterar".
*/
func (iterador *iterListaEnlazada[T]) VerActual() T {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	return iterador.actual.dato
}

/*
Precondiciones: Ninguna.
Postcondiciones: Devuelve true si hay un elemento siguiente en la lista por iterar, o false si no lo hay.
*/
func (iterador *iterListaEnlazada[T]) HaySiguiente() bool {
	return iterador.actual != nil
}

/*
Precondiciones: El iterador debe estar apuntando a un elemento válido.
Postcondiciones: Avanza el iterador al siguiente elemento en la lista.
Si no hay más elementos, entra en pánico con el mensaje: "El iterador termino de iterar".
*/
func (iterador *iterListaEnlazada[T]) Siguiente() {
	if !iterador.HaySiguiente() {
		panic("El iterador termino de iterar")
	}
	iterador.anterior = iterador.actual
	iterador.actual = iterador.actual.siguiente
}

/*
Precondiciones: Ninguna.
Postcondiciones: Inserta un nuevo elemento en la lista en la posición actual del iterador. Si el iterador estaba al inicio de
la lista, el nuevo nodo se convierte en el primer nodo. Si estaba al final, el nuevo nodo se convierte en el último.
El tamaño de la lista se incrementa en 1.
*/
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

/*
Precondiciones: El iterador debe estar apuntando a un elemento válido.
Postcondiciones: Elimina el elemento actual del iterador y devuelve su valor. El iterador avanza al siguiente elemento después
de la eliminación. Si no hay más elementos, entra en pánico con el mensaje: "El iterador termino de iterar".
*/
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

	iterador.lista.largo--
	iterador.actual = iterador.actual.siguiente
	return elemento
}

/**************  Fin iterador  **************/
