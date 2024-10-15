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
	cmp      funcCmp[K] // no estoy seguro de esto
}

// Hay que hacer todas las funciones otra vez

func CrearABB[K comparable, V any](function_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
	return &abb[K, V]{raiz: nil, cantidad: 0, funcCmp: function_cmp}
}

func (abb *abb[K, V]) Guardar(clave K, dato V) {
	//
}

func (abb *abb[K, V]) Pertenece(clave K) bool {
	return true
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

func (abb *abb[K, V]) Iterar(func(clave K, dato V) bool) {
	//
}

func (abb *abb[K, V]) Iterador() IterDiccionario[K, V] {
	//
}

func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {
	//
}

func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {

}
