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
	cmp      funcCmp[K]
}

//No se que hice
// type iteradorABB[K comparable, V any] struct {
// 	abb    *abb[K, V]
// 	indice int
// }

// func CrearABB[K comparable, V any](function_cmp func(K, K) int) DiccionarioOrdenado[K, V] {
// 	return &abb[K, V]{raiz: nil, cantidad: 0, cmp: function_cmp}
// }

// func (abb *abb[K, V]) IterarRango(desde *K, hasta *K, visitar func(clave K, dato V) bool) {

// }

// func (abb *abb[K, V]) IteradorRango(desde *K, hasta *K) IterDiccionario[K, V] {
// 	return abb.Iterador()
// }

// func (iter *iteradorABB[K, V]) avanzar(){
// 	if iter.abb.raiz.izquierdo
// }
