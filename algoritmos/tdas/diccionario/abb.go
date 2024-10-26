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
	dic    *abb[K, V]
	actual *nodoAbb[K, V]
	pila   pila.Pila[*nodoAbb[K, V]]
}

type iterDicAbbRango[K comparable, V any] struct {
	dic    *abb[K, V]
	actual *nodoAbb[K, V]
	pila   pila.Pila[*nodoAbb[K, V]]
	desde  *K
	hasta  *K
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
Precondiciones: nodo debe ser un puntero a un nodo válido, clave debe ser de tipo K, y funcCmp debe ser una función de comparación válida. pila debe ser una pila de nodos.
Postcondiciones: Los nodos en la ruta hacia la clave se apilan en pila.
*/
func apilarRecursivo[K comparable, V any](nodo *nodoAbb[K, V], clave K, funcCmp func(K, K) int, pila pila.Pila[*nodoAbb[K, V]]) {
	if nodo == nil {
		return
	}
	pila.Apilar(nodo)
	if nodo.clave == clave {
		return
	}
	if funcCmp(nodo.clave, clave) > 0 {
		apilarRecursivo(nodo.izquierdo, clave, funcCmp, pila)
	} else {
		apilarRecursivo(nodo.derecho, clave, funcCmp, pila)
	}
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

/*
Precondiciones: abb debe ser un puntero a un ABB válido y clave debe ser de tipo K.
Postcondiciones: Se elimina el nodo asociado a la clave. Si la clave no existe, se lanza un pánico. Se decrementa la cantidad de elementos.
*/
func (abb *abb[K, V]) Borrar(clave K) V {
	if abb.raiz == nil {
		panic("La clave no pertenece al diccionario")
	}

	var borrado *nodoAbb[K, V]
	abb.raiz, borrado = borrarRecursivo(abb.raiz, clave, abb.funcCmp)
	if borrado == nil {
		panic("La clave no pertenece al diccionario")
	}

	abb.cantidad--
	return borrado.dato
}

/*
Precondiciones: nodo debe ser un puntero a un nodo válido y funcCmp debe ser una función de comparación válida.
Postcondiciones: Se devuelve el nuevo nodo tras la eliminación y el nodo eliminado.
*/
func borrarRecursivo[K comparable, V any](nodo *nodoAbb[K, V], clave K, funcCmp func(K, K) int) (*nodoAbb[K, V], *nodoAbb[K, V]) {
	if nodo == nil {
		return nil, nil
	}

	cmp := funcCmp(clave, nodo.clave)
	if cmp < 0 {
		var borrado *nodoAbb[K, V]
		nodo.izquierdo, borrado = borrarRecursivo(nodo.izquierdo, clave, funcCmp)
		return nodo, borrado
	} else if cmp > 0 {
		var borrado *nodoAbb[K, V]
		nodo.derecho, borrado = borrarRecursivo(nodo.derecho, clave, funcCmp)
		return nodo, borrado
	} else {
		// nodo sin hijos
		if nodo.izquierdo == nil && nodo.derecho == nil {
			return nil, nodo
		}

		// nodo con un hijo
		if nodo.izquierdo == nil {
			return nodo.derecho, nodo
		}
		if nodo.derecho == nil {
			return nodo.izquierdo, nodo
		}

		// nodo con dos hijos, sucesor inorder
		sucesor := buscarMin(nodo.derecho)
		nodo.clave = sucesor.clave
		nodo.dato = sucesor.dato
		nodo.derecho, _ = borrarRecursivo(nodo.derecho, sucesor.clave, funcCmp)
		return nodo, sucesor
	}
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
	return &iterDicAbb[K, V]{dic: abb, actual: nodoActual, pila: nuevaPila}
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
			iter.actual = nodoActual
		}
	}
}

/*
Precondiciones:
Postcondiciones:
*/
func (iter *iterDicAbb[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

/*
Precondiciones: iter debe ser un puntero a un iterador válido y no debe estar vacío.
Postcondiciones: Se devuelve la clave y el dato del nodo actual en el iterador.
*/
func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if desde != nil {
		apilarRecursivo(abb.raiz, *desde, abb.funcCmp, pila)
	} else {
		var nodo *nodoAbb[K, V]
		nodo = abb.raiz
		if nodo != nil {
			pila.Apilar(abb.raiz)
			for nodo.izquierdo != nil {
				nodo = nodo.izquierdo
				pila.Apilar(nodo)
			}
		}
	}

	for !pila.EstaVacia() {
		nodoActual := pila.Desapilar()
		if hasta != nil {
			if abb.funcCmp(*hasta, nodoActual.clave) < 0 {
				return
			}
		}
		if desde != nil {
			if abb.funcCmp(nodoActual.clave, *desde) >= 0 {
				visitar(nodoActual.clave, nodoActual.dato)
				if nodoActual.derecho != nil {
					pila.Apilar(nodoActual.derecho)
					//quizas se pueda usar la variable nodoActual...
					proximoNodo := nodoActual.derecho
					for proximoNodo.izquierdo != nil {
						proximoNodo = proximoNodo.izquierdo
						pila.Apilar(proximoNodo)
					}
				}
			}
		} else {
			visitar(nodoActual.clave, nodoActual.dato) // cambiar despues
			if nodoActual.derecho != nil {
				pila.Apilar(nodoActual.derecho)
				//quizas se pueda usar la variable nodoActual...
				proximoNodo := nodoActual.derecho
				for proximoNodo.izquierdo != nil {
					proximoNodo = proximoNodo.izquierdo
					pila.Apilar(proximoNodo)
				}
			}
		}

	}
}

/*
Precondiciones: abb debe ser un puntero a un ABB válido, y desde y hasta deben ser punteros a claves de tipo K. visitar debe ser una función válida.
Postcondiciones: Se aplica la función visitar a cada clave y dato en el rango especificado.
*/
func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	if desde != nil {
		apilarRecursivo(abb.raiz, *desde, abb.funcCmp, pila)
	} else {
		var nodo *nodoAbb[K, V]
		nodo = abb.raiz
		if nodo != nil {
			pila.Apilar(abb.raiz)
			for nodo.izquierdo != nil {
				nodo = nodo.izquierdo
				pila.Apilar(nodo)
			}
		}
	}
	return &iterDicAbbRango[K, V]{dic: abb, actual: pila.VerTope(), pila: pila, desde: desde, hasta: hasta}
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
		return nodo.derecho != nil
	} else {
		if iter.hasta != nil {
			if iter.dic.funcCmp(*iter.hasta, iter.pila.VerTope().clave) < 0 {
				iter.pila.Apilar(nodo)
				return false
			}
		}
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
		nodoActual := iter.pila.Desapilar()
		if iter.desde != nil {
			if iter.dic.funcCmp(nodoActual.clave, *iter.desde) >= 0 {
				if nodoActual.derecho != nil {
					iter.pila.Apilar(nodoActual.derecho)

					proximoNodo := nodoActual.derecho
					for proximoNodo.izquierdo != nil {
						proximoNodo = proximoNodo.izquierdo
						iter.pila.Apilar(proximoNodo)
					}
				}
			}
			for iter.dic.funcCmp(iter.pila.VerTope().clave, *iter.desde) < 0 {
				iter.pila.Desapilar()
			}
		} else {
			if nodoActual.derecho != nil {
				iter.pila.Apilar(nodoActual.derecho)

				proximoNodo := nodoActual.derecho
				for proximoNodo.izquierdo != nil {
					proximoNodo = proximoNodo.izquierdo
					iter.pila.Apilar(proximoNodo)
				}
			}
		}
	}
}

/*
Precondiciones: El iterador debe estar inicializado y no debe estar vacío. Debe haber al menos un elemento en la pila que el iterador está utilizando.
Postcondiciones: Se devuelve la clave y el valor del elemento actual en el iterador. No se produce modificación en el estado del iterador.
*/
func (iter *iterDicAbbRango[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}
