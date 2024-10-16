package diccionario

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

// Hay que hacer todas las funciones otra vez

func CrearABB[K comparable, V any](function_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: function_cmp}
}

func crearNodo[K comparable, V any](clave K, dato V) *nodoAbb[K, V] {
	return &nodoAbb[K, V]{izquierdo: nil, derecho: nil, clave: clave, dato: dato}
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
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)
	nuevoNodo := crearNodo(clave, dato)
	if ptr == nil {
		abb.raiz = nuevoNodo
	}
	if dir == -1 {
		(*ptr).izquierdo = nuevoNodo
	} else if dir == 1 {
		(*ptr).derecho = nuevoNodo
	} else {
		(*ptr).dato = dato
	}
	abb.cantidad++
	// despues habrai que crear una funcion para rotar
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	cont := 0
	abb.Iterar(func(claveB K, dato V) bool {
		if claveB == clave {
			return false
		}
		cont++
		return true
	})
	return cont+1 != abb.cantidad // Si el indice es distinto significa que se corto la iteracion por ende encontro el elemento
}

func (abb *abb[K, V]) PerteneceV2(clave K) bool {
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)
	return ptr != nil && dir == 0 // Si al ubicarlo el ptr no es nulo y la direccion es 0
	// significa que ese puntero es el que tiene la clave buscada
}

func (abb *abb[K, V]) ObtenerV2(clave K) V {
	ptr, dir := abb.raiz.ubicar(clave, abb.raiz, 0, abb.funcCmp)

	if ptr != nil && dir == 0 {
		return (*ptr).dato
	}
	panic("NO encontrado")
}

func (abb *abb[K, V]) Obtener(clave K) V {
	return abb.raiz.dato // por ahora para no marcar error en el ide
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
	//
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}
