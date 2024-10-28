package diccionario

import (
	"tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
}

type abb[K comparable, V any] struct {
	raiz     *nodoAbb[K, V]
	cantidad int
	funcCmp  func(K, K) int
}

type iterDicAbb[K comparable, V any] struct {
	dic  *abb[K, V]
	pila pila.Pila[*nodoAbb[K, V]]
}

type iterDicAbbRango[K comparable, V any] struct {
	dic   *abb[K, V]
	pila  pila.Pila[*nodoAbb[K, V]]
	desde *K
	hasta *K
}

/*
Precondiciones: function_cmp debe ser una función válida que compare dos claves de tipo K y devuelva un entero.
Postcondiciones: Se devuelve un puntero a un nuevo diccionario ordenado vacío.
*/
func CrearABB[K comparable, V any](function_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: function_cmp}
}

/*
Precondiciones: clave y dato deben ser de los tipos correspondientes K y V.
Postcondiciones: Se devuelve un puntero a un nuevo nodo de tipo nodoAbb con la clave y el dato proporcionados, y altura inicial de 1.
*/
func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
}

/*
Precondiciones: nodo no debe ser nulo y funcCmp debe ser una función de comparación válida.
Postcondiciones: Se devuelve el nodo padre donde debería insertarse la clave y una dirección indicando si se debe ir a la izquierda, derecha o si la clave ya existe.
*/
// Retorna el nodo padre y en que lugar insertar un hijo
func (nodo *nodoAbb[K, V]) ubicar(clave K, ptrPadre *nodoAbb[K, V], funcCmp func(K, K) int) *nodoAbb[K, V] {
	if nodo == nil {
		return ptrPadre
	}
	if funcCmp(nodo.clave, clave) == 0 {
		return nodo
	}
	if funcCmp(nodo.clave, clave) > 0 { // Se supone que la comparacion devuelve un valor positivo si el primero es mas grande
		return nodo.izquierdo.ubicar(clave, nodo, funcCmp)
	} else {
		return nodo.derecho.ubicar(clave, nodo, funcCmp)
	}
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido, clave debe ser de tipo K, y dato debe ser de tipo V.
Postcondiciones: La clave y el dato se insertan en el árbol. Si la clave ya existe, se actualiza el dato. Se incrementa la cantidad de elementos.
*/
func (abb *abb[K, V]) Guardar(clave K, dato V) {

	nuevoNodo := crearNodo(clave, dato)
	if abb.raiz == nil {
		abb.raiz = nuevoNodo
		abb.cantidad++
	} else {
		ptr := abb.raiz.ubicar(clave, abb.raiz, abb.funcCmp)
		if abb.funcCmp(ptr.clave, clave) == 0 {
			ptr.dato = dato
		} else {
			if abb.funcCmp(ptr.clave, clave) > 0 {
				ptr.izquierdo = nuevoNodo
			} else {
				ptr.derecho = nuevoNodo
			}
			abb.cantidad++
		}
	}
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido y clave debe ser de tipo K.
Postcondiciones: Se devuelve true si la clave está en el árbol; de lo contrario, devuelve false.
*/
func (abb *abb[K, V]) Pertenece(clave K) bool {
	ptr := abb.raiz.ubicar(clave, abb.raiz, abb.funcCmp)
	if ptr == nil {
		return false
	}
	return abb.funcCmp(ptr.clave, clave) == 0
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido y clave debe ser de tipo K.
Postcondiciones: Se devuelve el dato asociado a la clave si existe. Si la clave no está, se lanza un pánico.
*/
func (abb *abb[K, V]) Obtener(clave K) V {
	ptr := abb.raiz.ubicar(clave, abb.raiz, abb.funcCmp)

	if ptr != nil {
		if abb.funcCmp(ptr.clave, clave) == 0 {
			return ptr.dato
		}
	}
	panic("La clave no pertenece al diccionario")
}

/*       ***********        nuevas funciones para  borrar        ************       */
func (nodo *nodoAbb[K, V]) buscarPadre(ptrPadre *nodoAbb[K, V], clave K, funcCmp func(K, K) int) *nodoAbb[K, V] {
	if nodo == nil { // se supone que nunca sera nil
		return ptrPadre
	}
	if funcCmp(nodo.clave, clave) == 0 {
		return ptrPadre
	}
	if funcCmp(nodo.clave, clave) > 0 { // Se supone que la comparacion devuelve un valor positivo si el primero es mas grande
		return nodo.izquierdo.buscarPadre(nodo, clave, funcCmp)
	} else {
		return nodo.derecho.buscarPadre(nodo, clave, funcCmp)
	}
}

func (nodo *nodoAbb[K, V]) borrar(clave K, funcCmp func(K, K) int) V {
	ptrPadre := nodo.buscarPadre(nodo, clave, funcCmp)
	var ptrHijo *nodoAbb[K, V]
	if funcCmp(ptrPadre.clave, clave) > 0 { // si el padre es mayor a la clave del hijo a borrar, implica que este hijo esta a su izquierda
		ptrHijo = ptrPadre.izquierdo
	}
	if funcCmp(ptrPadre.clave, clave) < 0 {
		ptrHijo = ptrPadre.derecho
	}
	// sin hijo
	if ptrHijo.izquierdo == nil && ptrHijo.derecho == nil {
		if funcCmp(ptrPadre.clave, clave) > 0 {
			ptrPadre.izquierdo = nil
			return ptrHijo.dato
		}
		ptrPadre.derecho = nil
		return ptrHijo.dato
	}
	// un hijo
	if ptrHijo.izquierdo == nil && ptrHijo.derecho != nil {
		if funcCmp(ptrPadre.clave, clave) > 0 {
			ptrPadre.izquierdo = ptrHijo.derecho
			return ptrHijo.dato
		}
		ptrPadre.derecho = ptrHijo.derecho
		return ptrHijo.dato
	}

	if ptrHijo.izquierdo != nil && ptrHijo.derecho == nil {
		if funcCmp(ptrPadre.clave, clave) > 0 {
			ptrPadre.izquierdo = ptrHijo.izquierdo
			return ptrHijo.dato
		}
		ptrPadre.derecho = ptrHijo.izquierdo
		return ptrHijo.dato
	}
	// dos hijos
	sucesor := buscarMin(ptrHijo.derecho)
	ptrPadre.borrar(sucesor.clave, funcCmp) //borra al sucesor
	valorAux := ptrHijo.dato
	ptrHijo.clave = sucesor.clave
	ptrHijo.dato = sucesor.dato
	return valorAux
}

func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.Pertenece(clave) {
		ptrPadre := abb.raiz.buscarPadre(abb.raiz, clave, abb.funcCmp)
		abb.cantidad--
		if abb.funcCmp(ptrPadre.clave, clave) == 0 {

			if ptrPadre.izquierdo == nil && ptrPadre.derecho == nil {
				abb.raiz = nil
				return ptrPadre.dato
			}
			// un hijo
			if ptrPadre.izquierdo == nil && ptrPadre.derecho != nil {
				abb.raiz = ptrPadre.derecho
				return ptrPadre.dato
			}
			if ptrPadre.izquierdo != nil && ptrPadre.derecho == nil {
				abb.raiz = ptrPadre.izquierdo
				return ptrPadre.dato
			}
			sucesor := buscarMin(ptrPadre.derecho)
			ptrPadre.borrar(sucesor.clave, abb.funcCmp) //borra al sucesor
			valorAux := ptrPadre.dato
			ptrPadre.clave = sucesor.clave
			ptrPadre.dato = sucesor.dato
			return valorAux
		}
		return abb.raiz.borrar(clave, abb.funcCmp)
	}
	panic("La clave no pertenece al diccionario")
}

/*
Precondiciones: nodo debe ser un puntero a un nodo válido.
Postcondiciones: Se devuelve el nodo con la clave mínima en el subárbol.
*/
func buscarMin[K comparable, V any](nodo *nodoAbb[K, V]) *nodoAbb[K, V] {
	for nodo.izquierdo != nil {
		nodo = nodo.izquierdo
	}
	return nodo
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido.
Postcondiciones: Se devuelve la cantidad de elementos en el árbol.
*/
func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido y f debe ser una función válida que acepte una clave y un dato.
Postcondiciones: Se aplica la función f a cada clave y dato en el árbol en orden inorden.
*/
func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if abb == nil {
		return
	}
	abb.raiz.izquierdo.iterar(f)
	f(abb.raiz.clave, abb.raiz.dato)
	abb.raiz.derecho.iterar(f)
}

/*
Precondiciones: nodo debe ser un puntero a un nodo válido y f debe ser una función válida.
Postcondiciones: Se aplica la función f a cada clave y dato en el subárbol en orden inorden.
*/
func (nodo *nodoAbb[K, V]) iterar(f func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	nodo.izquierdo.iterar(f)
	f(nodo.clave, nodo.dato)
	nodo.derecho.iterar(f)
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido.
Postcondiciones: Se devuelve un nuevo iterador para el ABB.
*/
func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	nuevaPila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	var nodoActual *nodoAbb[K, V]
	nodoActual = abb.raiz
	if nodoActual != nil {
		nuevaPila.Apilar(abb.raiz)
		for nodoActual.izquierdo != nil {
			nodoActual = nodoActual.izquierdo
			nuevaPila.Apilar(nodoActual)
		}
	}
	return &iterDicAbb[K, V]{dic: abb, pila: nuevaPila}
}

/*
Precondiciones: iter debe ser un puntero a un iterador válido.
Postcondiciones: Se devuelve true si hay un siguiente elemento en la pila; de lo contrario, false.
*/
func (iter *iterDicAbb[K, V]) HaySiguiente() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	nodo := iter.pila.Desapilar()
	if iter.pila.EstaVacia() {
		iter.pila.Apilar(nodo)
		return nodo.derecho != nil
	} else {
		iter.pila.Apilar(nodo)
		return true
	}
}

/*
Precondiciones: iter debe ser un puntero a un iterador válido.
Postcondiciones: Se mueve al siguiente elemento en el iterador, actualizando la pila según sea necesario.
*/
func (iter *iterDicAbb[K, V]) Siguiente() {
	if iter.HaySiguiente() {
		nodoActual := iter.pila.Desapilar()
		if nodoActual.derecho != nil { // si existe un hijo derecho del desapilado lo apilo
			nodoActual = nodoActual.derecho
			iter.pila.Apilar(nodoActual)
			for nodoActual.izquierdo != nil { // y apilo a todos sus hijos izquierdos
				nodoActual = nodoActual.izquierdo
				iter.pila.Apilar(nodoActual)
			}
		}
	} else {
		panic("El iterador termino de iterar")
	}
}

/*
Precondiciones:
Postcondiciones:
*/
func (iter *iterDicAbb[K, V]) VerActual() (K, V) {
	if iter.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

/*
Precondiciones: iter debe ser un puntero a un iterador válido y no debe estar vacío.
Postcondiciones: Se devuelve la clave y el dato del nodo actual en el iterador.
*/
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	if abb == nil {
		return
	}

	if desde == nil && hasta == nil { // si no tiene limites se iterar con el iterador interno sin rangos
		abb.Iterar(visitar)
	} else {
		if desde != nil {
			if abb.funcCmp(*desde, abb.raiz.clave) < 0 {
				abb.raiz.izquierdo.iterarRango(visitar, abb.funcCmp, desde, hasta)
			}
		} else {
			abb.raiz.izquierdo.iterarRango(visitar, abb.funcCmp, desde, hasta)
		}

		if desde != nil && hasta == nil {
			if abb.funcCmp(*desde, abb.raiz.clave) < 0 {
				visitar(abb.raiz.clave, abb.raiz.dato)
			}
		} else if desde == nil && hasta != nil {
			if abb.funcCmp(abb.raiz.clave, *hasta) < 0 {
				visitar(abb.raiz.clave, abb.raiz.dato)
			}
		} else {
			if (abb.funcCmp(*desde, abb.raiz.clave) < 0) && (abb.funcCmp(abb.raiz.clave, *hasta) < 0) {
				visitar(abb.raiz.clave, abb.raiz.dato)
			}
		}

		if hasta != nil {
			if abb.funcCmp(abb.raiz.clave, *hasta) < 0 {
				abb.raiz.derecho.iterarRango(visitar, abb.funcCmp, desde, hasta)
			}
		} else {
			abb.raiz.derecho.iterarRango(visitar, abb.funcCmp, desde, hasta)
		}
	}
}

func (nodo *nodoAbb[K, V]) iterarRango(f func(clave K, dato V) bool, funcCmp func(clave1 K, clave2 K) int, desde *K, hasta *K) {
	if nodo == nil {
		return
	}

	if desde != nil {
		if funcCmp(*desde, nodo.clave) < 0 {
			nodo.izquierdo.iterarRango(f, funcCmp, desde, hasta)
		}
	} else {
		nodo.izquierdo.iterarRango(f, funcCmp, desde, hasta)
	}

	if desde != nil && hasta == nil {
		if funcCmp(*desde, nodo.clave) < 0 {
			f(nodo.clave, nodo.dato)
		}
	} else if desde == nil && hasta != nil {
		if funcCmp(nodo.clave, *hasta) < 0 {
			f(nodo.clave, nodo.dato)
		}
	} else {
		if (funcCmp(*desde, nodo.clave) < 0) && (funcCmp(nodo.clave, *hasta) < 0) {
			f(nodo.clave, nodo.dato)
		}
	}

	if hasta != nil {
		if funcCmp(nodo.clave, *hasta) < 0 {
			nodo.derecho.iterarRango(f, funcCmp, desde, hasta)
		}
	} else {
		nodo.derecho.iterarRango(f, funcCmp, desde, hasta)
	}

}

/*
Precondiciones: abb debe ser un puntero a un ABB válido, y desde y hasta deben ser punteros a claves de tipo K. visitar debe ser una función válida.
Postcondiciones: Se aplica la función visitar a cada clave y dato en el rango especificado.
*/
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	if desde == nil && hasta == nil {
		return abb.Iterador()
	}
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarRango(desde, hasta, abb.raiz, abb.funcCmp, pila)
	return &iterDicAbbRango[K, V]{dic: abb, pila: pila, desde: desde, hasta: hasta}
}

func apilarRango[K comparable, V any](desde *K, hasta *K, nodo *nodoAbb[K, V], funcCmp func(K, K) int, pila pila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}
	if desde != nil && hasta != nil {
		if funcCmp(*desde, nodo.clave) <= 0 && funcCmp(nodo.clave, *hasta) <= 0 { // si esta dentro del rango lo apilo
			pila.Apilar(nodo)
			apilarRango(desde, hasta, nodo.izquierdo, funcCmp, pila)
		} else if funcCmp(nodo.clave, *desde) < 0 {
			apilarRango(desde, hasta, nodo.derecho, funcCmp, pila)
		} else if funcCmp(*hasta, nodo.clave) < 0 {
			apilarRango(desde, hasta, nodo.izquierdo, funcCmp, pila)
		}
	} else if desde == nil {

		if funcCmp(nodo.clave, *hasta) <= 0 { // si esta dentro del rango lo apilo
			pila.Apilar(nodo)
			apilarRango(desde, hasta, nodo.izquierdo, funcCmp, pila)
		} else if funcCmp(*hasta, nodo.clave) < 0 {
			apilarRango(desde, hasta, nodo.izquierdo, funcCmp, pila)
		}
	} else if hasta == nil {
		if funcCmp(nodo.clave, *desde) <= 0 {
			pila.Apilar(nodo)
			apilarRango(desde, hasta, nodo.izquierdo, funcCmp, pila)
		} else if funcCmp(nodo.clave, *desde) < 0 {
			apilarRango(desde, hasta, nodo.derecho, funcCmp, pila)
		}
	}
}

/*
Precondiciones: iter debe ser un puntero a un iterador de rango válido.
Postcondiciones: Se devuelve true si hay un siguiente elemento en la pila; de lo contrario, false.
*/
func (iter *iterDicAbbRango[K, V]) HaySiguiente() bool {
	if iter.pila.EstaVacia() {
		return false
	}
	nodo := iter.pila.Desapilar()
	if iter.pila.EstaVacia() {
		iter.pila.Apilar(nodo)
		if nodo.derecho != nil {
			pilaAux := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
			apilarRango(iter.desde, iter.hasta, nodo.derecho, iter.dic.funcCmp, pilaAux)
			return !pilaAux.EstaVacia()
		}
		return nodo.derecho != nil
	} else {
		iter.pila.Apilar(nodo)
		return true
	}
}

/*
Precondiciones: iter debe ser un puntero a un iterador de rango válido.
Postcondiciones: Se mueve al siguiente elemento en el iterador de rango, actualizando la pila según sea necesario.
*/
func (iter *iterDicAbbRango[K, V]) Siguiente() {
	if iter.HaySiguiente() {
		elemento := iter.pila.Desapilar()
		apilarRango(iter.desde, iter.hasta, elemento.derecho, iter.dic.funcCmp, iter.pila)
	} else {
		panic("El iterador termino de iterar")
	}
}

/*
Precondiciones: El iterador debe estar inicializado y no debe estar vacío. Debe haber al menos un elemento en la pila que el iterador está utilizando.
Postcondiciones: Se devuelve la clave y el valor del elemento actual en el iterador. No se produce modificación en el estado del iterador.
*/
func (iter *iterDicAbbRango[K, V]) VerActual() (K, V) {
	if iter.pila.EstaVacia() {
		panic("El iterador termino de iterar")
	}
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}
