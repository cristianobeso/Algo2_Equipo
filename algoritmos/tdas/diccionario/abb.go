package diccionario

import (
	"tdas/pila"
)

type nodoAbb[K comparable, V any] struct {
	izquierdo *nodoAbb[K, V]
	derecho   *nodoAbb[K, V]
	clave     K
	dato      V
	altura    int
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

// Hay que hacer todas las funciones otra vez

func CrearABB[K comparable, V any](function_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: function_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato, altura: 1}
}

func mayor(num1, num2 int) int {
	if num1 > num2 {
		return num1
	}
	return num2
}

// actualiza hasta el padre del ultimo hijo insertado
func (nodo *nodoAbb[K, V]) actualizarAltura(clave K, funcCmp func(K, K) int) {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarRecursivo(nodo, clave, funcCmp, pila)

	var alturaIzq, alturaDer int
	for !pila.EstaVacia() {
		nodoActual := pila.Desapilar()
		if nodoActual.izquierdo == nil {
			alturaIzq = 0
		} else {
			alturaIzq = nodoActual.izquierdo.altura
		}
		if nodoActual.derecho == nil {
			alturaDer = 0
		} else {
			alturaDer = nodoActual.derecho.altura
		}
		(*nodoActual).altura = mayor(alturaIzq, alturaDer) + 1
	}
}

func apilarRecursivo[K comparable, V any](nodo *nodoAbb[K, V], clave K, funcCmp func(K, K) int, pila pila.Pila[*nodoAbb[K, V]]) {
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

func (abb *abb[K, V]) Altura() int {
	return abb.raiz.altura
}

// Retorna el nodo padre y en que lugar insertar un hijo
func (nodo *nodoAbb[K, V]) ubicar(clave K, ptrPadre *nodoAbb[K, V], num int, funcCmp func(K, K) int) (*nodoAbb[K, V], int) {
	if nodo == nil {
		return ptrPadre, num
	}
	if nodo.clave == clave {
		return nodo, 0
	}
	if funcCmp(nodo.clave, clave) > 0 { // Se supone que la comparacion devuelve un valor positivo si el primero es mas grande
		return nodo.izquierdo.ubicar(clave, nodo, -1, funcCmp)
	} else {
		return nodo.derecho.ubicar(clave, nodo, 1, funcCmp)
	}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	var necesitaActualizar bool

	nuevoNodo := crearNodo(clave, dato)
	if abb.raiz == nil {
		abb.raiz = nuevoNodo
		abb.cantidad++
	}
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)

	if dir == 0 {
		ptr.dato = dato
	} else {
		necesitaActualizar = (ptr.izquierdo == nil && ptr.derecho == nil) // si no tiene hijos se actualizan los datos de los padres sucesivos
		//no sin antes insertar el hijo
		if dir == -1 {
			ptr.izquierdo = nuevoNodo
		} else if dir == 1 {
			ptr.derecho = nuevoNodo
		}
		abb.cantidad++
		if necesitaActualizar {
			abb.raiz.actualizarAltura(ptr.clave, abb.funcCmp)
		}
	}

	// despues habria que crear una funcion para rotar
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)
	return ptr != nil && dir == 0 // Si al ubicarlo el ptr no es nulo y la direccion es 0
	// significa que ese puntero es el que tiene la clave buscada
}

func (abb *abb[K, V]) Obtener(clave K) V {
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)

	if ptr != nil && dir == 0 {
		return (*ptr).dato
	}
	panic("NO encontrado")
}

func (abb *abb[K, V]) Borrar(clave K) V {
	return abb.raiz.dato // por ahora
}

func (abb *abb[K, V]) Cantidad() int {
	return abb.cantidad
}

func (abb *abb[K, V]) Iterar(f func(clave K, dato V) bool) {
	if abb == nil {
		return
	}
	abb.raiz.izquierdo.iterar(f)
	f(abb.raiz.clave, abb.raiz.dato)
	abb.raiz.derecho.iterar(f)
}

func (nodo *nodoAbb[K, V]) iterar(f func(clave K, dato V) bool) {
	if nodo == nil {
		return
	}
	nodo.izquierdo.iterar(f)
	f(nodo.clave, nodo.dato)
	nodo.derecho.iterar(f)
}

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

func (iter *iterDicAbb[K, V]) HaySiguiente() bool {
	return !iter.pila.EstaVacia()
}

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
	}
}

func (iter *iterDicAbb[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarRecursivo(abb.raiz, *desde, abb.funcCmp, pila)
	for !pila.EstaVacia() {
		nodoActual := pila.Desapilar()
		if abb.funcCmp(*hasta, nodoActual.clave) < 0 {
			return
		}
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
	}
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
	pila := pila.CrearPilaDinamica[*nodoAbb[K, V]]()
	apilarRecursivo(abb.raiz, *desde, abb.funcCmp, pila)
	return &iterDicAbbRango[K, V]{dic: abb, actual: pila.VerTope(), pila: pila, desde: desde, hasta: hasta}
}

func (iter *iterDicAbbRango[K, V]) HaySiguiente() bool {
	return (!iter.pila.EstaVacia())
}

func (iter *iterDicAbbRango[K, V]) Siguiente() {
	if iter.HaySiguiente() {
		nodoActual := iter.pila.Desapilar()

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
		if iter.dic.funcCmp(iter.pila.VerTope().clave, *iter.desde) < 0 { // si el nuevo actual es menor que la clave buscada lo desapilo
			iter.pila.Desapilar()
		}
	}
}

func (iter *iterDicAbbRango[K, V]) VerActual() (K, V) {
	return iter.pila.VerTope().clave, iter.pila.VerTope().dato
}
